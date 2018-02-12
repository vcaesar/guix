// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins/parts"
)

type TreeOuter interface {
	ListOuter
	PaintUnexpandedSelection(c guix.Canvas, r math.Rect)
}

type Tree struct {
	List
	parts.Focusable
	outer       TreeOuter
	treeAdapter guix.TreeAdapter
	listAdapter *TreeToListAdapter
	creator     TreeControlCreator
}

func (t *Tree) Init(outer TreeOuter, theme guix.Theme) {
	t.List.Init(outer, theme)
	t.Focusable.Init(outer)
	t.outer = outer
	t.creator = defaultTreeControlCreator{}

	// Interface compliance test
	_ = guix.Tree(t)
}

func (t *Tree) SetControlCreator(c TreeControlCreator) {
	t.creator = c
	if t.treeAdapter != nil {
		t.listAdapter = CreateTreeToListAdapter(t.treeAdapter, t.creator)
		t.DataReplaced()
	}
}

// guix.Tree complaince
func (t *Tree) SetAdapter(adapter guix.TreeAdapter) {
	if t.treeAdapter == adapter {
		return
	}
	if adapter != nil {
		t.treeAdapter = adapter
		t.listAdapter = CreateTreeToListAdapter(adapter, t.creator)
		t.List.SetAdapter(t.listAdapter)
	} else {
		t.listAdapter = nil
		t.treeAdapter = nil
		t.List.SetAdapter(nil)
	}
}

func (t *Tree) Adapter() guix.TreeAdapter {
	return t.treeAdapter
}

func (t *Tree) Show(item guix.AdapterItem) {
	t.listAdapter.ExpandItem(item)
	t.List.ScrollTo(item)
}

func (t *Tree) ContainsItem(item guix.AdapterItem) bool {
	return t.listAdapter != nil && t.listAdapter.Contains(item)
}

func (t *Tree) ExpandAll() {
	t.listAdapter.ExpandAll()
}

func (t *Tree) CollapseAll() {
	t.listAdapter.CollapseAll()
}

func (t *Tree) PaintUnexpandedSelection(c guix.Canvas, r math.Rect) {
	c.DrawRoundedRect(r, 2.0, 2.0, 2.0, 2.0, guix.CreatePen(1, guix.Gray50), guix.TransparentBrush)
}

// List override
func (t *Tree) PaintChild(c guix.Canvas, child *guix.Child, idx int) {
	t.List.PaintChild(c, child, idx)
	if t.selectedItem != nil {
		if deepest := t.listAdapter.DeepestNode(t.selectedItem); deepest != nil {
			if item := deepest.Item(); item != t.selectedItem {
				// The selected item is hidden by an unexpanded node.
				// Highlight the deepest visible node instead.
				if details, found := t.details[item]; found {
					if child == details.child {
						b := child.Bounds().Expand(child.Control.Margin())
						t.outer.PaintUnexpandedSelection(c, b)
					}
				}
			}
		}
	}
}

// InputEventHandler override
func (t *Tree) KeyPress(ev guix.KeyboardEvent) (consume bool) {
	switch ev.Key {
	case guix.KeyLeft:
		if item := t.Selected(); item != nil {
			node := t.listAdapter.DeepestNode(item)
			if node.Collapse() {
				return true
			}
			if p := node.Parent(); p != nil {
				return t.Select(p.Item())
			}
		}
	case guix.KeyRight:
		if item := t.Selected(); item != nil {
			node := t.listAdapter.DeepestNode(item)
			if node.Expand() {
				return true
			}
		}
	}
	return t.List.KeyPress(ev)
}

type defaultTreeControlCreator struct{}

func (defaultTreeControlCreator) Create(theme guix.Theme, control guix.Control, node *TreeToListNode) guix.Control {
	ll := theme.CreateLinearLayout()
	ll.SetDirection(guix.LeftToRight)

	btn := theme.CreateButton()
	btn.SetBackgroundBrush(guix.TransparentBrush)
	btn.SetBorderPen(guix.CreatePen(1, guix.Gray30))
	btn.SetMargin(math.Spacing{L: 2, R: 2, T: 1, B: 1})
	btn.OnClick(func(ev guix.MouseEvent) {
		if ev.Button == guix.MouseButtonLeft {
			node.ToggleExpanded()
		}
	})

	update := func() {
		btn.SetVisible(!node.IsLeaf())
		if node.IsExpanded() {
			btn.SetText("-")
		} else {
			btn.SetText("+")
		}
	}
	update()

	guix.WhileAttached(btn, node.OnChange, update)

	ll.AddChild(btn)
	ll.AddChild(control)
	ll.SetPadding(math.Spacing{L: 16 * node.Depth()})
	return ll
}

func (defaultTreeControlCreator) Size(theme guix.Theme, treeControlSize math.Size) math.Size {
	return treeControlSize
}
