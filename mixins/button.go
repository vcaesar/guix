// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/parts"
)

type ButtonOuter interface {
	LinearLayoutOuter
	IsChecked() bool
	SetChecked(bool)
}

type Button struct {
	LinearLayout
	parts.Focusable

	outer      ButtonOuter
	theme      guix.Theme
	label      guix.Label
	buttonType guix.ButtonType
	checked    bool
}

func (b *Button) Init(outer ButtonOuter, theme guix.Theme) {
	b.LinearLayout.Init(outer, theme)
	b.Focusable.Init(outer)

	b.buttonType = guix.PushButton
	b.theme = theme
	b.outer = outer

	// Interface compliance test
	_ = guix.Button(b)
}

func (b *Button) Label() guix.Label {
	return b.label
}

func (b *Button) Text() string {
	if b.label != nil {
		return b.label.Text()
	} else {
		return ""
	}
}

func (b *Button) SetText(text string) {
	if b.Text() == text {
		return
	}
	if text == "" {
		if b.label != nil {
			b.RemoveChild(b.label)
			b.label = nil
		}
	} else {
		if b.label == nil {
			b.label = b.theme.CreateLabel()
			b.label.SetMargin(math.ZeroSpacing)
			b.AddChild(b.label)
		}
		b.label.SetText(text)
	}
}

func (b *Button) Type() guix.ButtonType {
	return b.buttonType
}

func (b *Button) SetType(buttonType guix.ButtonType) {
	if buttonType != b.buttonType {
		b.buttonType = buttonType
		b.outer.Redraw()
	}
}

func (b *Button) IsChecked() bool {
	return b.checked
}

func (b *Button) SetChecked(checked bool) {
	if checked != b.checked {
		b.checked = checked
		b.outer.Redraw()
	}
}

// InputEventHandler override
func (b *Button) Click(ev guix.MouseEvent) (consume bool) {
	if ev.Button == guix.MouseButtonLeft {
		if b.buttonType == guix.ToggleButton {
			b.outer.SetChecked(!b.outer.IsChecked())
		}
		b.LinearLayout.Click(ev)
		return true
	}
	return b.LinearLayout.Click(ev)
}

func (b *Button) KeyPress(ev guix.KeyboardEvent) (consume bool) {
	consume = b.LinearLayout.KeyPress(ev)
	if ev.Key == guix.KeySpace || ev.Key == guix.KeyEnter {
		me := guix.MouseEvent{
			Button: guix.MouseButtonLeft,
		}
		return b.Click(me)
	}
	return
}
