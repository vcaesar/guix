// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parts

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/outer"
)

type BackgroundBorderPainterOuter interface {
	outer.Redrawer
}

type BackgroundBorderPainter struct {
	outer BackgroundBorderPainterOuter
	brush guix.Brush
	pen   guix.Pen
}

func (b *BackgroundBorderPainter) Init(outer BackgroundBorderPainterOuter) {
	b.outer = outer
	b.brush = guix.DefaultBrush
	b.pen = guix.DefaultPen
}

func (b *BackgroundBorderPainter) PaintBackground(c guix.Canvas, r math.Rect) {
	if b.brush.Color.A != 0 {
		w := b.pen.Width
		c.DrawRoundedRect(r, w, w, w, w, guix.TransparentPen, b.brush)
	}
}

func (b *BackgroundBorderPainter) PaintBorder(c guix.Canvas, r math.Rect) {
	if b.pen.Color.A != 0 && b.pen.Width != 0 {
		w := b.pen.Width
		c.DrawRoundedRect(r, w, w, w, w, b.pen, guix.TransparentBrush)
	}
}

func (b *BackgroundBorderPainter) BackgroundBrush() guix.Brush {
	return b.brush
}

func (b *BackgroundBorderPainter) SetBackgroundBrush(brush guix.Brush) {
	if b.brush != brush {
		b.brush = brush
		b.outer.Redraw()
	}
}

func (b *BackgroundBorderPainter) BorderPen() guix.Pen {
	return b.pen
}

func (b *BackgroundBorderPainter) SetBorderPen(pen guix.Pen) {
	if b.pen != pen {
		b.pen = pen
		b.outer.Redraw()
	}
}
