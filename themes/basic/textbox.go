// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins"
)

type TextBox struct {
	mixins.TextBox
	theme *Theme
}

func CreateTextBox(theme *Theme) guix.TextBox {
	t := &TextBox{}
	t.Init(t, theme.Driver(), theme, theme.DefaultFont())
	t.SetTextColor(theme.TextBoxDefaultStyle.FontColor)
	t.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	t.SetPadding(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	t.SetBackgroundBrush(theme.TextBoxDefaultStyle.Brush)
	t.SetBorderPen(theme.TextBoxDefaultStyle.Pen)
	t.OnMouseEnter(func(guix.MouseEvent) {
		t.SetBackgroundBrush(theme.TextBoxOverStyle.Brush)
		t.SetBorderPen(theme.TextBoxOverStyle.Pen)
	})
	t.OnMouseExit(func(guix.MouseEvent) {
		t.SetBackgroundBrush(theme.TextBoxDefaultStyle.Brush)
		t.SetBorderPen(theme.TextBoxDefaultStyle.Pen)
	})

	t.theme = theme

	return t
}

// mixins.TextBox overrides
func (t *TextBox) Paint(c guix.Canvas) {
	t.TextBox.Paint(c)

	if t.HasFocus() {
		r := t.Size().Rect()
		s := t.theme.FocusedStyle
		c.DrawRoundedRect(r, 3, 3, 3, 3, s.Pen, s.Brush)
	}
}
