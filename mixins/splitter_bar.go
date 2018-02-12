// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/base"
)

type SplitterBarOuter interface {
	base.ControlOuter
}

type SplitterBar struct {
	base.Control

	onDrag          func(wndPnt math.Point)
	outer           SplitterBarOuter
	theme           guix.Theme
	onDragStart     guix.Event
	onDragEnd       guix.Event
	backgroundColor guix.Color
	foregroundColor guix.Color
	isDragging      bool
}

func (b *SplitterBar) Init(outer SplitterBarOuter, theme guix.Theme) {
	b.Control.Init(outer, theme)

	b.outer = outer
	b.theme = theme
	b.onDragStart = guix.CreateEvent(func(guix.MouseEvent) {})
	b.onDragEnd = guix.CreateEvent(func(guix.MouseEvent) {})
	b.backgroundColor = guix.Red
	b.foregroundColor = guix.Green
}

func (b *SplitterBar) SetBackgroundColor(c guix.Color) {
	b.backgroundColor = c
}

func (b *SplitterBar) SetForegroundColor(c guix.Color) {
	b.foregroundColor = c
}

func (b *SplitterBar) OnSplitterDragged(f func(wndPnt math.Point)) {
	b.onDrag = f
}

func (b *SplitterBar) IsDragging() bool {
	return b.isDragging
}

func (b *SplitterBar) OnDragStart(f func(guix.MouseEvent)) guix.EventSubscription {
	return b.onDragStart.Listen(f)
}

func (b *SplitterBar) OnDragEnd(f func(guix.MouseEvent)) guix.EventSubscription {
	return b.onDragEnd.Listen(f)
}

// parts.DrawPaint overrides
func (b *SplitterBar) Paint(c guix.Canvas) {
	r := b.outer.Size().Rect()
	c.DrawRect(r, guix.CreateBrush(b.backgroundColor))
	if b.foregroundColor != b.backgroundColor {
		c.DrawRect(r.ContractI(1), guix.CreateBrush(b.foregroundColor))
	}
}

// InputEventHandler overrides
func (b *SplitterBar) MouseDown(e guix.MouseEvent) {
	b.isDragging = true
	b.onDragStart.Fire(e)
	var mms, mus guix.EventSubscription
	mms = e.Window.OnMouseMove(func(we guix.MouseEvent) {
		if b.onDrag != nil {
			b.onDrag(we.WindowPoint)
		}
	})
	mus = e.Window.OnMouseUp(func(we guix.MouseEvent) {
		mms.Unlisten()
		mus.Unlisten()
		b.isDragging = false
		b.onDragEnd.Fire(we)
	})

	b.InputEventHandler.MouseDown(e)
}
