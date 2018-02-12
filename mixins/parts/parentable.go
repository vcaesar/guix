// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parts

import (
	"github.com/vcaesar/guix"
)

type ParentableOuter interface{}

type Parentable struct {
	outer  ParentableOuter
	parent guix.Parent
}

func (p *Parentable) Init(outer ParentableOuter) {
	p.outer = outer
}

func (p *Parentable) Parent() guix.Parent {
	return p.parent
}

func (p *Parentable) SetParent(parent guix.Parent) {
	p.parent = parent
}
