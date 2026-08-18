# Rule System

Decouples rule definitions from `GridEntity` so different rulesets can be swapped in without editing grid code. First step toward the original project goal of hot-swappable rulesets.

## Purpose

`GridEntity` currently owns both the CA state (pixel buffer, RNG, zones, ticks) and the LVSD rule implementation (`SimstepLVSD`, `calculateInteraction`, `versusLVSD`, `battlemc`, outcome dispatch, `Cutoff*`, debug overlay). This entangles the "what" (grid storage) with the "how" (rule logic) and blocks adding new rulesets.

The refactor:

- Grid keeps the state and provides primitives for rules to use.
- A `Ruleset` interface defines the per-tick behavior contract.
- A registry lets code look up rulesets by name so future hot-swap and rule-switching UI can plug in without changing the grid.
- LVSD becomes the first concrete `Ruleset` implementation and continues to work unchanged from the user-facing perspective.

## Scope decisions

| Decision | Choice | Notes |
|---|---|---|
| Interface shape | `Ruleset` owns the tick loop via `Step(*GridEntity)` | Most flexible; each ruleset writes its own iteration (multi-pass, zone-first, region-first, whatever). LVSD keeps its 2x per-cell loop as-is. |
| Package layout | Single `entity` package, split files | `entity/rule.go` for interface + registry, `entity/rule_lvsd.go` for LVSD impl. No import surface change; primitives (`bavg`, `sliceToward`, `pxisort`, etc.) stay accessible without exports. Split into subpackages later when a 2nd ruleset lands. |
| Switching | Runtime `SetRuleset` supported from day one | Matches the original hot-swap goal. `SetRuleset` calls the new ruleset's `Init` so per-rule state is reset on install. |
| Rule-tunable state | Instance fields on the ruleset | `testCutoff` (LVSD-specific) moves off the package level onto a `LVSD` struct field so switching rulesets does not leak stale state. |
| Debug overlay | Ruleset responsibility | `ApplyDbgOverlay` is LVSD-flavored (red/blue/mid by average). It becomes `LVSD.Overlay(g, mode)`. Grid exposes a generic `grid.Overlay(mode)` that delegates. |

## Interface

```go
// Ruleset defines a plug-in CA rule.
type Ruleset interface {
    Name() string
    Init(g *GridEntity)                       // called on install; may reset per-rule state
    Step(g *GridEntity)                       // one tick
    Overlay(g *GridEntity, mode int) []byte   // debug overlay bytes; may be nil if unsupported
}
```

## Registry

```go
type ruleCtor func() Ruleset

// RegisterRuleset associates a name with a constructor. Called from init() in each rule file.
func RegisterRuleset(name string, ctor ruleCtor)

// NewRuleset constructs a registered ruleset by name. Returns nil if unknown.
func NewRuleset(name string) Ruleset

// RulesetNames returns the sorted names of all registered rulesets.
func RulesetNames() []string
```

Registration pattern: each rule file has an `init()` that calls `RegisterRuleset("name", func() Ruleset { return &Impl{} })`. Imports of `entity` guarantee rule files' `init()` fires.

## Grid additions

```go
type GridEntity struct {
    // ... existing fields ...
    ruleset Ruleset
}

func (grid *GridEntity) SetRuleset(r Ruleset)         // installs r, calls r.Init(grid)
func (grid *GridEntity) Ruleset() Ruleset             // returns the current ruleset (may be nil)
func (grid *GridEntity) Simstep()                     // delegates to grid.ruleset.Step(grid); no-op if nil
func (grid *GridEntity) Overlay(mode int) []byte      // delegates to grid.ruleset.Overlay(grid, mode); nil if none
```

`SimstepLVSD` is removed. `ApplyDbgOverlay` is removed. All callers go through the interface.

## LVSD extraction

`entity/rule_lvsd.go` gets:

- `type LVSD struct { cutoff byte }` — cutoff is now instance state, initialized to 127 by `NewLVSD`.
- `func NewLVSD() *LVSD` — constructor.
- `func init() { RegisterRuleset("lvsd", func() Ruleset { return NewLVSD() }) }`
- `Name()`, `Init(g)`, `Step(g)`, `Overlay(g, mode)` implementing `Ruleset`.
- Private helpers moved verbatim from `gridentity.go`, converted from `(grid *GridEntity)` receivers to `(r *LVSD)` methods that take `g *GridEntity` as an argument: `calculateInteraction`, `exec1v1`, `interactMine`, `interactStalemate`, `interactFriend`.
- Free helpers moved: `versusLVSD`, `battlemc`, `outcome` enum + constants.
- Overlay color vars (`overlayRed`, `overlayBlue`, `overlayMid`) move here since only overlay uses them.
- `CutoffUp() / CutoffDown() / CutoffIs()` become methods on `*LVSD` and operate on `r.cutoff`. Free-func versions in the `entity` package are removed.

## Callers touched

- `core/sceneSim.go`
  - `GameSimInit`: after `MakeGridDefault`, `g.maingrid.SetRuleset(entity.NewRuleset("lvsd"))`.
  - `Update`: `g.maingrid.SimstepLVSD()` -> `g.maingrid.Simstep()`.
  - `Update`: `g.maingrid.ApplyDbgOverlay(0)` -> `g.maingrid.Overlay(0)`.
  - `debugUpdate`: `entity.CutoffIs()` -> access via ruleset type-assert (see below).
- `core/ezinput.go`
  - Arrow-left/right cutoff calls go through a `GameSim.lvsd()` helper that type-asserts `g.maingrid.Ruleset().(*entity.LVSD)` and returns nil when the ruleset is something else. Callers guard with `if lv := g.lvsd(); lv != nil { lv.CutoffUp() }`.
- `entity/entity_test.go`
  - `benchmarkGridLVSD`: install ruleset after grid construction and call `g.Simstep()` instead of `g.SimstepLVSD()`.

## Out of scope (this pass)

- Splitting rules into subpackages (defer until a 2nd ruleset arrives).
- Hot-swapping via UI or key binding (interface supports it; wiring an input path is deferred).
- Rule-registered key bindings (interface can grow a `Controls()` method later; ezinput hardcodes LVSD arrow keys for now via the type-assert helper).
- Rule-registered debug overlay modes beyond a single int argument.
- Moving primitives (`bavg`, `sliceToward`, `pxisort`, etc.) to their own package.
- Runtime rule configuration (loading params from a file, etc.).

## Roadmap position

Per `docs/aperture.md`:

1. Aperture (done)
2. **Rule system refactor** (this doc)
3. Multi-attribute cell (SoA extension)
4. UI panels
5. Double-buffer + goroutine bands for parallel rule evaluation
6. GPU migration decision point
