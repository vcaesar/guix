// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parts

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/outer"
)

type LinearLayoutOuter interface {
	guix.Container
	outer.Sized
}

type LinearLayout struct {
	outer               LinearLayoutOuter
	direction           guix.Direction
	sizeMode            guix.SizeMode
	horizontalAlignment guix.HorizontalAlignment
	verticalAlignment   guix.VerticalAlignment
}

func (l *LinearLayout) Init(outer LinearLayoutOuter) {
	l.outer = outer
}

func (l *LinearLayout) LayoutChildren() {
	s := l.outer.Size().Contract(l.outer.Padding())
	o := l.outer.Padding().LT()
	children := l.outer.Children()
	major := 0
	if l.direction.RightToLeft() || l.direction.BottomToTop() {
		if l.direction.RightToLeft() {
			major = s.W
		} else {
			major = s.H
		}
	}
	for _, c := range children {
		cm := c.Control.Margin()
		cs := c.Control.DesiredSize(math.ZeroSize, s.Contract(cm).Max(math.ZeroSize))
		c.Control.SetSize(cs)

		// Calculate minor-axis alignment
		var minor int
		switch l.direction.Orientation() {
		case guix.Horizontal:
			switch l.verticalAlignment {
			case guix.AlignTop:
				minor = cm.T
			case guix.AlignMiddle:
				minor = (s.H - cs.H) / 2
			case guix.AlignBottom:
				minor = s.H - cs.H
			}
		case guix.Vertical:
			switch l.horizontalAlignment {
			case guix.AlignLeft:
				minor = cm.L
			case guix.AlignCenter:
				minor = (s.W - cs.W) / 2
			case guix.AlignRight:
				minor = s.W - cs.W
			}
		}

		// Peform layout
		switch l.direction {
		case guix.LeftToRight:
			major += cm.L
			c.Offset = math.Point{X: major, Y: minor}.Add(o)
			major += cs.W
			major += cm.R
			s.W -= cs.W + cm.W()
		case guix.RightToLeft:
			major -= cm.R
			c.Offset = math.Point{X: major - cs.W, Y: minor}.Add(o)
			major -= cs.W
			major -= cm.L
			s.W -= cs.W + cm.W()
		case guix.TopToBottom:
			major += cm.T
			c.Offset = math.Point{X: minor, Y: major}.Add(o)
			major += cs.H
			major += cm.B
			s.H -= cs.H + cm.H()
		case guix.BottomToTop:
			major -= cm.B
			c.Offset = math.Point{X: minor, Y: major - cs.H}.Add(o)
			major -= cs.H
			major -= cm.T
			s.H -= cs.H + cm.H()
		}
	}
}

func (l *LinearLayout) DesiredSize(min, max math.Size) math.Size {
	if l.sizeMode.Fill() {
		return max
	}

	bounds := min.Rect()
	children := l.outer.Children()

	horizontal := l.direction.Orientation().Horizontal()
	offset := math.Point{X: 0, Y: 0}
	for _, c := range children {
		cs := c.Control.DesiredSize(math.ZeroSize, max)
		cm := c.Control.Margin()
		cb := cs.Expand(cm).Rect().Offset(offset)
		if horizontal {
			offset.X += cb.W()
		} else {
			offset.Y += cb.H()
		}
		bounds = bounds.Union(cb)
	}

	return bounds.Size().Expand(l.outer.Padding()).Clamp(min, max)
}

func (l *LinearLayout) Direction() guix.Direction {
	return l.direction
}

func (l *LinearLayout) SetDirection(d guix.Direction) {
	if l.direction != d {
		l.direction = d
		l.outer.Relayout()
	}
}

func (l *LinearLayout) SizeMode() guix.SizeMode {
	return l.sizeMode
}

func (l *LinearLayout) SetSizeMode(mode guix.SizeMode) {
	if l.sizeMode != mode {
		l.sizeMode = mode
		l.outer.Relayout()
	}
}

func (l *LinearLayout) HorizontalAlignment() guix.HorizontalAlignment {
	return l.horizontalAlignment
}

func (l *LinearLayout) SetHorizontalAlignment(alignment guix.HorizontalAlignment) {
	if l.horizontalAlignment != alignment {
		l.horizontalAlignment = alignment
		l.outer.Relayout()
	}
}

func (l *LinearLayout) VerticalAlignment() guix.VerticalAlignment {
	return l.verticalAlignment
}

func (l *LinearLayout) SetVerticalAlignment(alignment guix.VerticalAlignment) {
	if l.verticalAlignment != alignment {
		l.verticalAlignment = alignment
		l.outer.Relayout()
	}
}
