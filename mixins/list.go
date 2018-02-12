// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"fmt"

	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/base"
	"github.com/vcaesar/guix/mixins/parts"
)

type ListOuter interface {
	base.ContainerOuter
	ContainsItem(guix.AdapterItem) bool
	PaintBackground(c guix.Canvas, r math.Rect)
	PaintMouseOverBackground(c guix.Canvas, r math.Rect)
	PaintSelection(c guix.Canvas, r math.Rect)
	PaintBorder(c guix.Canvas, r math.Rect)
}

type itemDetails struct {
	child               *guix.Child
	index               int
	mark                int
	onClickSubscription guix.EventSubscription
}

type List struct {
	base.Container
	parts.BackgroundBorderPainter
	parts.Focusable

	outer ListOuter

	theme                    guix.Theme
	adapter                  guix.ListAdapter
	scrollBar                guix.ScrollBar
	scrollBarChild           *guix.Child
	scrollBarEnabled         bool
	selectedItem             guix.AdapterItem
	onSelectionChanged       guix.Event
	details                  map[guix.AdapterItem]itemDetails
	orientation              guix.Orientation
	scrollOffset             int
	itemSize                 math.Size
	itemCount                int // Count number of items in the adapter
	layoutMark               int
	mousePosition            math.Point
	itemMouseOver            *guix.Child
	onItemClicked            guix.Event
	dataChangedSubscription  guix.EventSubscription
	dataReplacedSubscription guix.EventSubscription
}

func (l *List) Init(outer ListOuter, theme guix.Theme) {
	l.outer = outer
	l.Container.Init(outer, theme)
	l.BackgroundBorderPainter.Init(outer)
	l.Focusable.Init(outer)

	l.theme = theme
	l.scrollBar = theme.CreateScrollBar()
	l.scrollBarChild = l.AddChild(l.scrollBar)
	l.scrollBarEnabled = true
	l.scrollBar.OnScroll(func(from, to int) { l.SetScrollOffset(from) })

	l.SetOrientation(guix.Vertical)
	l.SetBackgroundBrush(guix.TransparentBrush)
	l.SetMouseEventTarget(true)

	l.details = make(map[guix.AdapterItem]itemDetails)

	// Interface compliance test
	_ = guix.List(l)
}

func (l *List) UpdateItemMouseOver() {
	if !l.IsMouseOver() {
		if l.itemMouseOver != nil {
			l.itemMouseOver = nil
			l.Redraw()
		}
		return
	}
	for _, details := range l.details {
		if details.child.Bounds().Contains(l.mousePosition) {
			if l.itemMouseOver != details.child {
				l.itemMouseOver = details.child
				l.Redraw()
				return
			}
		}
	}
}

func (l *List) LayoutChildren() {
	if l.adapter == nil {
		l.outer.RemoveAll()
		return
	}

	if !l.RelayoutSuspended() {
		// Disable relayout on AddChild / RemoveChild as we're performing layout here.
		l.SetRelayoutSuspended(true)
		defer l.SetRelayoutSuspended(false)
	}

	s := l.outer.Size().Contract(l.Padding())
	o := l.Padding().LT()

	var itemSize math.Size
	if l.orientation.Horizontal() {
		itemSize = math.Size{W: l.itemSize.W, H: s.H}
	} else {
		itemSize = math.Size{W: s.W, H: l.itemSize.H}
	}

	startIndex, endIndex := l.VisibleItemRange(true)
	majorAxisItemSize := l.MajorAxisItemSize()

	d := startIndex*majorAxisItemSize - l.scrollOffset

	mark := l.layoutMark
	l.layoutMark++

	for idx := startIndex; idx < endIndex; idx++ {
		item := l.adapter.ItemAt(idx)

		details, found := l.details[item]
		if found {
			if details.mark == mark {
				panic(fmt.Errorf("Adapter for control '%s' returned duplicate item (%v) for indices %v and %v",
					guix.Path(l.outer), item, details.index, idx))
			}
		} else {
			control := l.adapter.Create(l.theme, idx)
			details.onClickSubscription = control.OnClick(func(ev guix.MouseEvent) {
				l.ItemClicked(ev, item)
			})
			details.child = l.AddChildAt(0, control)
		}
		details.mark = mark
		details.index = idx
		l.details[item] = details

		c := details.child
		cm := c.Control.Margin()
		cs := itemSize.Contract(cm).Max(math.ZeroSize)
		if l.orientation.Horizontal() {
			c.Layout(math.CreateRect(d, cm.T, d+cs.W, cm.T+cs.H).Offset(o))
		} else {
			c.Layout(math.CreateRect(cm.L, d, cm.L+cs.W, d+cs.H).Offset(o))
		}
		d += majorAxisItemSize
	}

	// Reap unused items
	for item, details := range l.details {
		if details.mark != mark {
			details.onClickSubscription.Unlisten()
			l.RemoveChild(details.child.Control)
			delete(l.details, item)
		}
	}

	if l.scrollBarEnabled {
		ss := l.scrollBar.DesiredSize(math.ZeroSize, s)
		if l.Orientation().Horizontal() {
			l.scrollBarChild.Layout(math.CreateRect(0, s.H-ss.H, s.W, s.H).Canon().Offset(o))
		} else {
			l.scrollBarChild.Layout(math.CreateRect(s.W-ss.W, 0, s.W, s.H).Canon().Offset(o))
		}

		// Only show the scroll bar if needed
		entireContentVisible := startIndex == 0 && endIndex == l.itemCount
		l.scrollBar.SetVisible(!entireContentVisible)
	}

	l.UpdateItemMouseOver()
}

