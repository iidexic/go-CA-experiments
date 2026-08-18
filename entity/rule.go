package entity

import "sort"

// Ruleset defines a plug-in CA rule. See docs/rules.md for the design.
type Ruleset interface {
	Name() string
	Init(g *GridEntity)
	Step(g *GridEntity)
	Overlay(g *GridEntity, mode int) []byte
}

type ruleCtor func() Ruleset

var rulesetRegistry = map[string]ruleCtor{}

// RegisterRuleset associates a name with a constructor. Called from init() in each rule file.
func RegisterRuleset(name string, ctor ruleCtor) {
	rulesetRegistry[name] = ctor
}

// NewRuleset constructs a registered ruleset by name. Returns nil if unknown.
func NewRuleset(name string) Ruleset {
	ctor, ok := rulesetRegistry[name]
	if !ok {
		return nil
	}
	return ctor()
}

// RulesetNames returns the sorted names of all registered rulesets.
func RulesetNames() []string {
	names := make([]string, 0, len(rulesetRegistry))
	for name := range rulesetRegistry {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// SetRuleset installs r and calls its Init on the grid.
func (grid *GridEntity) SetRuleset(r Ruleset) {
	grid.ruleset = r
	if r != nil {
		r.Init(grid)
	}
}

// Ruleset returns the currently installed ruleset. May be nil.
func (grid *GridEntity) Ruleset() Ruleset { return grid.ruleset }

// Simstep runs one tick via the installed ruleset. No-op if none installed.
func (grid *GridEntity) Simstep() {
	if grid.ruleset == nil {
		return
	}
	grid.ruleset.Step(grid)
}

// Overlay returns a debug overlay from the current ruleset. Returns nil if
// none installed or the ruleset does not produce overlays.
func (grid *GridEntity) Overlay(mode int) []byte {
	if grid.ruleset == nil {
		return nil
	}
	return grid.ruleset.Overlay(grid, mode)
}
