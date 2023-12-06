// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build notyet

package main

import (
	"log"

	"goki.dev/colors"
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/gimain"
	"goki.dev/gi/v2/giv"
	"goki.dev/gi/v2/icons"
	"goki.dev/gi/v2/units"
	"goki.dev/gi/v2/xyz"
	"goki.dev/ki/v2"
	"goki.dev/mat32/v2"
)

func main() { gimain.Run(app) }

func app() {
	width := 1600
	height := 1200

	rec := ki.Node{}          // receiver for events
	rec.InitName(&rec, "rec") // this is essential for root objects not owned by other Ki tree nodes

	gi.SetAppName("xyzview")
	gi.SetAppAbout(`This is a viewer for the 3D graphics aspect of the <b>GoGi</b> graphical interface system, within the <b>GoKi</b> tree framework.  See <a href="https://github.com/goki">GoKi on GitHub</a>.
<p>The <a href="https://goki.dev/gi/v2/blob/master/examples/xyzviewer/README.md">README</a> page for this example app has further info.</p>`)

	win := gi.NewMainWindow("xyz-viewer", "GoGi 3D Viewer", width, height)

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	mfr := win.SetMainFrame()
	mfr.SetProp("spacing", units.Ex(1))

	tbar := gi.NewToolbar(mfr, "tbar")
	tbar.SetStretchMaxWidth()

	//////////////////////////////////////////
	//    Scene

	gi.NewSpace(mfr, "scspc")
	scvw := xyz.NewSceneView(mfr, "sceneview")
	scvw.SetStretchMax()
	scvw.Config()
	sc := scvw.Scene()

	// first, add lights, set camera
	sc.BackgroundColor = colors.FromRGB(230, 230, 255) // sky blue-ish
	xyz.NewAmbientLight(sc, "ambient", 0.3, xyz.DirectSun)

	dir := xyz.NewDirLight(sc, "dir", 1, xyz.DirectSun)
	dir.Pos.Set(0, 2, 1) // default: 0,1,1 = above and behind us (we are at 0,0,X)

	// point := xyz.NewPointLight(sc, "point", 1, xyz.DirectSun)
	// point.Pos.Set(0, 5, 5)

	// spot := xyz.NewSpotLight(sc, "spot", 1, xyz.DirectSun)
	// spot.Pose.Pos.Set(0, 5, 5)

	sc.Camera.LookAt(mat32.Vec3Zero, mat32.Vec3Y) // defaults to looking at origin

	objgp := xyz.NewGroup(sc, sc, "obj-gp")

	_, err := sc.OpenNewObj("objs/airplane_prop_001.obj", objgp)
	if err != nil {
		log.Println(err)
	}

	curFn := ""
	exts := ".obj,.dae,.gltf"

	tbar.AddAction(gi.ActOpts{Label: "Open...", Icon: icons.Open, Tooltip: "Open a 3D object file for viewing."}, win.This(), func(recv, send ki.Ki, sig int64, data any) {
		giv.FileViewDialog(vp, curFn, exts, giv.DlgOpts{Title: "Open 3D Object", Prompt: "Open a 3D object file for viewing."}, nil,
			win.This(), func(recv, send ki.Ki, sig int64, data any) {
				if sig == int64(gi.DialogAccepted) {
					dlg, _ := send.Embed(gi.TypeDialog).(*gi.Dialog)
					fn := giv.FileViewDialogValue(dlg)
					curFn = fn
					updt := sc.UpdateStart()
					objgp.DeleteChildren(true)
					sc.DeleteMeshes()
					sc.DeleteTextures()
					ki.DelMgr.DestroyDeleted() // this is actually essential to prevent leaking memory!
					_, err := sc.OpenNewObj(fn, objgp)
					if err != nil {
						log.Println(err)
					}
					sc.SetCamera("default")
					sc.UpdateEnd(updt)
				}
			})
	})

	appnm := gi.AppName()
	mmen := win.MainMenu
	mmen.ConfigMenus([]string{appnm, "File", "Edit", "Window"})

	amen := win.MainMenu.ChildByName(appnm, 0).(*gi.Button)
	amen.Menu.AddAppMenu(win)

	win.MainMenuUpdated()
	vp.UpdateEndNoSig(updt)
	win.StartEventLoop()
}
