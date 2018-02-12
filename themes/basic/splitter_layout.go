// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins"
)

type SplitterLayout struct {
	mixins.SplitterLayout
	theme *Theme
}

func CreateSplitterLayout(theme *Theme) guix.SplitterLayout {
	l := &SplitterLayout{}
	l.theme = theme
	l.Init(l, theme)
	return l
}

// mixins.SplitterLayout overrides
func (l *SplitterLayout) CreateSplitterBar() guix.Control {
	b := &mixins.SplitterBar{}
	b.Init(b, l.theme)
	b.SetBackgroundColor(l.theme.SplitterBarDefaultStyle.Brush.Color)
	b.SetForegroundColor(l.theme.SplitterBarDefaultStyle.Pen.Color)
	b.OnSplitterDragged(func(wndPnt math.Point) { l.SplitterDragged(b, wndPnt) })
	updateForegroundColor := func() {
		switch {
		case b.IsDragging():
			b.SetForegroundColor(l.theme.HighlightStyle.Pen.Color)
		case b.IsMouseOver():
			b.SetForegroundColor(l.theme.SplitterBarOverStyle.Pen.Color)
		default:
			b.SetForegroundColor(l.theme.SplitterBarDefaultStyle.Pen.Color)
		}
		b.Redraw()
	}
	b.OnDragStart(func(guix.MouseEvent) { updateForegroundColor() })
	b.OnDragEnd(func(guix.MouseEvent) { updateForegroundColor() })
	b.OnDragStart(func(guix.MouseEvent) { updateForegroundColor() })
	b.OnMouseEnter(func(guix.MouseEvent) { updateForegroundColor() })
	b.OnMouseExit(func(guix.MouseEvent) { updateForegroundColor() })
	return b
}
