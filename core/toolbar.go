// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/math32"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/states"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/tree"
)

// Toolbar is a [Frame] that is useful for holding [Button]s that do things.
// It automatically moves items that do not fit into an overflow menu, and
// manages additional items that are always placed onto this overflow menu.
// In general it should be possible to use a single toolbar + overflow to
// manage all an app's functionality, in a way that is portable across
// mobile and desktop environments.
// See [Widget.MakeToolbar] for the standard toolbar config method for
// any given widget, and [Scene.AppBars] for [ConfigFuncs] for [Scene]
// elements who should be represented in the main AppBar (e.g., TopAppBar).
type Toolbar struct {
	Frame

	// OverflowItems are items moved from the main toolbar that will be shown in the overflow menu.
	OverflowItems tree.Slice `set:"-" json:"-" xml:"-"`

	// OverflowMenus are functions for the overflow menu; use [Toolbar.AddOverflowMenu] to add.
	// These are processed in reverse order (last in, first called)
	// so that the default items are added last.
	OverflowMenus []func(m *Scene) `set:"-" json:"-" xml:"-"`

	// Makers contains functions for making the plan for the toolbar.
	// You can use [Toolbar.AddMaker] to add a new one.
	Makers []func(p *Plan)

	// This is the overflow button
	OverflowButton *Button `copier:"-"`
}

func (tb *Toolbar) OnInit() {
	tb.Frame.OnInit()
	ToolbarStyles(tb)
}

// AddMaker adds the given function(s) for making the plan for the toolbar.
func (tb *Toolbar) AddMaker(m ...func(p *Plan)) *Toolbar {
	tb.Makers = append(tb.Makers, m...)
	return tb
}

func (tb *Toolbar) IsVisible() bool {
	// do not render toolbars with no buttons
	return tb.WidgetBase.IsVisible() && len(tb.Kids) > 0
}

// AppChooser returns the app [Chooser] used for searching for
// items. It will only be non-nil if this toolbar has been configured
// with an app chooser, which typically only happens for app bars.
func (tb *Toolbar) AppChooser() *Chooser {
	ch, _ := tb.ChildByName("app-chooser").(*Chooser)
	return ch
}

func (tb *Toolbar) Make(p *Plan) {
	for _, f := range tb.Makers {
		f(p)
	}
}

func (tb *Toolbar) SizeUp() {
	tb.AllItemsToChildren()
	tb.Frame.SizeUp()
}

// todo: try doing move to overflow in Final

func (tb *Toolbar) SizeDown(iter int) bool {
	redo := tb.Frame.SizeDown(iter)
	if iter == 0 {
		return true // ensure a second pass
	}
	tb.MoveToOverflow()
	return redo
}

func (tb *Toolbar) SizeFromChildren(iter int, pass LayoutPasses) math32.Vector2 {
	csz := tb.Frame.SizeFromChildren(iter, pass)
	if pass == SizeUpPass || (pass == SizeDownPass && iter == 0) {
		dim := tb.Styles.Direction.Dim()
		ovsz := tb.OverflowButton.Geom.Size.Actual.Total.Dim(dim)
		csz.SetDim(dim, ovsz) // present the minimum size initially
		return csz
	}
	return csz
}

// AllItemsToChildren moves the overflow items back to the children,
// so the full set is considered for the next layout round,
// and ensures the overflow button is made and moves it
// to the end of the list.
func (tb *Toolbar) AllItemsToChildren() {
	if tb.OverflowButton == nil {
		ic := icons.MoreVert
		if tb.Styles.Direction != styles.Row {
			ic = icons.MoreHoriz
		}
		tb.OverflowButton = NewButton(tb).SetIcon(ic).SetTooltip("Additional menu items")
		tb.OverflowButton.SetName("overflow-menu")
		tb.OverflowButton.Menu = tb.OverflowMenu
	}
	if len(tb.OverflowItems) > 0 {
		tb.Kids = append(tb.Kids, tb.OverflowItems...)
		tb.OverflowItems = nil
	}
	ovi := -1
	for i, k := range tb.Kids {
		_, wb := AsWidget(k)
		if wb.This() == tb.OverflowButton.This() {
			ovi = i
			break
		}
	}
	if ovi >= 0 {
		tb.Kids.DeleteAtIndex(ovi)
	}
	tb.Kids = append(tb.Kids, tb.OverflowButton.This())
	tb.OverflowButton.Update()
}

