// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/base"
	"github.com/vcaesar/guix/mixins/parts"
)

type ProgressBarOuter interface {
	base.ControlOuter
	PaintProgress(guix.Canvas, math.Rect, float32)
}

type ProgressBar struct {
	base.Control
	parts.BackgroundBorderPainter

	outer            ProgressBarOuter
	desiredSize      math.Size
	progress, target int
}

func (b *ProgressBar) Init(outer ProgressBarOuter, theme guix.Theme) {
	b.outer = outer
	b.Control.Init(outer, theme)
	b.BackgroundBorderPainter.Init(outer)
	b.desiredSize = math.MaxSize
	b.target = 100

	// Interface compliance test
	_ = guix.ProgressBar(b)
}

func (b *ProgressBar) Paint(c guix.Canvas) {
	frac := math.Saturate(float32(b.progress) / float32(b.target))
	r := b.outer.Size().Rect()
	b.PaintBackground(c, r)
	b.outer.PaintProgress(c, r, frac)
	b.PaintBorder(c, r)
}

func (b *ProgressBar) PaintProgress(c guix.Canvas, r math.Rect, frac float32) {
	r.Max.X = math.Lerp(r.Min.X, r.Max.X, frac)
	c.DrawRect(r, guix.CreateBrush(guix.Gray50))
}

func (b *ProgressBar) DesiredSize(min, max math.Size) math.Size {
	return b.desiredSize.Clamp(min, max)
}

// guix.ProgressBar compliance
func (b *ProgressBar) SetDesiredSize(size math.Size) {
	b.desiredSize = size
	b.Relayout()
}

func (b *ProgressBar) SetProgress(progress int) {
	if b.progress != progress {
		b.progress = progress
		b.Redraw()
	}
}

func (b *ProgressBar) Progress() int {
	return b.progress
}

func (b *ProgressBar) SetTarget(target int) {
	if b.target != target {
		b.target = target
		b.Redraw()
	}
}

func (b *ProgressBar) Target() int {
	return b.target
}
