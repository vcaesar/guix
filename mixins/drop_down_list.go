// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/base"
	"github.com/vcaesar/guix/mixins/parts"
)

type DropDownListOuter interface {
	base.ContainerOuter
}

type DropDownList struct {
	base.Container
	parts.BackgroundBorderPainter
	parts.Focusable

	outer DropDownListOuter

	theme       guix.Theme
	list        guix.List
	listShowing bool
	itemSize    math.Size
	overlay     guix.BubbleOverlay
	selected    *guix.Child
	onShowList  guix.Event
	onHideList  guix.Event
}

func (l *DropDownList) Init(outer DropDownListOuter, theme guix.Theme) {
	l.outer = outer
	l.Container.Init(outer, theme)
	l.BackgroundBorderPainter.Init(outer)
	l.Focusable.Init(outer)

	l.theme = theme
	l.list = theme.CreateList()
	l.list.OnSelectionChanged(func(item guix.AdapterItem) {
		l.outer.RemoveAll()
		adapter := l.list.Adapter()
		if item != nil && adapter != nil {
			l.selected = l.AddChild(adapter.Create(l.theme, adapter.ItemIndex(item)))
		} else {
			l.selected = nil
		}
		l.Relayout()
	})
	l.list.OnItemClicked(func(guix.MouseEvent, guix.AdapterItem) {
		l.HideList()
	})
	l.list.OnKeyPress(func(ev guix.KeyboardEvent) {
		switch ev.Key {
		case guix.KeyEnter, guix.KeyEscape:
			l.HideList()
		}
	})
	l.list.OnLostFocus(l.HideList)
	l.OnDetach(l.HideList)
	l.SetMouseEventTarget(true)

	// Interface compliance test
	_ = guix.DropDownList(l)
}

func (l *DropDownList) LayoutChildren() {
	if !l.RelayoutSuspended() {
		// Disable relayout on AddChild / RemoveChild as we're performing layout here.
		l.SetRelayoutSuspended(true)
		defer l.SetRelayoutSuspended(false)
	}

	if l.selected != nil {
		s := l.outer.Size().Contract(l.Padding()).Max(math.ZeroSize)
		o := l.Padding().LT()
		l.selected.Layout(s.Rect().Offset(o))
	}
}

func (l *DropDownList) DesiredSize(min, max math.Size) math.Size {
	if l.selected != nil {
		return l.selected.Control.DesiredSize(min, max).Expand(l.outer.Padding()).Clamp(min, max)
	} else {
		return l.itemSize.Expand(l.outer.Padding()).Clamp(min, max)
	}
}

func (l *DropDownList) DataReplaced() {
	adapter := l.list.Adapter()
	itemSize := adapter.Size(l.theme)
	l.itemSize = itemSize
	l.outer.Relayout()
}

func (l *DropDownList) ListShowing() bool {
	return l.listShowing
}

func (l *DropDownList) ShowList() bool {
	if l.listShowing || l.overlay == nil {
		return false
	}
	l.listShowing = true
	s := l.Size()
	at := math.Point{X: s.W / 2, Y: s.H}
	l.overlay.Show(l.list, guix.TransformCoordinate(at, l.outer, l.overlay))
	guix.SetFocus(l.list)
	if l.onShowList != nil {
		l.onShowList.Fire()
	}
	return true
}

func (l *DropDownList) HideList() {
	if l.listShowing {
		l.listShowing = false
		l.overlay.Hide()
		if l.Attached() {
			guix.SetFocus(l)
		}
		if l.onHideList != nil {
			l.onHideList.Fire()
		}
	}
}

func (l *DropDownList) List() guix.List {
	return l.list
}

// InputEventHandler override
func (l *DropDownList) Click(ev guix.MouseEvent) (consume bool) {
	l.InputEventHandler.Click(ev)
	if l.ListShowing() {
		l.HideList()
	} else {
		l.ShowList()
	}
	return true
}

// guix.DropDownList compliance
func (l *DropDownList) SetBubbleOverlay(overlay guix.BubbleOverlay) {
	l.overlay = overlay
}

func (l *DropDownList) BubbleOverlay() guix.BubbleOverlay {
	return l.overlay
}

func (l *DropDownList) Adapter() guix.ListAdapter {
	return l.list.Adapter()
}

func (l *DropDownList) SetAdapter(adapter guix.ListAdapter) {
	if l.list.Adapter() != adapter {
		l.list.SetAdapter(adapter)
		if adapter != nil {
			adapter.OnDataChanged(func(bool) { l.DataReplaced() })
			adapter.OnDataReplaced(l.DataReplaced)
		}
		// TODO: Unlisten
		l.DataReplaced()
	}
}

func (l *DropDownList) Selected() guix.AdapterItem {
	return l.list.Selected()
}

func (l *DropDownList) Select(item guix.AdapterItem) {
	if l.list.Selected() != item {
		l.list.Select(item)
		l.LayoutChildren()
	}
}

func (l *DropDownList) OnSelectionChanged(f func(guix.AdapterItem)) guix.EventSubscription {
	return l.list.OnSelectionChanged(f)
}

func (l *DropDownList) OnShowList(f func()) guix.EventSubscription {
	if l.onShowList == nil {
		l.onShowList = guix.CreateEvent(f)
	}
	return l.onShowList.Listen(f)
}

func (l *DropDownList) OnHideList(f func()) guix.EventSubscription {
	if l.onHideList == nil {
		l.onHideList = guix.CreateEvent(f)
	}
	return l.onHideList.Listen(f)
}

// InputEventHandler overrides
func (l *DropDownList) KeyPress(ev guix.KeyboardEvent) (consume bool) {
	if ev.Key == guix.KeySpace || ev.Key == guix.KeyEnter {
		me := guix.MouseEvent{
			Button: guix.MouseButtonLeft,
		}
		return l.Click(me)
	}
	return l.InputEventHandler.KeyPress(ev)
}

// parts.Container overrides
func (l *DropDownList) Paint(c guix.Canvas) {
	r := l.outer.Size().Rect()
	l.PaintBackground(c, r)
	l.Container.Paint(c)
	l.PaintBorder(c, r)
}