func (tb *Toolbar) ParentSize() float32 {
	ma := tb.Styles.Direction.Dim()
	psz := tb.ParentWidget().Geom.Size.Alloc.Content.Sub(tb.Geom.Size.Space)
	avail := psz.Dim(ma)
	return avail
}

// MoveToOverflow moves overflow out of children to the OverflowItems list
func (tb *Toolbar) MoveToOverflow() {
	ma := tb.Styles.Direction.Dim()
	avail := tb.ParentSize()
	ovsz := tb.OverflowButton.Geom.Size.Actual.Total.Dim(ma)
	avsz := avail - ovsz
	sz := &tb.Geom.Size
	sz.Alloc.Total.SetDim(ma, avail)
	sz.SetContentFromTotal(&sz.Alloc)
	n := len(tb.Kids)
	ovidx := n - 1
	hasOv := false
	szsum := float32(0)
	tb.WidgetKidsIter(func(i int, kwi Widget, kwb *WidgetBase) bool {
		if i >= n-1 {
			return tree.Break
		}
		ksz := kwb.Geom.Size.Alloc.Total.Dim(ma)
		szsum += ksz
		if szsum > avsz {
			if !hasOv {
				ovidx = i
				hasOv = true
			}
			tb.OverflowItems = append(tb.OverflowItems, kwi)
		}
		return tree.Continue
	})
	if ovidx != n-1 {
		tb.Kids.Move(n-1, ovidx)
		tb.Kids = tb.Kids[:ovidx+1]
	}
	if len(tb.OverflowItems) == 0 && len(tb.OverflowMenus) == 0 {
		tb.OverflowButton.SetState(true, states.Invisible)
	} else {
		tb.OverflowButton.SetState(false, states.Invisible)
		tb.OverflowButton.Update()
	}
}

// OverflowMenu is the overflow menu function
func (tb *Toolbar) OverflowMenu(m *Scene) {
	nm := len(tb.OverflowMenus)
	if len(tb.OverflowItems) > 0 {
		for _, k := range tb.OverflowItems {
			if k.This() == tb.OverflowButton.This() {
				continue
			}
			cl := k.This().Clone()
			m.AddChild(cl)
			cl.This().(Widget).Build()
		}
		if nm > 1 { // default includes sep
			NewSeparator(m)
		}
	}
	// reverse order so defaults are last
	for i := nm - 1; i >= 0; i-- {
		fn := tb.OverflowMenus[i]
		fn(m)
	}
}

// AddOverflowMenu adds the given menu function to the overflow menu list.
// These functions are called in reverse order such that the last added function
// is called first when constructing the menu.
func (tb *Toolbar) AddOverflowMenu(fun func(m *Scene)) {
	tb.OverflowMenus = append(tb.OverflowMenus, fun)
}

//////////////////////////////////////////////////////////////////////////////
// 	ToolbarStyles

// ToolbarStyles can be applied to any layout (e.g., Frame) to achieve
// standard toolbar styling.
func ToolbarStyles(ly Layouter) {
	lb := ly.AsLayout()
	ly.Style(func(s *styles.Style) {
		s.Border.Radius = styles.BorderRadiusFull
		s.Background = colors.C(colors.Scheme.SurfaceContainer)
		s.Gap.Zero()
		s.Align.Items = styles.Center
	})
	ly.AsWidget().StyleFinal(func(s *styles.Style) {
		if s.Direction == styles.Row {
			s.Grow.Set(1, 0)
			s.Padding.SetHorizontal(units.Dp(16))
		} else {
			s.Grow.Set(0, 1)
			s.Padding.SetVertical(units.Dp(16))
		}
	})
	ly.OnWidgetAdded(func(w Widget) {
		if bt := AsButton(w); bt != nil {
			bt.Type = ButtonAction
			return
		}
		if sp, ok := w.(*Separator); ok {
			sp.Style(func(s *styles.Style) {
				s.Direction = lb.Styles.Direction.Other()
			})
		}
	})
}

// BasicBar is just a styled Frame layout for holding buttons
// and other widgets. Use this when the more advanced features
// of the [Toolbar] are not needed.
type BasicBar struct {
	Frame
}

func (tb *BasicBar) OnInit() {
	tb.Frame.OnInit()
	ToolbarStyles(tb)
}

// UpdateBar calls ApplyStyleUpdate to update to current state
func (tb *BasicBar) UpdateBar() {
	tb.ApplyStyleUpdate()
}
