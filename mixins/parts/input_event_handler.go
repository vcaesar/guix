// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parts

import (
	"github.com/vcaesar/guix"
)

type InputEventHandlerOuter interface{}

type InputEventHandler struct {
	outer         InputEventHandlerOuter
	isMouseOver   bool
	isMouseDown   map[guix.MouseButton]bool
	onClick       guix.Event
	onDoubleClick guix.Event
	onKeyPress    guix.Event
	onKeyStroke   guix.Event
	onMouseMove   guix.Event
	onMouseEnter  guix.Event
	onMouseExit   guix.Event
	onMouseDown   guix.Event
	onMouseUp     guix.Event
	onMouseScroll guix.Event
	onKeyDown     guix.Event
	onKeyUp       guix.Event
	onKeyRepeat   guix.Event
}

func (m *InputEventHandler) getOnClick() guix.Event {
	if m.onClick == nil {
		m.onClick = guix.CreateEvent(m.Click)
	}
	return m.onClick
}

func (m *InputEventHandler) getOnDoubleClick() guix.Event {
	if m.onDoubleClick == nil {
		m.onDoubleClick = guix.CreateEvent(m.DoubleClick)
	}
	return m.onDoubleClick
}

func (m *InputEventHandler) getOnKeyPress() guix.Event {
	if m.onKeyPress == nil {
		m.onKeyPress = guix.CreateEvent(m.KeyPress)
	}
	return m.onKeyPress
}

func (m *InputEventHandler) getOnKeyStroke() guix.Event {
	if m.onKeyStroke == nil {
		m.onKeyStroke = guix.CreateEvent(m.KeyStroke)
	}
	return m.onKeyStroke
}

func (m *InputEventHandler) getOnMouseMove() guix.Event {
	if m.onMouseMove == nil {
		m.onMouseMove = guix.CreateEvent(m.MouseMove)
	}
	return m.onMouseMove
}

func (m *InputEventHandler) getOnMouseEnter() guix.Event {
	if m.onMouseEnter == nil {
		m.onMouseEnter = guix.CreateEvent(m.MouseEnter)
	}
	return m.onMouseEnter
}

func (m *InputEventHandler) getOnMouseExit() guix.Event {
	if m.onMouseExit == nil {
		m.onMouseExit = guix.CreateEvent(m.MouseExit)
	}
	return m.onMouseExit
}

func (m *InputEventHandler) getOnMouseDown() guix.Event {
	if m.onMouseDown == nil {
		m.onMouseDown = guix.CreateEvent(m.MouseDown)
	}
	return m.onMouseDown
}

func (m *InputEventHandler) getOnMouseUp() guix.Event {
	if m.onMouseUp == nil {
		m.onMouseUp = guix.CreateEvent(m.MouseUp)
	}
	return m.onMouseUp
}

func (m *InputEventHandler) getOnMouseScroll() guix.Event {
	if m.onMouseScroll == nil {
		m.onMouseScroll = guix.CreateEvent(m.MouseScroll)
	}
	return m.onMouseScroll
}

func (m *InputEventHandler) getOnKeyDown() guix.Event {
	if m.onKeyDown == nil {
		m.onKeyDown = guix.CreateEvent(m.KeyDown)
	}
	return m.onKeyDown
}

func (m *InputEventHandler) getOnKeyUp() guix.Event {
	if m.onKeyUp == nil {
		m.onKeyUp = guix.CreateEvent(m.KeyUp)
	}
	return m.onKeyUp
}

func (m *InputEventHandler) getOnKeyRepeat() guix.Event {
	if m.onKeyRepeat == nil {
		m.onKeyRepeat = guix.CreateEvent(m.KeyRepeat)
	}
	return m.onKeyRepeat
}

func (m *InputEventHandler) Init(outer InputEventHandlerOuter) {
	m.outer = outer
	m.isMouseDown = make(map[guix.MouseButton]bool)
}

func (m *InputEventHandler) Click(ev guix.MouseEvent) (consume bool) {
	m.getOnClick().Fire(ev)
	return false
}

func (m *InputEventHandler) DoubleClick(ev guix.MouseEvent) (consume bool) {
	m.getOnDoubleClick().Fire(ev)
	return false
}

