// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package views

import (
	"maps"
	"testing"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"github.com/stretchr/testify/assert"
)

func TestMapView(t *testing.T) {
	b := core.NewBody()
	NewMapView(b).SetMap(&map[string]int{"Go": 1, "C++": 3, "Python": 5})
	b.AssertRender(t, "map-view/basic")
}

func TestMapViewChange(t *testing.T) {
	b := core.NewBody()
	m := map[string]int{"Go": 1, "C++": 3, "Python": 5}

	n := 0
	value := map[string]int{}
	mv := NewMapView(b).SetMap(&m)
	mv.OnChange(func(e events.Event) {
		n++
		maps.Copy(value, m)
	})
	b.AssertRender(t, "map-view/change", func() {
		// [3] is value of second row, which is Go since it is sorted alphabetically
		mv.MapGrid().Child(3).(*core.Spinner).TrailingIconButton().Send(events.Click)
		assert.Equal(t, 1, n)
		assert.Equal(t, m, value)
		assert.Equal(t, map[string]int{"Go": 2, "C++": 3, "Python": 5}, m)
	})
}

func TestMapViewReadOnly(t *testing.T) {
	b := core.NewBody()
	NewMapView(b).SetMap(&map[string]int{"Go": 1, "C++": 3, "Python": 5}).SetReadOnly(true)
	b.AssertRender(t, "map-view/read-only")
}
