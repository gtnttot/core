// Copyright (c) 2018, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"fmt"
	"image"
	"time"

	"goki.dev/goosi/events/key"
	"goki.dev/mat32/v2"
)

var (
	// DoubleClickInterval is the maximum time interval between button press
	// events to count as a double-click.
	// This is also in gi.Prefs and updated from there.
	DoubleClickInterval = 500 * time.Millisecond

	// ScrollWheelSpeed controls how fast the scroll wheel moves (typically
	// interpreted as pixels per wheel step) -- only relevant for some OS's which
	// do not have a native preference for this setting, e.g., X11
	// This is also in gi.Prefs and updated from there
	ScrollWheelSpeed = float32(20)
)

// Buttons is a mouse button.
type Buttons int32 //enums:enum

// TODO: have a separate axis concept for wheel up/down? How does that relate
// to joystick events?

const (
	NoButton Buttons = iota
	Left
	Middle
	Right
)

// Mouse is a basic mouse event for all mouse events except Scroll
type Mouse struct {
	Base

	// TODO: have a field to hold what other buttons are down, for detecting
	// drags or button-chords.

	// TODO: add a Device ID, for multiple input devices?
}

func NewMouse(typ Types, but Buttons, where image.Point, mods key.Modifiers) *Mouse {
	ev := &Mouse{}
	ev.Typ = typ
	ev.SetUnique()
	ev.Button = but
	ev.Where = where
	ev.Mods = mods
	return ev
}

func (ev *Mouse) String() string {
	return fmt.Sprintf("%v{Button: %v, Pos: %v, Mods: %v, Time: %v}", ev.Type(), ev.Button, ev.Where, key.ModsString(ev.Mods), ev.Time())
}

func (ev Mouse) HasPos() bool {
	return true
}

func NewMouseMove(but Buttons, where, prev image.Point, mods key.Modifiers) *Mouse {
	ev := &Mouse{}
	ev.Typ = MouseMove
	// not unique
	ev.Button = but
	ev.Where = where
	ev.Prev = prev
	ev.Mods = mods
	return ev
}

func NewMouseDrag(but Buttons, where, prev, start image.Point, mods key.Modifiers) *Mouse {
	ev := &Mouse{}
	ev.Typ = MouseDrag
	// not unique
	ev.Button = but
	ev.Where = where
	ev.Prev = prev
	ev.Start = start
	ev.Mods = mods
	return ev
}

// MouseScroll is for mouse scrolling, recording the delta of the scroll
type MouseScroll struct {
	Mouse

	// Delta is the amount of scrolling in each axis
	Delta image.Point
}

func (ev *MouseScroll) String() string {
	return fmt.Sprintf("%v{Delta: %v, Pos: %v, Mods: %v, Time: %v}", ev.Type(), ev.Delta, ev.Where, key.ModsString(ev.Mods), ev.Time())
}

func NewScroll(where, delta image.Point, mods key.Modifiers) *MouseScroll {
	ev := &MouseScroll{}
	ev.Typ = Scroll
	// not unique, but delta integrated!
	ev.Where = where
	ev.Delta = delta
	ev.Mods = mods
	return ev
}

// DimDelta returns the scrolling delta change for the given axis.
// for X dimension, it returns any non-zero Y delta because often
// scroll wheels only report Y deltas.
func (ev MouseScroll) DimDelta(dim mat32.Dims) int {
	if dim == mat32.X {
		// TODO(kai): remove this function unless we are adding back this
		// if ev.Delta.X == 0 {
		// 	return ev.Delta.Y
		// }
		return ev.Delta.X
	}
	return ev.Delta.Y
}