func (m *InputEventHandler) KeyPress(ev guix.KeyboardEvent) (consume bool) {
	m.getOnKeyPress().Fire(ev)
	return false
}

func (m *InputEventHandler) KeyStroke(ev guix.KeyStrokeEvent) (consume bool) {
	m.getOnKeyStroke().Fire(ev)
	return false
}

func (m *InputEventHandler) MouseScroll(ev guix.MouseEvent) (consume bool) {
	m.getOnMouseScroll().Fire(ev)
	return false
}

func (m *InputEventHandler) MouseMove(ev guix.MouseEvent) {
	m.getOnMouseMove().Fire(ev)
}

func (m *InputEventHandler) MouseEnter(ev guix.MouseEvent) {
	m.isMouseOver = true
	m.getOnMouseEnter().Fire(ev)
}

func (m *InputEventHandler) MouseExit(ev guix.MouseEvent) {
	m.isMouseOver = false
	m.getOnMouseExit().Fire(ev)
}

func (m *InputEventHandler) MouseDown(ev guix.MouseEvent) {
	m.isMouseDown[ev.Button] = true
	m.getOnMouseDown().Fire(ev)
}

func (m *InputEventHandler) MouseUp(ev guix.MouseEvent) {
	m.isMouseDown[ev.Button] = false
	m.getOnMouseUp().Fire(ev)
}

func (m *InputEventHandler) KeyDown(ev guix.KeyboardEvent) {
	m.getOnKeyDown().Fire(ev)
}

func (m *InputEventHandler) KeyUp(ev guix.KeyboardEvent) {
	m.getOnKeyUp().Fire(ev)
}

func (m *InputEventHandler) KeyRepeat(ev guix.KeyboardEvent) {
	m.getOnKeyRepeat().Fire(ev)
}

func (m *InputEventHandler) OnClick(f func(guix.MouseEvent)) guix.EventSubscription {
	return m.getOnClick().Listen(f)
}

func (m *InputEventHandler) OnDoubleClick(f func(guix.MouseEvent)) guix.EventSubscription {
	return m.getOnDoubleClick().Listen(f)
}

func (m *InputEventHandler) OnKeyPress(f func(guix.KeyboardEvent)) guix.EventSubscription {
	return m.getOnKeyPress().Listen(f)
}

func (m *InputEventHandler) OnKeyStroke(f func(guix.KeyStrokeEvent)) guix.EventSubscription {
	return m.getOnKeyStroke().Listen(f)
}

func (m *InputEventHandler) OnMouseMove(f func(guix.MouseEvent)) guix.EventSubscription {
	return m.getOnMouseMove().Listen(f)
}

func (m *InputEventHandler) OnMouseEnter(f func(guix.MouseEvent)) guix.EventSubscription {
	return m.getOnMouseEnter().Listen(f)
}

func (m *InputEventHandler) OnMouseExit(f func(guix.MouseEvent)) guix.EventSubscription {
	return m.getOnMouseExit().Listen(f)
}

func (m *InputEventHandler) OnMouseDown(f func(guix.MouseEvent)) guix.EventSubscription {
	return m.getOnMouseDown().Listen(f)
}

func (m *InputEventHandler) OnMouseUp(f func(guix.MouseEvent)) guix.EventSubscription {
	return m.getOnMouseUp().Listen(f)
}

func (m *InputEventHandler) OnMouseScroll(f func(guix.MouseEvent)) guix.EventSubscription {
	return m.getOnMouseScroll().Listen(f)
}

func (m *InputEventHandler) OnKeyDown(f func(guix.KeyboardEvent)) guix.EventSubscription {
	return m.getOnKeyDown().Listen(f)
}

func (m *InputEventHandler) OnKeyUp(f func(guix.KeyboardEvent)) guix.EventSubscription {
	return m.getOnKeyUp().Listen(f)
}

func (m *InputEventHandler) OnKeyRepeat(f func(guix.KeyboardEvent)) guix.EventSubscription {
	return m.getOnKeyRepeat().Listen(f)
}

func (m *InputEventHandler) IsMouseOver() bool {
	return m.isMouseOver
}

func (m *InputEventHandler) IsMouseDown(button guix.MouseButton) bool {
	return m.isMouseDown[button]
}
