// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"strings"

	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/base"
)

type LabelOuter interface {
	base.ControlOuter
}

type Label struct {
	base.Control

	outer               LabelOuter
	font                guix.Font
	color               guix.Color
	horizontalAlignment guix.HorizontalAlignment
	verticalAlignment   guix.VerticalAlignment
	multiline           bool
	text                string
}

func (l *Label) Init(outer LabelOuter, theme guix.Theme, font guix.Font, color guix.Color) {
	if font == nil {
		panic("Cannot create a label with a nil font")
	}
	l.Control.Init(outer, theme)
	l.outer = outer
	l.font = font
	l.color = color
	l.horizontalAlignment = guix.AlignLeft
	l.verticalAlignment = guix.AlignMiddle
	// Interface compliance test
	_ = guix.Label(l)
}

func (l *Label) Text() string {
	return l.text
}

func (l *Label) SetText(text string) {
	if l.text != text {
		l.text = text
		l.outer.Relayout()
	}
}

func (l *Label) Font() guix.Font {
	return l.font
}

func (l *Label) SetFont(font guix.Font) {
	if l.font != font {
		l.font = font
		l.Relayout()
	}
}

func (l *Label) Color() guix.Color {
	return l.color
}

func (l *Label) SetColor(color guix.Color) {
	if l.color != color {
		l.color = color
		l.outer.Redraw()
	}
}

func (l *Label) Multiline() bool {
	return l.multiline
}

func (l *Label) SetMultiline(multiline bool) {
	if l.multiline != multiline {
		l.multiline = multiline
		l.outer.Relayout()
	}
}

func (l *Label) DesiredSize(min, max math.Size) math.Size {
	t := l.text
	if !l.multiline {
		t = strings.Replace(t, "\n", " ", -1)
	}
	s := l.font.Measure(&guix.TextBlock{Runes: []rune(t)})
	return s.Clamp(min, max)
}

func (l *Label) SetHorizontalAlignment(horizontalAlignment guix.HorizontalAlignment) {
	if l.horizontalAlignment != horizontalAlignment {
		l.horizontalAlignment = horizontalAlignment
		l.Redraw()
	}
}

func (l *Label) HorizontalAlignment() guix.HorizontalAlignment {
	return l.horizontalAlignment
}

func (l *Label) SetVerticalAlignment(verticalAlignment guix.VerticalAlignment) {
	if l.verticalAlignment != verticalAlignment {
		l.verticalAlignment = verticalAlignment
		l.Redraw()
	}
}

func (l *Label) VerticalAlignment() guix.VerticalAlignment {
	return l.verticalAlignment
}

// parts.DrawPaint overrides
func (l *Label) Paint(c guix.Canvas) {
	r := l.outer.Size().Rect()
	t := l.text
	if !l.multiline {
		t = strings.Replace(t, "\n", " ", -1)
	}

	runes := []rune(t)
	offsets := l.font.Layout(&guix.TextBlock{
		Runes:     runes,
		AlignRect: r,
		H:         l.horizontalAlignment,
		V:         l.verticalAlignment,
	})
	c.DrawRunes(l.font, runes, offsets, l.color)
}
