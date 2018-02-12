// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins"
)

type Tree struct {
	mixins.Tree
	theme *Theme
}

var expandedPoly = guix.Polygon{
	guix.PolygonVertex{Position: math.Point{X: 2, Y: 3}},
	guix.PolygonVertex{Position: math.Point{X: 8, Y: 3}},
	guix.PolygonVertex{Position: math.Point{X: 5, Y: 8}},
}

var collapsedPoly = guix.Polygon{
	guix.PolygonVertex{Position: math.Point{X: 3, Y: 2}},
	guix.PolygonVertex{Position: math.Point{X: 8, Y: 5}},
	guix.PolygonVertex{Position: math.Point{X: 3, Y: 8}},
}

func CreateTree(theme *Theme) guix.Tree {
	t := &Tree{}
	t.Init(t, theme)
	t.SetPadding(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	t.SetBorderPen(guix.TransparentPen)
	t.theme = theme
	t.SetControlCreator(treeControlCreator{})

	return t
}

// mixins.Tree overrides
func (t *Tree) Paint(c guix.Canvas) {
	r := t.Size().Rect()

	t.Tree.Paint(c)

	if t.HasFocus() {
		s := t.theme.FocusedStyle
		c.DrawRoundedRect(r, 3, 3, 3, 3, s.Pen, s.Brush)
	}
}

func (t *Tree) PaintMouseOverBackground(c guix.Canvas, r math.Rect) {
	c.DrawRoundedRect(r, 2.0, 2.0, 2.0, 2.0, guix.TransparentPen, guix.CreateBrush(guix.Gray15))
}

// mixins.List overrides
func (l *Tree) PaintSelection(c guix.Canvas, r math.Rect) {
	s := l.theme.HighlightStyle
	c.DrawRoundedRect(r, 2.0, 2.0, 2.0, 2.0, s.Pen, s.Brush)
}

type treeControlCreator struct{}

func (treeControlCreator) Create(theme guix.Theme, control guix.Control, node *mixins.TreeToListNode) guix.Control {
	img := theme.CreateImage()
	imgSize := math.Size{W: 10, H: 10}

	ll := theme.CreateLinearLayout()
	ll.SetDirection(guix.LeftToRight)

	btn := theme.CreateButton()
	btn.SetBackgroundBrush(guix.TransparentBrush)
	btn.SetBorderPen(guix.CreatePen(1, guix.Gray30))
	btn.SetMargin(math.Spacing{L: 1, R: 1, T: 1, B: 1})
	btn.OnClick(func(ev guix.MouseEvent) {
		if ev.Button == guix.MouseButtonLeft {
			node.ToggleExpanded()
		}
	})
	btn.AddChild(img)

	update := func() {
		expanded := node.IsExpanded()
		canvas := theme.Driver().CreateCanvas(imgSize)
		btn.SetVisible(!node.IsLeaf())
		switch {
		case !btn.IsMouseDown(guix.MouseButtonLeft) && expanded:
			canvas.DrawPolygon(expandedPoly, guix.TransparentPen, guix.CreateBrush(guix.Gray70))
		case !btn.IsMouseDown(guix.MouseButtonLeft) && !expanded:
			canvas.DrawPolygon(collapsedPoly, guix.TransparentPen, guix.CreateBrush(guix.Gray70))
		case expanded:
			canvas.DrawPolygon(expandedPoly, guix.TransparentPen, guix.CreateBrush(guix.Gray30))
		case !expanded:
			canvas.DrawPolygon(collapsedPoly, guix.TransparentPen, guix.CreateBrush(guix.Gray30))
		}
		canvas.Complete()
		img.SetCanvas(canvas)
	}
	btn.OnMouseDown(func(guix.MouseEvent) { update() })
	btn.OnMouseUp(func(guix.MouseEvent) { update() })
	update()

	guix.WhileAttached(btn, node.OnChange, update)

	ll.AddChild(btn)
	ll.AddChild(control)
	ll.SetPadding(math.Spacing{L: 16 * node.Depth()})
	return ll
}

func (treeControlCreator) Size(theme guix.Theme, treeControlSize math.Size) math.Size {
	return treeControlSize
}
