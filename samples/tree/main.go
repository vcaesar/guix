// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/drivers/gl"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/samples/flags"
)

// item is used as an guix.AdapterItem to identifiy each of the nodes.
// Each node's item must be equality-unique for the entire tree.
type item int

var nextUniqueItem item // the next item to used by node.add

// node is an implementation of guix.TreeNode.
type node struct {
	name     string  // The name and item for this node.
	item     item    // The unique item for this node.
	changed  func()  // Called when a new item is added to this node.
	children []*node // The list of all child nodes.
}

// add appends a new child node to n with the specified name.
func (n *node) add(name string) *node {
	child := &node{
		name:    name,
		item:    nextUniqueItem,
		changed: n.changed,
	}
	nextUniqueItem++
	n.children = append(n.children, child)
	n.changed()
	return child
}

// Count implements guix.TreeNodeContainer.
func (n *node) Count() int {
	return len(n.children)
}

// NodeAt implements guix.TreeNodeContainer.
func (n *node) NodeAt(index int) guix.TreeNode {
	return n.children[index]
}

// ItemIndex implements guix.TreeNodeContainer.
func (n *node) ItemIndex(item guix.AdapterItem) int {
	for i, c := range n.children {
		if c.item == item || c.ItemIndex(item) >= 0 {
			return i
		}
	}
	return -1
}

// Item implements guix.TreeNode.
func (n *node) Item() guix.AdapterItem {
	return n.item
}

// Create implements guix.TreeNode.
func (n *node) Create(theme guix.Theme) guix.Control {
	layout := theme.CreateLinearLayout()
	layout.SetDirection(guix.LeftToRight)

	label := theme.CreateLabel()
	label.SetText(n.name)

	textbox := theme.CreateTextBox()
	textbox.SetText(n.name)
	textbox.SetPadding(math.ZeroSpacing)
	textbox.SetMargin(math.ZeroSpacing)

	addButton := theme.CreateButton()
	addButton.SetText("+")
	addButton.OnClick(func(guix.MouseEvent) { n.add("<new>") })

	edit := func() {
		layout.RemoveAll()
		layout.AddChild(textbox)
		layout.AddChild(addButton)
		guix.SetFocus(textbox)
	}

	commit := func() {
		n.name = textbox.Text()
		label.SetText(n.name)
		layout.RemoveAll()
		layout.AddChild(label)
		layout.AddChild(addButton)
	}

	// When the user clicks the label, replace it with an editable text-box
	label.OnClick(func(guix.MouseEvent) { edit() })

	// When the text-box loses focus, replace it with a label again.
	textbox.OnLostFocus(commit)

	layout.AddChild(label)
	layout.AddChild(addButton)
	return layout
}

// adapter is an implementation of guix.TreeAdapter.
type adapter struct {
	guix.AdapterBase
	node
}

// Size implements guix.TreeAdapter.
func (a *adapter) Size(t guix.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: 18}
}

// addSpecies adds the list of species to animals.
// A map of name to item is returned.
func addSpecies(animals *node) map[string]item {
	items := make(map[string]item)

	add := func(to *node, name string) *node {
		n := to.add(name)
		items[name] = n.item
		return n
	}

	mammals := add(animals, "Mammals")
	add(mammals, "Cats")
	add(mammals, "Dogs")
	add(mammals, "Horses")
	add(mammals, "Duck-billed platypuses")

	birds := add(animals, "Birds")
	add(birds, "Peacocks")
	add(birds, "Doves")

	reptiles := add(animals, "Reptiles")
	add(reptiles, "Lizards")
	add(reptiles, "Turtles")
	add(reptiles, "Crocodiles")
	add(reptiles, "Snakes")

	arthropods := add(animals, "Arthropods")

	crustaceans := add(arthropods, "Crustaceans")
	add(crustaceans, "Crabs")
	add(crustaceans, "Lobsters")

	insects := add(arthropods, "Insects")
	add(insects, "Ants")
	add(insects, "Bees")

	arachnids := add(arthropods, "Arachnids")
	add(arachnids, "Spiders")
	add(arachnids, "Scorpions")

	return items
}

func appMain(driver guix.Driver) {
	theme := flags.CreateTheme(driver)

	layout := theme.CreateLinearLayout()
	layout.SetDirection(guix.TopToBottom)

	adapter := &adapter{}

	// hook up node changed function to the adapter OnDataChanged event.
	adapter.changed = func() { adapter.DataChanged(false) }

	// add all the species to the 'Animals' root node.
	items := addSpecies(adapter.add("Animals"))

	tree := theme.CreateTree()
	tree.SetAdapter(adapter)
	tree.Select(items["Doves"])
	tree.Show(tree.Selected())

	layout.AddChild(tree)

	row := theme.CreateLinearLayout()
	row.SetDirection(guix.LeftToRight)
	layout.AddChild(row)

	expandAll := theme.CreateButton()
	expandAll.SetText("Expand All")
	expandAll.OnClick(func(guix.MouseEvent) { tree.ExpandAll() })
	row.AddChild(expandAll)

	collapseAll := theme.CreateButton()
	collapseAll.SetText("Collapse All")
	collapseAll.OnClick(func(guix.MouseEvent) { tree.CollapseAll() })
	row.AddChild(collapseAll)

	window := theme.CreateWindow(800, 600, "Tree view")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(layout)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

func main() {
	gl.StartDriver(appMain)
}
