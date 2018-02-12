// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/mixins/base"
	"github.com/vcaesar/guix/mixins/parts"
)

type LinearLayoutOuter interface {
	base.ContainerOuter
}

type LinearLayout struct {
	base.Container
	parts.LinearLayout
	parts.BackgroundBorderPainter
}

func (l *LinearLayout) Init(outer LinearLayoutOuter, theme guix.Theme) {
	l.Container.Init(outer, theme)
	l.LinearLayout.Init(outer)
	l.BackgroundBorderPainter.Init(outer)
	l.SetMouseEventTarget(true)
	l.SetBackgroundBrush(guix.TransparentBrush)
	l.SetBorderPen(guix.TransparentPen)

	// Interface compliance test
	_ = guix.LinearLayout(l)
}

func (l *LinearLayout) Paint(c guix.Canvas) {
	r := l.Size().Rect()
	l.BackgroundBorderPainter.PaintBackground(c, r)
	l.PaintChildren.Paint(c)
	l.BackgroundBorderPainter.PaintBorder(c, r)
}