func (l *List) SetSize(size math.Size) {
	l.Layoutable.SetSize(size)
	// Ensure scroll offset is still valid
	l.SetScrollOffset(l.scrollOffset)
}

func (l *List) DesiredSize(min, max math.Size) math.Size {
	if l.adapter == nil {
		return min
	}
	count := math.Max(l.itemCount, 1)
	var s math.Size
	if l.orientation.Horizontal() {
		s = math.Size{W: l.itemSize.W * count, H: l.itemSize.H}
	} else {
		s = math.Size{W: l.itemSize.W, H: l.itemSize.H * count}
	}
	if l.scrollBarEnabled {
		if l.orientation.Horizontal() {
			s.H += l.scrollBar.DesiredSize(min, max).H
		} else {
			s.W += l.scrollBar.DesiredSize(min, max).W
		}
	}
	return s.Expand(l.outer.Padding()).Clamp(min, max)
}

func (l *List) ScrollBarEnabled(bool) bool {
	return l.scrollBarEnabled
}

func (l *List) SetScrollBarEnabled(enabled bool) {
	if l.scrollBarEnabled != enabled {
		l.scrollBarEnabled = enabled
		l.Relayout()
	}
}

func (l *List) SetScrollOffset(scrollOffset int) {
	if l.adapter == nil {
		return
	}
	s := l.outer.Size().Contract(l.outer.Padding())
	if l.orientation.Horizontal() {
		maxScroll := math.Max(l.itemSize.W*l.itemCount-s.W, 0)
		scrollOffset = math.Clamp(scrollOffset, 0, maxScroll)
		l.scrollBar.SetScrollPosition(scrollOffset, scrollOffset+s.W)
	} else {
		maxScroll := math.Max(l.itemSize.H*l.itemCount-s.H, 0)
		scrollOffset = math.Clamp(scrollOffset, 0, maxScroll)
		l.scrollBar.SetScrollPosition(scrollOffset, scrollOffset+s.H)
	}
	if l.scrollOffset != scrollOffset {
		l.scrollOffset = scrollOffset
		l.LayoutChildren()
	}
}

func (l *List) MajorAxisItemSize() int {
	return l.orientation.Major(l.itemSize.WH())
}

func (l *List) VisibleItemRange(includePartiallyVisible bool) (startIndex, endIndex int) {
	if l.itemCount == 0 {
		return 0, 0
	}
	s := l.outer.Size()
	p := l.outer.Padding()
	majorAxisItemSize := l.MajorAxisItemSize()
	if majorAxisItemSize == 0 {
		return 0, 0
	}
	startIndex = l.scrollOffset
	if !includePartiallyVisible {
		startIndex += majorAxisItemSize - 1
	}
	if l.orientation.Horizontal() {
		endIndex = l.scrollOffset + s.W - p.W()
	} else {
		endIndex = l.scrollOffset + s.H - p.H()
	}
	if includePartiallyVisible {
		endIndex += majorAxisItemSize - 1
	}
	startIndex = math.Max(startIndex/majorAxisItemSize, 0)
	endIndex = math.Min(endIndex/majorAxisItemSize, l.itemCount)

	return startIndex, endIndex
}

