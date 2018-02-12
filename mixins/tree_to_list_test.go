// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"testing"

	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
)

type testTreeNode struct {
	item     guix.AdapterItem
	children []*testTreeNode
}

func (n *testTreeNode) Count() int                           { return len(n.children) }
func (n *testTreeNode) NodeAt(index int) guix.TreeNode       { return n.children[index] }
func (n *testTreeNode) Item() guix.AdapterItem               { return n.item }
func (n *testTreeNode) Create(theme guix.Theme) guix.Control { return nil }

func (n *testTreeNode) ItemIndex(item guix.AdapterItem) int {
	for i, c := range n.children {
		if item == c.item {
			return i
		}
		if idx := c.ItemIndex(item); idx >= 0 {
			return i
		}
	}
	return -1
}

type testTreeAdapter struct {
	guix.AdapterBase
	testTreeNode
}

func (n *testTreeNode) Size(theme guix.Theme) math.Size { return math.ZeroSize }

// n creates and returns a testTreeNode with the item i and children c.
func n(i guix.AdapterItem, c ...*testTreeNode) *testTreeNode {
	return &testTreeNode{item: i, children: c}
}

// a creates and returns a list and tree adapters for the children c.
func a(c ...*testTreeNode) (list_adapter *TreeToListAdapter, tree_adapter *testTreeAdapter) {
	adapter := &testTreeAdapter{}
	adapter.children = c
	return CreateTreeToListAdapter(adapter, nil), adapter
}

func test(t *testing.T, name string, adapter *TreeToListAdapter, expected ...guix.AdapterItem) {
	if len(expected) != adapter.Count() {
		t.Errorf("%s: Count was not as expected.\nExpected: %v\nGot:      %v",
			name, len(expected), adapter.Count())
	}
	for expected_index, expected_item := range expected {
		got_item := adapter.ItemAt(expected_index)
		got_index := adapter.ItemIndex(expected_item)
		if expected_item != got_item {
			t.Errorf("%s: Item at index %v was not as expected.\nExpected: %v\nGot:      %v",
				name, expected_index, expected_item, got_item)
		}
		if expected_index != got_index {
			t.Errorf("%s: Index of item %v was not as expected.\nExpected: %v\nGot:      %v",
				name, expected_item, expected_item, got_item)
		}
	}
}

func TestTreeToListNodeFlat(t *testing.T) {
	list_adapter, _ := a(n(10), n(20), n(30))
	test(t, "flat", list_adapter,
		guix.AdapterItem(10),
		guix.AdapterItem(20),
		guix.AdapterItem(30),
	)
}

func TestTreeToListNodeDeep(t *testing.T) {

	list_adapter, tree_adapter := a(
		n(100,
			n(110),
			n(120,
				n(121),
				n(122),
				n(123)),
			n(130),
			n(140,
				n(141),
				n(142))))

	test(t, "unexpanded", list_adapter,
		guix.AdapterItem(100),
	)

	list_adapter.node.children[0].Expand()
	test(t, "single expanded", list_adapter,
		guix.AdapterItem(100), // (0) 100
		guix.AdapterItem(110), // (1)  ╠══ 110
		guix.AdapterItem(120), // (2)  ╠══ 120
		guix.AdapterItem(130), // (3)  ╠══ 130
		guix.AdapterItem(140), // (4)  ╚══ 140
	)

	list_adapter.ExpandAll()
	test(t, "fully expanded", list_adapter,
		guix.AdapterItem(100), // (0) 100
		guix.AdapterItem(110), // (1)  ╠══ 110
		guix.AdapterItem(120), // (2)  ╠══ 120
		guix.AdapterItem(121), // (3)  ║    ╠══ 121
		guix.AdapterItem(122), // (4)  ║    ╠══ 122
		guix.AdapterItem(123), // (5)  ║    ╚══ 123
		guix.AdapterItem(130), // (6)  ╠══ 130
		guix.AdapterItem(140), // (7)  ╚══ 140
		guix.AdapterItem(141), // (8)       ╠══ 141
		guix.AdapterItem(142), // (9)       ╚══ 142
	)

	list_adapter.node.NodeAt(2).Collapse()
	test(t, "one collapsed", list_adapter,
		guix.AdapterItem(100), // (0) 100
		guix.AdapterItem(110), // (1)  ╠══ 110
		guix.AdapterItem(120), // (2)  ╠══ 120
		guix.AdapterItem(130), // (3)  ╠══ 130
		guix.AdapterItem(140), // (4)  ╚══ 140
		guix.AdapterItem(141), // (5)       ╠══ 141
		guix.AdapterItem(142), // (6)       ╚══ 142
	)

	tree_adapter.children[0].children = append(tree_adapter.children[0].children, n(150))
	test(t, "mutate, no data-changed", list_adapter,
		guix.AdapterItem(100), // (0) 100
		guix.AdapterItem(110), // (1)  ╠══ 110
		guix.AdapterItem(120), // (2)  ╠══ 120
		guix.AdapterItem(130), // (3)  ╠══ 130
		guix.AdapterItem(140), // (4)  ╚══ 140
		guix.AdapterItem(141), // (5)       ╠══ 141
		guix.AdapterItem(142), // (6)       ╚══ 142
	)

	tree_adapter.DataChanged(false)
	test(t, "data-changed", list_adapter,
		guix.AdapterItem(100), // (0) 100
		guix.AdapterItem(110), // (1)  ╠══ 110
		guix.AdapterItem(120), // (2)  ╠══ 120
		guix.AdapterItem(130), // (3)  ╠══ 130
		guix.AdapterItem(140), // (4)  ╠══ 140
		guix.AdapterItem(141), // (5)  ║    ╠══ 141
		guix.AdapterItem(142), // (6)  ║    ╚══ 142
		guix.AdapterItem(150), // (7)  ╚══ 150
	)
}
