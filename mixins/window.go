// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/outer"
	"github.com/vcaesar/guix/mixins/parts"
)

type WindowOuter interface {
	guix.Window
	outer.Attachable
	outer.IsVisibler
	outer.LayoutChildren
	outer.PaintChilder
	outer.Painter
	outer.Parenter
	outer.Sized
}

type Window struct {
	parts.Attachable
	parts.BackgroundBorderPainter
	parts.Container
	parts.Paddable
	parts.PaintChildren

	driver             guix.Driver
	outer              WindowOuter
	viewport           guix.Viewport
	windowedSize       math.Size
	mouseController    *guix.MouseController
	keyboardController *guix.KeyboardController
	focusController    *guix.FocusController
	layoutPending      bool
	drawPending        bool
	updatePending      bool
	onClose            guix.Event // Raised by viewport
	onResize           guix.Event // Raised by viewport
	onMouseMove        guix.Event // Raised by viewport
	onMouseEnter       guix.Event // Raised by viewport
	onMouseExit        guix.Event // Raised by viewport
	onMouseDown        guix.Event // Raised by viewport
	onMouseUp          guix.Event // Raised by viewport
	onMouseScroll      guix.Event // Raised by viewport
	onKeyDown          guix.Event // Raised by viewport
	onKeyUp            guix.Event // Raised by viewport
	onKeyRepeat        guix.Event // Raised by viewport
	onKeyStroke        guix.Event // Raised by viewport

	onClick       guix.Event // Raised by MouseController
	onDoubleClick guix.Event // Raised by MouseController

	viewportSubscriptions []guix.EventSubscription
}

func (w *Window) requestUpdate() {
	if !w.updatePending {
		w.updatePending = true
		w.driver.Call(w.update)
	}
}

func (w *Window) update() {
	if !w.Attached() {
		// Window was detached between requestUpdate() and update()
		w.updatePending = false
		w.layoutPending = false
		w.drawPending = false
		return
	}
	w.updatePending = false
	if w.layoutPending {
		w.layoutPending = false
		w.drawPending = true
		w.outer.LayoutChildren()
	}
	if w.drawPending {
		w.drawPending = false
		w.Draw()
	}
}

func (w *Window) Init(outer WindowOuter, driver guix.Driver, width, height int, title string) {
	w.Attachable.Init(outer)
	w.BackgroundBorderPainter.Init(outer)
	w.Container.Init(outer)
	w.Paddable.Init(outer)
	w.PaintChildren.Init(outer)
	w.outer = outer
	w.driver = driver

	w.onClose = guix.CreateEvent(func() {})
	w.onResize = guix.CreateEvent(func() {})
	w.onMouseMove = guix.CreateEvent(func(guix.MouseEvent) {})
	w.onMouseEnter = guix.CreateEvent(func(guix.MouseEvent) {})
	w.onMouseExit = guix.CreateEvent(func(guix.MouseEvent) {})
	w.onMouseDown = guix.CreateEvent(func(guix.MouseEvent) {})
	w.onMouseUp = guix.CreateEvent(func(guix.MouseEvent) {})
	w.onMouseScroll = guix.CreateEvent(func(guix.MouseEvent) {})
	w.onKeyDown = guix.CreateEvent(func(guix.KeyboardEvent) {})
	w.onKeyUp = guix.CreateEvent(func(guix.KeyboardEvent) {})
	w.onKeyRepeat = guix.CreateEvent(func(guix.KeyboardEvent) {})
	w.onKeyStroke = guix.CreateEvent(func(guix.KeyStrokeEvent) {})

	w.onClick = guix.CreateEvent(func(guix.MouseEvent) {})
	w.onDoubleClick = guix.CreateEvent(func(guix.MouseEvent) {})

	w.focusController = guix.CreateFocusController(outer)
	w.mouseController = guix.CreateMouseController(outer, w.focusController)
	w.keyboardController = guix.CreateKeyboardController(outer)

	w.onResize.Listen(func() {
		w.outer.LayoutChildren()
		w.Draw()
	})

	w.SetBorderPen(guix.TransparentPen)

	w.setViewport(driver.CreateWindowedViewport(width, height, title))

	// Window starts shown
	w.Attach()

	// Interface compliance test
	_ = guix.Window(w)
}

func (w *Window) Draw() guix.Canvas {
	if s := w.viewport.SizeDips(); s != math.ZeroSize {
		c := w.driver.CreateCanvas(s)
		w.outer.Paint(c)
		c.Complete()
		w.viewport.SetCanvas(c)
		return c
	} else {
		return nil
	}
}

func (w *Window) Paint(c guix.Canvas) {
	w.PaintBackground(c, c.Size().Rect())
	w.PaintChildren.Paint(c)
	w.PaintBorder(c, c.Size().Rect())
}

func (w *Window) LayoutChildren() {
	s := w.Size().Contract(w.Padding()).Max(math.ZeroSize)
	o := w.Padding().LT()
	for _, c := range w.outer.Children() {
		c.Layout(c.Control.DesiredSize(math.ZeroSize, s).Rect().Offset(o))
	}
}

func (w *Window) Size() math.Size {
	return w.viewport.SizeDips()
}

func (w *Window) SetSize(size math.Size) {
	w.viewport.SetSizeDips(size)
}

func (w *Window) Parent() guix.Parent {
	return nil
}

func (w *Window) Viewport() guix.Viewport {
	return w.viewport
}

func (w *Window) Title() string {
	return w.viewport.Title()
}

func (w *Window) SetTitle(t string) {
	w.viewport.SetTitle(t)
}

func (w *Window) Scale() float32 {
	return w.viewport.Scale()
}

func (w *Window) SetScale(scale float32) {
	w.viewport.SetScale(scale)
}