func (l *List) SizeChanged() {
	l.itemSize = l.adapter.Size(l.theme)
	l.scrollBar.SetScrollLimit(l.itemCount * l.MajorAxisItemSize())
	l.SetScrollOffset(l.scrollOffset)
	l.outer.Relayout()
}

func (l *List) DataChanged(recreateControls bool) {
	if recreateControls {
		for item, details := range l.details {
			details.onClickSubscription.Unlisten()
			l.RemoveChild(details.child.Control)
			delete(l.details, item)
		}
	}
	l.itemCount = l.adapter.Count()
	l.SizeChanged()
}

func (l *List) DataReplaced() {
	l.selectedItem = nil
	l.DataChanged(true)
}

func (l *List) Paint(c guix.Canvas) {
	r := l.outer.Size().Rect()
	l.outer.PaintBackground(c, r)
	l.Container.Paint(c)
	l.outer.PaintBorder(c, r)
}

func (l *List) PaintSelection(c guix.Canvas, r math.Rect) {
	c.DrawRoundedRect(r, 2.0, 2.0, 2.0, 2.0, guix.WhitePen, guix.TransparentBrush)
}

func (l *List) PaintMouseOverBackground(c guix.Canvas, r math.Rect) {
	c.DrawRoundedRect(r, 2.0, 2.0, 2.0, 2.0, guix.TransparentPen, guix.CreateBrush(guix.Gray90))
}

func (l *List) SelectPrevious() {
	if l.selectedItem != nil {
		selectedIndex := l.adapter.ItemIndex(l.selectedItem)
		l.Select(l.adapter.ItemAt(math.Mod(selectedIndex-1, l.itemCount)))
	} else {
		l.Select(l.adapter.ItemAt(0))
	}
}

func (l *List) SelectNext() {
	if l.selectedItem != nil {
		selectedIndex := l.adapter.ItemIndex(l.selectedItem)
		l.Select(l.adapter.ItemAt(math.Mod(selectedIndex+1, l.itemCount)))
	} else {
		l.Select(l.adapter.ItemAt(0))
	}
}

func (l *List) ContainsItem(item guix.AdapterItem) bool {
	return l.adapter != nil && l.adapter.ItemIndex(item) >= 0
}

func (l *List) RemoveAll() {
	for _, details := range l.details {
		details.onClickSubscription.Unlisten()
		l.outer.RemoveChild(details.child.Control)
	}
	l.details = make(map[guix.AdapterItem]itemDetails)
}

// PaintChildren overrides
func (l *List) PaintChild(c guix.Canvas, child *guix.Child, idx int) {
	if child == l.itemMouseOver {
		b := child.Bounds().Expand(child.Control.Margin())
		l.outer.PaintMouseOverBackground(c, b)
	}
	l.PaintChildren.PaintChild(c, child, idx)
	if selected, found := l.details[l.selectedItem]; found {
		if child == selected.child {
			b := child.Bounds().Expand(child.Control.Margin())
			l.outer.PaintSelection(c, b)
		}
	}
}

// InputEventHandler override
func (l *List) MouseMove(ev guix.MouseEvent) {
	l.InputEventHandler.MouseMove(ev)
	l.mousePosition = ev.Point
	l.UpdateItemMouseOver()
}

func (l *List) MouseExit(ev guix.MouseEvent) {
	l.InputEventHandler.MouseExit(ev)
	l.itemMouseOver = nil
}

func (l *List) MouseScroll(ev guix.MouseEvent) (consume bool) {
	if ev.ScrollY == 0 {
		return l.InputEventHandler.MouseScroll(ev)
	}
	prevOffset := l.scrollOffset
	if l.orientation.Horizontal() {
		delta := ev.ScrollY * l.itemSize.W / 8
		l.SetScrollOffset(l.scrollOffset - delta)
	} else {
		delta := ev.ScrollY * l.itemSize.H / 8
		l.SetScrollOffset(l.scrollOffset - delta)
	}
	return prevOffset != l.scrollOffset
}

