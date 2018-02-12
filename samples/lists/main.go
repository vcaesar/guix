// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/drivers/gl"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/samples/flags"
)

// Number picker uses the guix.DefaultAdapter for driving a list
func numberPicker(theme guix.Theme, overlay guix.BubbleOverlay) guix.Control {
	items := []string{
		"zero", "one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen",
		"sixteen", "seventeen", "eighteen", "nineteen", "twenty",
	}

	adapter := guix.CreateDefaultAdapter()
	adapter.SetItems(items)

	layout := theme.CreateLinearLayout()
	layout.SetDirection(guix.TopToBottom)

	label0 := theme.CreateLabel()
	label0.SetText("Numbers:")
	layout.AddChild(label0)

	dropList := theme.CreateDropDownList()
	dropList.SetAdapter(adapter)
	dropList.SetBubbleOverlay(overlay)
	layout.AddChild(dropList)

	list := theme.CreateList()
	list.SetAdapter(adapter)
	list.SetOrientation(guix.Vertical)
	layout.AddChild(list)

	label1 := theme.CreateLabel()
	label1.SetMargin(math.Spacing{T: 30})
	label1.SetText("Selected number:")
	layout.AddChild(label1)

	selected := theme.CreateLabel()
	layout.AddChild(selected)

	dropList.OnSelectionChanged(func(item guix.AdapterItem) {
		if list.Selected() != item {
			list.Select(item)
		}
	})

	list.OnSelectionChanged(func(item guix.AdapterItem) {
		if dropList.Selected() != item {
			dropList.Select(item)
		}
		selected.SetText(fmt.Sprintf("%s - %d", item, adapter.ItemIndex(item)))
	})

	return layout
}

type customAdapter struct {
	guix.AdapterBase
}

func (a *customAdapter) Count() int {
	return 1000
}

func (a *customAdapter) ItemAt(index int) guix.AdapterItem {
	return index // This adapter uses integer indices as AdapterItems
}

func (a *customAdapter) ItemIndex(item guix.AdapterItem) int {
	return item.(int) // Inverse of ItemAt()
}

func (a *customAdapter) Size(theme guix.Theme) math.Size {
	return math.Size{W: 100, H: 100}
}

func (a *customAdapter) Create(theme guix.Theme, index int) guix.Control {
	phase := float32(index) / 1000
	c := guix.Color{
		R: 0.5 + 0.5*math.Sinf(math.TwoPi*(phase+0.000)),
		G: 0.5 + 0.5*math.Sinf(math.TwoPi*(phase+0.333)),
		B: 0.5 + 0.5*math.Sinf(math.TwoPi*(phase+0.666)),
		A: 1.0,
	}
	i := theme.CreateImage()
	i.SetBackgroundBrush(guix.CreateBrush(c))
	i.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	i.OnMouseEnter(func(ev guix.MouseEvent) {
		i.SetBorderPen(guix.CreatePen(2, guix.Gray80))
	})
	i.OnMouseExit(func(ev guix.MouseEvent) {
		i.SetBorderPen(guix.TransparentPen)
	})
	i.OnMouseDown(func(ev guix.MouseEvent) {
		i.SetBackgroundBrush(guix.CreateBrush(c.MulRGB(0.7)))
	})
	i.OnMouseUp(func(ev guix.MouseEvent) {
		i.SetBackgroundBrush(guix.CreateBrush(c))
	})
	return i
}

// Color picker uses the customAdapter for driving a list
func colorPicker(theme guix.Theme) guix.Control {
	layout := theme.CreateLinearLayout()
	layout.SetDirection(guix.TopToBottom)

	label0 := theme.CreateLabel()
	label0.SetText("Color palette:")
	layout.AddChild(label0)

	adapter := &customAdapter{}

	list := theme.CreateList()
	list.SetAdapter(adapter)
	list.SetOrientation(guix.Horizontal)
	layout.AddChild(list)

	label1 := theme.CreateLabel()
	label1.SetMargin(math.Spacing{T: 30})
	label1.SetText("Selected color:")
	layout.AddChild(label1)

	selected := theme.CreateImage()
	selected.SetExplicitSize(math.Size{W: 32, H: 32})
	layout.AddChild(selected)

	list.OnSelectionChanged(func(item guix.AdapterItem) {
		if item != nil {
			control := list.ItemControl(item)
			selected.SetBackgroundBrush(control.(guix.Image).BackgroundBrush())
		}
	})

	return layout
}

func appMain(driver guix.Driver) {
	theme := flags.CreateTheme(driver)

	overlay := theme.CreateBubbleOverlay()

	holder := theme.CreatePanelHolder()
	holder.AddPanel(numberPicker(theme, overlay), "Default adapter")
	holder.AddPanel(colorPicker(theme), "Custom adapter")

	window := theme.CreateWindow(800, 600, "Lists")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(holder)
	window.AddChild(overlay)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

func main() {
	gl.StartDriver(appMain)
}
