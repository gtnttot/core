// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testdata

import "fmt"

// Person represents a person and their attributes.
// The zero value of a Person is not valid.
//
//gti:add -type-var -instance
//ki:flagtype NodeFlags -field Flag
type Person struct {
	// Name is the name of the person
	//gi:toolbar -hide
	Name string
	// Age is the age of the person
	//gi:view inline
	Age int
}

func (p Person) Introduction() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}
