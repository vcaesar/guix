// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parts

import (
	"fmt"

	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/outer"
)

type LayoutableOuter interface {
	outer.Parenter
	outer.Redrawer
}

type Layoutable struct {
	outer             LayoutableOuter
	driver            guix.Driver
	margin            math.Spacing
	size              math.Size
	relayoutRequested bool
	inLayoutChildren  bool // True when calling LayoutChildren
}

func (l *Layoutable) Init(outer LayoutableOuter, theme guix.Theme) {
	l.outer = outer
	l.driver = theme.Driver()
}

func (l *Layoutable) SetMargin(m math.Spacing) {
	l.margin = m
	if p := l.outer.Parent(); p != nil {
		p.Relayout()
	}
}

func (l *Layoutable) Margin() math.Spacing {
	return l.margin
}

func (l *Layoutable) Size() math.Size {
	return l.size
}

func (l *Layoutable) SetSize(size math.Size) {
	if size.W < 0 {
		panic(fmt.Errorf("SetSize() called with a negative width. Size: %v", size))
	}
	if size.H < 0 {
		panic(fmt.Errorf("SetSize() called with a negative height. Size: %v", size))
	}

	sizeChanged := l.size != size
	l.size = size
	if l.relayoutRequested || sizeChanged {
		l.relayoutRequested = false
		l.inLayoutChildren = true
		callLayoutChildrenIfSupported(l.outer)
		l.inLayoutChildren = false
		l.outer.Redraw()
	}
}

func (l *Layoutable) Relayout() {
	l.driver.AssertUIGoroutine()
	if l.inLayoutChildren {
		panic("Cannot call Relayout() while in LayoutChildren")
	}
	if !l.relayoutRequested {
		if p := l.outer.Parent(); p != nil {
			l.relayoutRequested = true
			p.Relayout()
		}
	}
}
