# Aperture

View-transform layer between `GridEntity` and the screen. Decouples the CA grid's pixel dimensions and world position from the window, and provides pan/zoom navigation.

## Purpose

The CA grid holds an `Img` (an `*ebiten.Image` sized to the grid's pixel dimensions) and a `Px` byte slice representing cell state. Previously the grid was drawn directly to screen with a fixed `Op` (`screen.DrawImage(grid.Img, grid.Op)`), locking the grid's on-screen size and position to its underlying dimensions.

Aperture replaces that direct draw with a view-transform:

- The grid keeps ownership of `Img` and `Px`.
- The aperture owns the mapping from grid-pixel space to screen-pixel space.
- The aperture owns the draw call and any user interaction that navigates the view (pan, zoom, cell selection).

## Scope decisions

| Decision | Choice | Notes |
|---|---|---|
| Zoom mode | Integer steps initially, API accepts `float32` | Continuous zoom is a future extension; only integer values are set at first. |
| Pan input | Right-click drag | Does not clash with existing keybindings (`CutoffUp/Down` on arrows). |
| Instance count | Singleton per scene | One `*Aperture` on `GameSim`. Registry/multi-aperture deferred until minimap or split-view is needed. |
| Grid ownership | Aperture holds `*GridEntity` pointer | Singleton scope means 1:1 relationship; passing the grid on every `Draw` call adds no value. |
| Rendering | `Img.SubImage(sourceRect)` + `DrawImageOptions` scale+translate | Ebiten native, avoids intermediate buffers. |

## State

```go
type Aperture struct {
    grid          *GridEntity     // source of pixel data
    gridTopLeft   image.Point     // top-left of visible region, in grid-pixel coords
    screenBounds  image.Rectangle // where on screen the aperture draws
    zoom          float32         // screen pixels per grid pixel (integer-valued for now)
    minZoom       float32
    maxZoom       float32

    // pan-drag state
    dragging      bool
    dragAnchor    image.Point     // cursor position at drag start (screen coords)
    dragOrigin    image.Point     // gridTopLeft at drag start
}
```

`gridTopLeft` + `zoom` + `screenBounds.Size()` derive everything else. The visible grid rectangle is `image.Rect(gridTopLeft.X, gridTopLeft.Y, gridTopLeft.X + screenBounds.Dx()/int(zoom), gridTopLeft.Y + screenBounds.Dy()/int(zoom))`.

## API

```go
// Construction
NewAperture(grid *GridEntity, screenBounds image.Rectangle) *Aperture

// Navigation
(a *Aperture) Pan(dxGrid, dyGrid int)            // shift gridTopLeft in grid pixels
(a *Aperture) Zoom(delta float32, anchor image.Point) // scale around a screen-coord anchor
(a *Aperture) SetZoom(z float32, anchor image.Point)  // absolute zoom, same anchor semantics

// Input helpers (called by input layer)
(a *Aperture) BeginPan(cursor image.Point)
(a *Aperture) UpdatePan(cursor image.Point)
(a *Aperture) EndPan()

// Coordinate conversion
(a *Aperture) ScreenToGrid(sx, sy int) (gx, gy int, inside bool)
(a *Aperture) GridToScreen(gx, gy int) (sx, sy int)

// Render
(a *Aperture) Draw(screen *ebiten.Image)
```

## Zoom semantics

- `Zoom(delta, anchor)` treats `delta` as an integer step for now: `newZoom = clamp(zoom + delta, minZoom, maxZoom)`.
- The grid point currently under `anchor` (screen coord) stays under `anchor` after the zoom. This is the standard "zoom to cursor" behavior.
- Defaults: `minZoom = 1`, `maxZoom = 32`. Chosen because a 1:1 grid pixel is the smallest useful view and 32× is well beyond typical inspection needs. Revisit if too limiting.

## Pan semantics

- `Pan(dx, dy)` shifts `gridTopLeft` by grid-pixel deltas. Right-drag input passes screen-pixel deltas divided by zoom.
- `gridTopLeft` is clamped so at least one column and one row of grid remain visible; over-pan into empty space is prevented.
- `BeginPan`/`UpdatePan`/`EndPan` are the input-driven entry points. `UpdatePan` computes the delta from `dragAnchor` and adjusts `gridTopLeft` relative to `dragOrigin`, avoiding drift from repeated float rounding.

## Rendering

`Draw(screen)` steps:

1. Compute the source rectangle on `grid.Img` from `gridTopLeft` and `screenBounds.Size() / zoom`, clamped to grid bounds.
2. Take `sub := grid.Img.SubImage(sourceRect).(*ebiten.Image)`.
3. Build `op := &ebiten.DrawImageOptions{}`, apply `op.GeoM.Scale(float64(zoom), float64(zoom))` then `op.GeoM.Translate(float64(screenBounds.Min.X), float64(screenBounds.Min.Y))`.
4. `screen.DrawImage(sub, op)`.

Nearest-neighbor sampling is Ebiten's default and matches the "integer zoom, pixel-perfect CA" aesthetic.

## Input map (initial)

Wired in `core/ezinput.go`:

- Right mouse button press: `aperture.BeginPan(cursor)`
- Right mouse button held: `aperture.UpdatePan(cursor)`
- Right mouse button release: `aperture.EndPan()`
- Mouse wheel up: `aperture.Zoom(+1, cursor)`
- Mouse wheel down: `aperture.Zoom(-1, cursor)`

Existing key bindings (arrows, R, G, Q, Enter, Escape) remain unchanged.

## Wiring

- `GameSim` gains an `aperture *entity.Aperture` field.
- `GameSimInit` constructs the aperture with `screenBounds = image.Rect(0, 0, g.gWidth, g.gHeight)` after `maingrid` is built.
- `GameSim.Draw` replaces `screen.DrawImage(g.maingrid.Img, g.maingrid.Op)` with `g.aperture.Draw(screen)`, gated on `g.maingrid.Visible`.
- `GridEntity.Op` is no longer read from `GameSim.Draw`; kept on the struct for now in case debug overlay path uses it.

## Out of scope (this pass)

- Continuous (non-integer) zoom
- Cell-selection interactions (planned after aperture lands)
- Overlays drawn on top of aperture output (planned with `GridToScreen` consumers)
- Multiple apertures, minimap, split-view
- Rebinding keyboard controls
- Touch / gamepad input

## Deferred behavior (known, revisit later)

- Aperture currently spans the full window (`image.Rect(0, 0, gWidth, gHeight)`) and the grid is anchored at its top-left. Previously the grid was drawn centered in the window with a border. Long-term the aperture will be smaller than the window (to reserve space for UI panels) and clamped to never exceed the grid size, at which point centering / positioning becomes a layout concern rather than an aperture concern. Not fixing here.

## Roadmap position

Per the recommended order in the branch cleanup discussion:

1. **Aperture** (this doc)
2. Rule system refactor — interface + registry
3. Multi-attribute cell (SoA extension)
4. UI panels (ebitenui or equivalent)
5. Double-buffer + goroutine bands for parallel rule evaluation
6. GPU migration decision point
