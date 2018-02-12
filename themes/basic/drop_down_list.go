// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins"
)

type DropDownList struct {
	mixins.DropDownList
	theme *Theme
}

func CreateDropDownList(theme *Theme) guix.DropDownList {
	l := &DropDownList{}
	l.Init(l, theme)
	l.OnGainedFocus(l.Redraw)
	l.OnLostFocus(l.Redraw)
	l.List().OnAttach(l.Redraw)
	l.List().OnDetach(l.Redraw)
	l.OnMouseEnter(func(guix.MouseEvent) {
		l.SetBorderPen(theme.DropDownListOverStyle.Pen)
	})
	l.OnMouseExit(func(guix.MouseEvent) {
		l.SetBorderPen(theme.DropDownListDefaultStyle.Pen)
	})
	l.SetPadding(math.CreateSpacing(2))
	l.SetBorderPen(theme.DropDownListDefaultStyle.Pen)
	l.SetBackgroundBrush(theme.DropDownListDefaultStyle.Brush)
	l.theme = theme
	return l
}

// mixin.List overrides
func (l *DropDownList) Paint(c guix.Canvas) {
	l.DropDownList.Paint(c)
	if l.HasFocus() || l.ListShowing() {
		r := l.Size().Rect().ContractI(1)
		c.DrawRoundedRect(r, 3.0, 3.0, 3.0, 3.0, l.theme.FocusedStyle.Pen, l.theme.FocusedStyle.Brush)
	}
}

func (l *DropDownList) DrawSelection(c guix.Canvas, r math.Rect) {
	c.DrawRoundedRect(r, 2.0, 2.0, 2.0, 2.0, l.theme.HighlightStyle.Pen, l.theme.HighlightStyle.Brush)
}