func (l *List) KeyPress(ev guix.KeyboardEvent) (consume bool) {
	if l.itemCount > 0 {
		if l.orientation.Horizontal() {
			switch ev.Key {
			case guix.KeyLeft:
				l.SelectPrevious()
				return true
			case guix.KeyRight:
				l.SelectNext()
				return true
			case guix.KeyPageUp:
				l.SetScrollOffset(l.scrollOffset - l.Size().W)
				return true
			case guix.KeyPageDown:
				l.SetScrollOffset(l.scrollOffset + l.Size().W)
				return true
			}
		} else {
			switch ev.Key {
			case guix.KeyUp:
				l.SelectPrevious()
				return true
			case guix.KeyDown:
				l.SelectNext()
				return true
			case guix.KeyPageUp:
				l.SetScrollOffset(l.scrollOffset - l.Size().H)
				return true
			case guix.KeyPageDown:
				l.SetScrollOffset(l.scrollOffset + l.Size().H)
				return true
			}
		}
	}
	return l.Container.KeyPress(ev)
}

// guix.List compliance
func (l *List) Adapter() guix.ListAdapter {
	return l.adapter
}

func (l *List) SetAdapter(adapter guix.ListAdapter) {
	if l.adapter != adapter {
		if l.adapter != nil {
			l.dataChangedSubscription.Unlisten()
			l.dataReplacedSubscription.Unlisten()
		}
		l.adapter = adapter
		if l.adapter != nil {
			l.dataChangedSubscription = l.adapter.OnDataChanged(l.DataChanged)
			l.dataReplacedSubscription = l.adapter.OnDataReplaced(l.DataReplaced)
		}
		l.DataReplaced()
	}
}

func (l *List) Orientation() guix.Orientation {
	return l.orientation
}

func (l *List) SetOrientation(o guix.Orientation) {
	l.scrollBar.SetOrientation(o)
	if l.orientation != o {
		l.orientation = o
		l.Relayout()
	}
}

func (l *List) ScrollTo(item guix.AdapterItem) {
	idx := l.adapter.ItemIndex(item)
	startIndex, endIndex := l.VisibleItemRange(false)
	if idx < startIndex {
		if l.Orientation().Horizontal() {
			l.SetScrollOffset(l.itemSize.W * idx)
		} else {
			l.SetScrollOffset(l.itemSize.H * idx)
		}
	} else if idx >= endIndex {
		count := endIndex - startIndex
		if l.Orientation().Horizontal() {
			l.SetScrollOffset(l.itemSize.W * (idx - count + 1))
		} else {
			l.SetScrollOffset(l.itemSize.H * (idx - count + 1))
		}
	}
}

func (l *List) IsItemVisible(item guix.AdapterItem) bool {
	_, found := l.details[item]
	return found
}

func (l *List) ItemControl(item guix.AdapterItem) guix.Control {
	if item, found := l.details[item]; found {
		return item.child.Control
	}
	return nil
}

func (l *List) ItemClicked(ev guix.MouseEvent, item guix.AdapterItem) {
	if l.onItemClicked != nil {
		l.onItemClicked.Fire(ev, item)
	}
	l.Select(item)
}

func (l *List) OnItemClicked(f func(guix.MouseEvent, guix.AdapterItem)) guix.EventSubscription {
	if l.onItemClicked == nil {
		l.onItemClicked = guix.CreateEvent(f)
	}
	return l.onItemClicked.Listen(f)
}

func (l *List) Selected() guix.AdapterItem {
	return l.selectedItem
}

func (l *List) Select(item guix.AdapterItem) bool {
	if l.selectedItem != item {
		if !l.outer.ContainsItem(item) {
			return false
		}
		l.selectedItem = item
		if l.onSelectionChanged != nil {
			l.onSelectionChanged.Fire(item)
		}
		l.Redraw()
	}
	l.ScrollTo(item)
	return true
}

func (l *List) OnSelectionChanged(f func(guix.AdapterItem)) guix.EventSubscription {
	if l.onItemClicked == nil {
		l.onSelectionChanged = guix.CreateEvent(f)
	}
	return l.onSelectionChanged.Listen(f)
}
