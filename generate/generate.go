// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package generate provides the generation
// of useful methods, variables, and constants
// for GoKi code.
package generate

import (
	"bytes"
	"fmt"

	"goki.dev/goki/config"
)

// Generate is the main entry point to code generation
// that does all of the generation according to the
// given config info.
func Generate(c *config.Config) error {
	// err := enumgen.Generate(&c.Generate.Enumgen)
	// if err != nil {
	// 	return fmt.Errorf("error running enumgen: %w", err)
	// }

	g := NewGenerator(c)
	err := g.ParsePackage()
	if err != nil {
		return fmt.Errorf("Generate: error parsing package: %w", err)
	}
	for _, pkg := range g.Pkgs {
		g.Pkg = pkg
		g.Buf = bytes.Buffer{}
		g.PrintHeader()
		err := g.Write()
		if err != nil {
			return fmt.Errorf("Generate: error writing code: %w", err)
		}
	}
	return nil
}
