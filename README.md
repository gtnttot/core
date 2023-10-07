<h1 align="center">
    <a href="https://goki.dev">
        <img alt="GoGi Blue Logo" width="150" height="150" src="logo/gogi_logo_blue_transparent.png">
    </a>
</h1>

<p align="center">
    <a href="https://goreportcard.com/report/goki.dev/gi/v2"><img src="https://goreportcard.com/badge/goki.dev/gi/v2" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/goki.dev/gi/v2"><img src="https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat" alt="pkg.go.dev docs"></a>
    <a href="https://goki.dev/gi/v2/actions/workflows/ci.yml"><img src="https://goki.dev/gi/v2/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
    <a href="https://www.tickgit.com/browse?repo=goki.dev/gi/v2"><img src="https://badgen.net/https/api.tickgit.com/badgen/goki.dev/gi/v2" alt="TODOs"></a>
    <a href="https://goki.dev/gi/v2/releases/"><img src="https://img.shields.io/github/release/goki/gi?include_prereleases=&sort=semver&color=blue" alt="GitHub release"></a>
</p>

**NOTE:** GoKi is currently undergoing a period of significant developement to make it easier to make useful, fast, and beautiful apps and support running apps on mobile. As such, some of the information in this repository and on the [GoKi website](https://GoKi.dev) may be incorrect. Furthermore, there may be breaking changes soon, so starting new apps with this framework is not recommended at this time; if you do, please be ready to adjust to any breaking changes. If you want to accelerate the improvement of GoKi, please contribute by following the [Contribution Guidelines](https://goki.dev/docs/general/contributionguidelines/). Developement of Gi is currently happening on this branch. For the latest stable version of Gi, import version 1.3.19 and see the [v1 branch](https://goki.dev/gi/v2/tree/v1).

GoGi is part of the [GoKi](https://GoKi.dev) Go language (golang) full strength tree structure system (ki = 木 = tree in Japanese)

`package gi` is a scenegraph-based 2D and 3D GUI / graphics interface (Gi) in Go, that functions similar to HTML / CSS / SVG  and Qt.

NOTE: Requires Go version `1.18+` -- now using the new generics.

See the [Wiki](https://goki.dev/gi/v2/wiki) for more docs (increasingly extensive), [Install](https://goki.dev/gi/v2/wiki/Install) instructions (mostly basic `go build` procedure, but does now depend on `cgo` on all platforms due to `glfw`, so see details for each platform -- for mac you must now install the [Vulkan SDK](https://vulkan.lunarg.com), and [Google Groups goki-gi](https://groups.google.com/forum/#!forum/goki-gi) email list, and the new github [Discussions](https://goki.dev/gi/v2/discussions) tool.

GoGi uses the [GoKi](https://goki.dev/ki/v2) tree infrastructure to implement a scenegraph-based GUI framework in full native idiomatic Go, with minimal OS-specific backend interfaces based originally on the [Shiny](https://github.com/golang/exp/tree/master/shiny) drivers, now using [go-gl/glfw](https://github.com/go-gl/glfw) and vulkan-based [vgpu](https://goki.dev/vgpu/v2), and supporting MacOS, Linux, and Windows.

The overall design integrates existing standards and conventions from widely-used frameworks, including Qt (overall widget design), HTML / CSS (styling), and SVG (rendering).  The core `Layout` object automates most of the complexity associated with GUI construction (including scrolling), so the programmer mainly just needs to add the elements, and set their style properties -- similar to HTML.  The main 2D framework also integrates with a 3D scenegraph, supporting interesting combinations of these frameworks (see `gi3d` package and [examples/gi3d](https://goki.dev/gi/v2/tree/master/examples/gi3d)).  Currently GoGi is focused on desktop systems, but nothing should prevent adaptation to mobile. 

See [Gide](https://goki.dev/gi/v2de) for a complete, complex application written in GoGi (an IDE), and likewise the [Emergent](https://github.com/emer/emergent) neural network simulation environment (the prime motivator for the whole project), along with the various examples in this repository for lots of useful demonstrations -- start with the  [Widgets](https://goki.dev/gi/v2/tree/master/examples/widgets) example which has a bit of a tutorial introduction.

# Main Features

* Has all the standard widgets: `Button`, `Menu`, `Slider`, `TextField`, `SpinBox`, `ComboBox` etc, with tooltips, hover, focus, copy / paste (full native clipboard support), drag-n-drop -- the full set of standard GUI functionality.  See `gi/examples/widgets` for a demo of all the widgets.

* `Layout` auto-organizes and auto-sizes everything to configure interfaces that "just work" across different scales, resolutions, platforms.  Automatically remembers and reinstates window positions and sizes across sessions, and supports standard `Ctrl+` and `Ctrl-` zooming of display scale.

* CSS-based styling allows customization of everything -- native style properties are HTML compatible (with all standard `em`, `px`, `pct` etc units), including HTML "rich text" styling for all text rendering (e.g., in `Label` widget) -- can decorate any text with inline tags (`<strong>`, `<em>` etc), and even include links.  Styling is now separated out into `gist` package, for easier navigation.

* Compiles in seconds, compared to many minutes to hours for comparable alternatives such as Qt, and with minimal cgo dependency.  As of April 2019 we now depend on the [glfw](https://github.com/go-gl/glfw) cross-platform GUI infrastructure system, and as of May 2022 vulkan provides all the rendering (2D via vdraw, 3D via vphong):  [vgpu](https://goki.dev/vgpu/v2).

* Fully self-contained -- does *not* use OS-specific native widgets -- results in simpler, consistent code across platforms, and is `HiDPI` capable and scalable using standard `Ctrl/Cmd+Plus or Minus` key, and in `Preferences`.  This also allows a complete 2D GUI to be embedded into a 3D scene, for example.

* `SVG` element (in `svg` sub-package) supports SVG rendering -- used for Icons internally and available for advanced graphics displays -- see `gi/examples/svg` for viewer and start on editor, along with a number of test .svg files.

* **Model / View** paradigm with `reflect`ion-based view elements that display and manipulate all the standard Go types (in `giv` sub-package), from individual types (e.g., int, float display in a `SpinBox`, "enum" const int types in a `ComboBox` chooser) to composite data structures, including `StructView` editor of `struct` fields, `MapView` and `SliceView` displays of `map` and `slice` elements (including full editing / adding / deleting of elements), and full-featured `TableView` for a `slice`-of-`struct` and `TreeView` for GoKi trees.
    + `TreeView` enables a built-in GUI editor / inspector for designing gui elements themselves.  Just press `Control+Alt+I` in any window to pull up this editor / inspector.  Scene graphs can be automatically saved / loaded from JSON files, to provide a basic GUI designer framework -- just load and add appropriate connections..
    
* GoGi is a "standard" *retained-mode* (scenegraph-based) GUI, as compared to *immediate-mode* GUIs such as [Gio](https://gioui.org).  As such, GoGi automatically takes care of everything for you, but as a result you sacrifice control over every last detail.  Immediate mode gives you full control, but also the full burden of control -- you have to code every last behavior yourself.  In GoGi, you have extensive control through styling and closure-based "callback" methods, in the same way you would in a standard front-end web application (so it will likely be more familiar to many users), but if you want to do something very different, you will likely need to code a new type of Widget, which can be more difficult as then you need to know more about the overall infrastructure.  Thus, if you are likely to be doing fairly standard things and don't feel the need for absolute control, GoGi will likely be an easier experience.

![Screenshot of Widgets demo](screenshot.png?raw=true "Screenshot of Widgets demo")

![Screenshot of Gi3D demo](screenshot_gi3d.png?raw=true "Screenshot of Gi3D demo")

![Screenshot of GiEditor, Dark mode](screenshot_dark.png?raw=true "Screenshot of GiEditor, Dark Mode")

# Code Overview

There are three main types of 2D nodes:

* `Viewport` nodes that manage their own `image.RGBA` bitmap and can upload that directly to the `goosi.Texture` (GPU based) that then uploads directly to the `goosi.Window`.  The parent `Window` has a master `Viewport` that backs the entire window, and is what most `Widget`'s render into.
    + Popup `Dialog` and `Menu`'s have their own viewports that are layered on top of the main window viewport.
    + `SVG` and its subclass `Icon` are containers for SVG-rendering nodes.

* `Widget` nodes that use the full CSS-based styling (e.g., the Box model etc), are typically placed within a `Layout` -- they use `units` system with arbitrary DPI to transform sizes into actual rendered `dots` (term for actual raw resolution-dependent pixels -- "pixel" has been effectively co-opted as a 96dpi display-independent unit at this point).  Widgets have non-overlapping bounding boxes (`BBox` -- cached for all relevant reference frames).

* `SVG` rendering nodes that directly set properties on the `girl.Paint` object and typically have their own geometry etc -- they should be within a parent `SVG` viewport, and their geom units are determined entirely by the transforms etc and we do not support any further unit specification -- just raw float values.

General Widget method conventions:
* `SetValue` kinds of methods are wrapped in `UpdateStart` / `End`, but do NOT emit a signal.
* `SetValueAction` calls `SetValue` and emits the signal.
This allows other users of the widget that also recv the signal to not trigger themselves, but typically you want the update, so it makes sense to have that in the basic version.  `ValueView` in particular requires this kind of behavior.

The best way to see how the system works are in the `examples` directory, and by interactively modifying any existing gui using the interactive reflective editor via `Control+Alt+I`.

# Backend

The `oswin` and `oswin/driver/vkos` packages provide interface abstractions for hardware-level implementations, now using [vgpu](https://goki.dev/vgpu/v2) and [glfw](https://github.com/go-gl/glfw) (version 3.3) provides the basic platform-specific details along with a few other bits of platform-specific code.

All of the main "front end" code just deals with `image.RGBA` through the [girl](https://goki.dev/gi/v2/tree/master/girl) rendering library, using `girl.Paint` methods, which was adapted from [fogleman/gg](https://github.com/fogleman/gg), and we use [srwiley/rasterx](https://github.com/srwiley/rasterx) for CPU-based rasterization to the image, which is fast and SVG performant.   The [vgpu/vdraw](https://goki.dev/vgpu/v2/vdraw) package performs optimized GPU texture-based compositing to assemble the final display in a way that minimizes the copying of image data up to the GPU, and supports overlays such as popups and sprites.  Any 3D scene elements are accessed directly within the GPU.

# Status / News

* Version 1.3 released May, 2022, uses the new vulkan based  [vgpu](https://goki.dev/vgpu/v2) rendering framework.

* Version 1.2 released Feb, 2021, had lots of bug fixes.

* Version 1.1 released Nov, 2020, has the styling parameters and code broken out in the [gist](https://goki.dev/gi/v2/tree/master/gist) style package, and basic rendering code, including a complete text layout and rendering system, in the [girl](https://goki.dev/gi/v2/tree/master/girl) render library.

* Version 1.0 released April, 2020!  The 3D `gi3d` component is ready for use, and the code has been widely tested by students and researchers, including extensive testing under `gide`.  The API will remain stable at this point.

* Active users should join [Google Groups goki-gi](https://groups.google.com/forum/#!forum/goki-gi) emailing list to receive more detailed status updates.

* Please file [Issues](https://github.com/goki/gi/issues) for anything that does not work.

* 3/2019: `python` wrapper is now available!  you can do most of GoGi from python now.  See [README.md](https://goki.dev/gi/v2/tree/master/python/README.md) file there for more details.

# Detailed Render Logic

There are 4 major steps to making the `Scene` tree of `Widget` elements ready to render:

1. `Config` configures `Parts` and other settings on the `Widget` based on settings reflected in the fields of the Widget.  This is "structural" in terms of sub-widget parts, but not about styling.

2. `ApplyStyle` applies styling functions to the configured `Widget` to determine how it will look when rendered.  Generally it should be called after any `Config` step, and the `ReConfig` method on a widget does just this.

3. `DoLayout` gathers sizes and uses styling info etc to arrange widgets on the screen.  This generally must be done after changes made in `Config` or `ApplyStyle`, which set flags indicating that this is needed at the next render pass.  Critically, DoLayout is where the Styling settings made in ApplyStyle are "cached out" into raw pixel dots, based on the DPI and Size of the window.

4. `Render` itself does the actual drawing of the widgets to the `Scene` `Pixels` image, which is then uploaded to the `RenderWin` window that you actually see.

The window updating happens in a separate step, typically 60 times per second, which checks for flags on the `Scene` indicating if any new rendering or layout updating is needed.  This separate updating pass is designed to eliminate the possibility of collisions between updating and rendering, using atomically set flags to signal what needs to be updated at any point (these flags are not set until _after_ changes have been made, and unchanged nodes are not re-renderd).

## Initial configuration

First, there is code, e.g., in user's `mainrun` function, or other libraries, that calls `NewScene` and `NewButton` etc to construct widgets for a scene.  During this process, various parameters (fields) on the Widgets are set, that then determine e.g., how the Parts of the widget will then be configured.

## Stage.Run()

When the `scene` is put into a `Stage` container and then `Run()`, we call `ConfigScene()` on the Scene to make sure that everything is fully configured according to the settings made during initial construction.

## RenderWindow

Then, when the Scene is actually rendered to a `RenderWin` window that you can actually see, it is continually subject to various levels of updating depending on how much it has changed.

* If the `RenderWin` has a different size or DPI than the Scene was previously presented at, then `ApplyStyleScene` is called to allow `Style` functions to do things differently depending on the rendering context (i.e., "responsive" styling).  Then, a full Layout and render pass is done.

* If anything was subject to `Config` or `ApplyStyle` changes, then a new `DoLayout` pass is done.

* If only State changes were made due to events (e.g., a mouse hovering over a button), then just a `Render` update is done.

The computing of actual style sizes in terms of dots, which happens in `SetUnitContext`, has to happen twice: first in `ApplyStyle` so that the layout `GetSizes` functions have correct info about how much size a label etc takes when rendered, and then after the Layout is done, to reflect what is actually going to be rendered.

## Scrolling, other Spaces

* LayoutScroll: optional pass invoked by scrollbars to move elements relative to their previously-assigned positions.

* SVG nodes skip the Size and Layout passes, and render directly into parent SVG viewport


# v2 Rewrite TODO / Discussion

## TODO

* Get rid of extra UpdateStart / End in Config* calls -- all are wrapped in outer Config so unnec.

* get rid of all the extra Config calls in SetStyle and Render!

* fix all `Set..`, `...Action` state-changing methods to detect whether Config, layout, or just render is needed.  Also add return val so chain-based setting works.

* in general, "SetNeedsFullReRender" -> SetNeedsLayout(vp, updt) or true

* TextField:KeyChord -- need dialog!  should have a better mech for this.

* does StageMgr need ordmap or could just use slices?  do we ever actually lookup by name??

## Window is not a Ki

* contains a stack of Viewports, the first of which is the main window, rest are popups.

## Viewport is not a Ki, is the base struct for each separate tree

* Viewport represents: main window and each popup -> window, tooltip, dialog, etc.

* Viewport has a Frame field which is the base of each tree

* is passed as arg, not stored on obj: all methods take the viewport!
* so much cleaner vs. mutexes, safe finding, etc.

## Init2D -> Config

* ConfigAll called during window Show, but otherwise not called!
* manually call Config on individual elements if properties change.

## Style2D -> SetStyle

* called automatically for any state change

* Widget SetState method is called with bitflag to set: checks if state is not yet set: if so, sets flag, calls SetStyle() wrapped in UpdateStart / End.

## Robust updating logic redo

See [render.go](gi/render.go) for updated version:

* Config calls SetStyle, within an overall Update Start / End block

* ki.UpdateEnd() in Config or SetStyle calls SetRender method, which sets flag in node (all below it will be rendered) and a NeedsRender flag on the Viewport (need the viewport for this -- cache in Config, only use for setting this one flag!).

* Window iterates through Viewports on a timer loop, and checks if any have VpNeedsRender flag set, and not VpIsRendering flag -- calls RenderAll on them, passing viewport down.  *first* step is to clear VpNeedsRender flag so any further updates can register need, and each node's NeedsRender flag is cleared after rendering *starts*.  Better to err on side of extra rendering.

* If Config modifies the tree, then FullRender is needed: GetSize, DoLayout, then Render

* If State changes, Style then Render needs to be called..


