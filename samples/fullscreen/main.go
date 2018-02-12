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

	window := theme.CreateWindow(200, 150, "Window")
	window.OnClose(driver.Terminate)
	window.SetScale(flags.DefaultScaleFactor)
	window.SetPadding(math.Spacing{L: 10, R: 10, T: 10, B: 10})
	button := theme.CreateButton()
	button.SetHorizontalAlignment(guix.AlignCenter)
	button.SetSizeMode(guix.Fill)
	toggle := func() {
		fullscreen := !window.Fullscreen()
		window.SetFullscreen(fullscreen)
		if fullscreen {
			button.SetText("Make windowed")
		} else {
			button.SetText("Make fullscreen")
		}
	}
	button.SetText("Make fullscreen")
	button.OnClick(func(guix.MouseEvent) { toggle() })
	window.AddChild(button)
}

func main() {
	gl.StartDriver(appMain)
}