func (w *Window) Position() math.Point {
	return w.viewport.Position()
}

func (w *Window) SetPosition(pos math.Point) {
	w.viewport.SetPosition(pos)
}

func (w *Window) Fullscreen() bool {
	return w.viewport.Fullscreen()
}

func (w *Window) SetFullscreen(fullscreen bool) {
	title := w.viewport.Title()
	if fullscreen != w.Fullscreen() {
		old := w.viewport
		if fullscreen {
			w.windowedSize = old.SizeDips()
			w.setViewport(w.driver.CreateFullscreenViewport(0, 0, title))
		} else {
			width, height := w.windowedSize.WH()
			w.setViewport(w.driver.CreateWindowedViewport(width, height, title))
		}
		old.Close()
	}
}

func (w *Window) Show() {
	w.Attach()
	w.viewport.Show()
}

func (w *Window) Hide() {
	w.Detach()
	w.viewport.Hide()
}

func (w *Window) Close() {
	w.Detach()
	w.viewport.Close()
}

func (w *Window) Focus() guix.Focusable {
	return w.focusController.Focus()
}

func (w *Window) SetFocus(c guix.Control) bool {
	fc := w.focusController
	if c == nil {
		fc.SetFocus(nil)
		return true
	}
	if f := fc.Focusable(c); f != nil {
		fc.SetFocus(f)
		return true
	}
	return false
}

func (w *Window) IsVisible() bool {
	return true
}

func (w *Window) OnClose(f func()) guix.EventSubscription {
	return w.onClose.Listen(f)
}

func (w *Window) OnResize(f func()) guix.EventSubscription {
	return w.onResize.Listen(f)
}

func (w *Window) OnClick(f func(guix.MouseEvent)) guix.EventSubscription {
	return w.onClick.Listen(f)
}

func (w *Window) OnDoubleClick(f func(guix.MouseEvent)) guix.EventSubscription {
	return w.onDoubleClick.Listen(f)
}

func (w *Window) OnMouseMove(f func(guix.MouseEvent)) guix.EventSubscription {
	return w.onMouseMove.Listen(func(ev guix.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnMouseEnter(f func(guix.MouseEvent)) guix.EventSubscription {
	return w.onMouseEnter.Listen(func(ev guix.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnMouseExit(f func(guix.MouseEvent)) guix.EventSubscription {
	return w.onMouseExit.Listen(func(ev guix.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnMouseDown(f func(guix.MouseEvent)) guix.EventSubscription {
	return w.onMouseDown.Listen(func(ev guix.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnMouseUp(f func(guix.MouseEvent)) guix.EventSubscription {
	return w.onMouseUp.Listen(func(ev guix.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnMouseScroll(f func(guix.MouseEvent)) guix.EventSubscription {
	return w.onMouseScroll.Listen(func(ev guix.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnKeyDown(f func(guix.KeyboardEvent)) guix.EventSubscription {
	return w.onKeyDown.Listen(f)
}

func (w *Window) OnKeyUp(f func(guix.KeyboardEvent)) guix.EventSubscription {
	return w.onKeyUp.Listen(f)
}

func (w *Window) OnKeyRepeat(f func(guix.KeyboardEvent)) guix.EventSubscription {
	return w.onKeyRepeat.Listen(f)
}

func (w *Window) OnKeyStroke(f func(guix.KeyStrokeEvent)) guix.EventSubscription {
	return w.onKeyStroke.Listen(f)
}

func (w *Window) Relayout() {
	w.layoutPending = true
	w.requestUpdate()
}

func (w *Window) Redraw() {
	w.drawPending = true
	w.requestUpdate()
}

func (w *Window) Click(ev guix.MouseEvent) {
	w.onClick.Fire(ev)
}

func (w *Window) DoubleClick(ev guix.MouseEvent) {
	w.onDoubleClick.Fire(ev)
}

func (w *Window) KeyPress(ev guix.KeyboardEvent) {
	if ev.Key == guix.KeyTab {
		if ev.Modifier&guix.ModShift != 0 {
			w.focusController.FocusPrev()
		} else {
			w.focusController.FocusNext()
		}
	}
}
func (w *Window) KeyStroke(guix.KeyStrokeEvent) {}

func (w *Window) setViewport(v guix.Viewport) {
	for _, s := range w.viewportSubscriptions {
		s.Unlisten()
	}
	w.viewport = v
	w.viewportSubscriptions = []guix.EventSubscription{
		v.OnClose(func() { w.onClose.Fire() }),
		v.OnResize(func() { w.onResize.Fire() }),
		v.OnMouseMove(func(ev guix.MouseEvent) { w.onMouseMove.Fire(ev) }),
		v.OnMouseEnter(func(ev guix.MouseEvent) { w.onMouseEnter.Fire(ev) }),
		v.OnMouseExit(func(ev guix.MouseEvent) { w.onMouseExit.Fire(ev) }),
		v.OnMouseDown(func(ev guix.MouseEvent) { w.onMouseDown.Fire(ev) }),
		v.OnMouseUp(func(ev guix.MouseEvent) { w.onMouseUp.Fire(ev) }),
		v.OnMouseScroll(func(ev guix.MouseEvent) { w.onMouseScroll.Fire(ev) }),
		v.OnKeyDown(func(ev guix.KeyboardEvent) { w.onKeyDown.Fire(ev) }),
		v.OnKeyUp(func(ev guix.KeyboardEvent) { w.onKeyUp.Fire(ev) }),
		v.OnKeyRepeat(func(ev guix.KeyboardEvent) { w.onKeyRepeat.Fire(ev) }),
		v.OnKeyStroke(func(ev guix.KeyStrokeEvent) { w.onKeyStroke.Fire(ev) }),
	}
	w.Relayout()
}
