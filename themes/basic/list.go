// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins"
)

type List struct {
	mixins.List
	theme *Theme
}

func CreateList(theme *Theme) guix.List {
	l := &List{}
	l.Init(l, theme)
	l.OnGainedFocus(l.Redraw)
	l.OnLostFocus(l.Redraw)
	l.SetPadding(math.CreateSpacing(2))
	l.SetBorderPen(guix.TransparentPen)
	l.theme = theme
	return l
}

// mixin.List overrides
func (l *List) Paint(c guix.Canvas) {
	l.List.Paint(c)
	if l.HasFocus() {
		r := l.Size().Rect().ContractI(1)
		c.DrawRoundedRect(r, 3.0, 3.0, 3.0, 3.0, l.theme.FocusedStyle.Pen, l.theme.FocusedStyle.Brush)
	}
}

func (l *List) PaintSelection(c guix.Canvas, r math.Rect) {
	c.DrawRoundedRect(r, 2.0, 2.0, 2.0, 2.0, l.theme.HighlightStyle.Pen, l.theme.HighlightStyle.Brush)
}

func (l *List) PaintMouseOverBackground(c guix.Canvas, r math.Rect) {
	c.DrawRoundedRect(r, 2.0, 2.0, 2.0, 2.0, guix.TransparentPen, guix.CreateBrush(guix.Gray15))
}
