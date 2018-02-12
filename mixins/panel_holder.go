// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"fmt"

	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/base"
)

type PanelTab interface {
	guix.Control
	SetText(string)
	SetActive(bool)
}

type PanelTabCreater interface {
	CreatePanelTab() PanelTab
}

type PanelHolderOuter interface {
	base.ContainerNoControlOuter
	guix.PanelHolder
	PanelTabCreater
}

type PanelEntry struct {
	Tab                   PanelTab
	Panel                 guix.Control
	MouseDownSubscription guix.EventSubscription
}

type PanelHolder struct {
	base.Container

	outer PanelHolderOuter

	theme     guix.Theme
	tabLayout guix.LinearLayout
	entries   []PanelEntry
	selected  PanelEntry
}

func insertIndex(holder guix.PanelHolder, at math.Point) int {
	count := holder.PanelCount()
	bestIndex := count
	bestScore := float32(1e20)
	score := func(point math.Point, index int) {
		score := point.Sub(at).Len()
		if score < bestScore {
			bestIndex = index
			bestScore = score
		}
	}
	for i := 0; i < holder.PanelCount(); i++ {
		tab := holder.Tab(i)
		size := tab.Size()
		ml := math.Point{Y: size.H / 2}
		mr := math.Point{Y: size.H / 2, X: size.W}
		score(guix.TransformCoordinate(ml, tab, holder), i)
		score(guix.TransformCoordinate(mr, tab, holder), i+1)
	}
	return bestIndex
}

func beginTabDragging(holder guix.PanelHolder, panel guix.Control, name string, window guix.Window) {
	var mms, mos guix.EventSubscription
	mms = window.OnMouseMove(func(ev guix.MouseEvent) {
		for _, c := range guix.TopControlsUnder(ev.WindowPoint, ev.Window) {
			if over, ok := c.C.(guix.PanelHolder); ok {
				insertAt := insertIndex(over, c.P)
				if over == holder {
					if insertAt > over.PanelIndex(panel) {
						insertAt--
					}
				}
				holder.RemovePanel(panel)
				holder = over
				holder.AddPanelAt(panel, name, insertAt)
				holder.Select(insertAt)
			}
		}
	})
	mos = window.OnMouseUp(func(guix.MouseEvent) {
		mms.Unlisten()
		mos.Unlisten()
	})
}

func (p *PanelHolder) Init(outer PanelHolderOuter, theme guix.Theme) {
	p.Container.Init(outer, theme)

	p.outer = outer
	p.theme = theme

	p.tabLayout = theme.CreateLinearLayout()
	p.tabLayout.SetDirection(guix.LeftToRight)
	p.Container.AddChild(p.tabLayout)
	p.SetMargin(math.Spacing{L: 1, T: 2, R: 1, B: 1})
	p.SetMouseEventTarget(true) // For drag-drop targets

	// Interface compliance test
	_ = guix.PanelHolder(p)
}

func (p *PanelHolder) LayoutChildren() {
	s := p.Size()

	tabHeight := p.tabLayout.DesiredSize(math.ZeroSize, s).H
	panelRect := math.CreateRect(0, tabHeight, s.W, s.H).Contract(p.Padding())

	for _, child := range p.Children() {
		if child.Control == p.tabLayout {
			child.Control.SetSize(math.Size{W: s.W, H: tabHeight})
			child.Offset = math.ZeroPoint
		} else {
			rect := panelRect.Contract(child.Control.Margin())
			child.Control.SetSize(rect.Size())
			child.Offset = rect.Min
		}
	}
}

func (p *PanelHolder) DesiredSize(min, max math.Size) math.Size {
	return max
}

func (p *PanelHolder) SelectedPanel() guix.Control {
	return p.selected.Panel
}

// guix.PanelHolder compliance
func (p *PanelHolder) AddPanel(panel guix.Control, name string) {
	p.AddPanelAt(panel, name, len(p.entries))
}

func (p *PanelHolder) AddPanelAt(panel guix.Control, name string, index int) {
	if index < 0 || index > p.PanelCount() {
		panic(fmt.Errorf("Index %d is out of bounds. Acceptable range: [%d - %d]",
			index, 0, p.PanelCount()))
	}
	tab := p.outer.CreatePanelTab()
	tab.SetText(name)
	mds := tab.OnMouseDown(func(ev guix.MouseEvent) {
		p.Select(p.PanelIndex(panel))
		beginTabDragging(p.outer, panel, name, ev.Window)
	})

	p.entries = append(p.entries, PanelEntry{})
	copy(p.entries[index+1:], p.entries[index:])
	p.entries[index] = PanelEntry{
		Panel: panel,
		Tab:   tab,
		MouseDownSubscription: mds,
	}
	p.tabLayout.AddChildAt(index, tab)

	if p.selected.Panel == nil {
		p.Select(index)
	}
}

func (p *PanelHolder) RemovePanel(panel guix.Control) {
	index := p.PanelIndex(panel)
	if index < 0 {
		panic("PanelHolder does not contain panel")
	}

	entry := p.entries[index]
	entry.MouseDownSubscription.Unlisten()
	p.entries = append(p.entries[:index], p.entries[index+1:]...)
	p.tabLayout.RemoveChildAt(index)

	if panel == p.selected.Panel {
		if p.PanelCount() > 0 {
			p.Select(math.Max(index-1, 0))
		} else {
			p.Select(-1)
		}
	}
}

func (p *PanelHolder) Select(index int) {
	if index >= p.PanelCount() {
		panic(fmt.Errorf("Index %d is out of bounds. Acceptable range: [%d - %d]",
			index, -1, p.PanelCount()-1))
	}

	if p.selected.Panel != nil {
		p.selected.Tab.SetActive(false)
		p.Container.RemoveChild(p.selected.Panel)
	}

	if index >= 0 {
		p.selected = p.entries[index]
	} else {
		p.selected = PanelEntry{}
	}

	if p.selected.Panel != nil {
		p.Container.AddChild(p.selected.Panel)
		p.selected.Tab.SetActive(true)
	}
}

func (p *PanelHolder) PanelCount() int {
	return len(p.entries)
}

func (p *PanelHolder) PanelIndex(panel guix.Control) int {
	for i, e := range p.entries {
		if e.Panel == panel {
			return i
		}
	}
	return -1
}

func (p *PanelHolder) Panel(index int) guix.Control {
	return p.entries[index].Panel
}

func (p *PanelHolder) Tab(index int) guix.Control {
	return p.entries[index].Tab
}
