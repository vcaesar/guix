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

func appMain(driver guix.Driver) {
	theme := flags.CreateTheme(driver)
	layout := theme.CreateLinearLayout()
	layout.SetSizeMode(guix.Fill)

	buttonState := map[guix.Button]func() bool{}
	update := func() {
		for button, f := range buttonState {
			button.SetChecked(f())
		}
	}

	button := func(name string, action func(), isSelected func() bool) guix.Button {
		b := theme.CreateButton()
		b.SetText(name)
		b.OnClick(func(guix.MouseEvent) { action(); update() })
		layout.AddChild(b)
		buttonState[b] = isSelected
		return b
	}

	button("TopToBottom",
		func() { layout.SetDirection(guix.TopToBottom) },
		func() bool { return layout.Direction().TopToBottom() },
	)
	button("LeftToRight",
		func() { layout.SetDirection(guix.LeftToRight) },
		func() bool { return layout.Direction().LeftToRight() },
	)
	button("BottomToTop",
		func() { layout.SetDirection(guix.BottomToTop) },
		func() bool { return layout.Direction().BottomToTop() },
	)
	button("RightToLeft",
		func() { layout.SetDirection(guix.RightToLeft) },
		func() bool { return layout.Direction().RightToLeft() },
	)

	button("AlignLeft",
		func() { layout.SetHorizontalAlignment(guix.AlignLeft) },
		func() bool { return layout.HorizontalAlignment().AlignLeft() },
	)
	button("AlignCenter",
		func() { layout.SetHorizontalAlignment(guix.AlignCenter) },
		func() bool { return layout.HorizontalAlignment().AlignCenter() },
	)
	button("AlignRight",
		func() { layout.SetHorizontalAlignment(guix.AlignRight) },
		func() bool { return layout.HorizontalAlignment().AlignRight() },
	)

	button("AlignTop",
		func() { layout.SetVerticalAlignment(guix.AlignTop) },
		func() bool { return layout.VerticalAlignment().AlignTop() },
	)
	button("AlignMiddle",
		func() { layout.SetVerticalAlignment(guix.AlignMiddle) },
		func() bool { return layout.VerticalAlignment().AlignMiddle() },
	)
	button("AlignBottom",
		func() { layout.SetVerticalAlignment(guix.AlignBottom) },
		func() bool { return layout.VerticalAlignment().AlignBottom() },
	)

	update()

	window := theme.CreateWindow(800, 600, "Linear layout")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(layout)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

func main() {
	gl.StartDriver(appMain)
}
