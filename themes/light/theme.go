// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package light

import (
	"fmt"

	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/gxfont"
	"github.com/vcaesar/guix/themes/basic"
)

func CreateTheme(driver guix.Driver) guix.Theme {
	defaultFont, err := driver.CreateFont(gxfont.Default, 12)
	if err == nil {
		defaultFont.LoadGlyphs(32, 126)
	} else {
		fmt.Printf("Warning: Failed to load default font - %v\n", err)
	}

	defaultMonospaceFont, err := driver.CreateFont(gxfont.Monospace, 12)
	if err == nil {
		defaultFont.LoadGlyphs(32, 126)
	} else {
		fmt.Printf("Warning: Failed to load default monospace font - %v\n", err)
	}

	scrollBarRailDefaultBg := guix.Black
	scrollBarRailDefaultBg.A = 0.7

	scrollBarRailOverBg := guix.Gray20
	scrollBarRailOverBg.A = 0.7

	neonBlue := guix.ColorFromHex(0xFF5C8CFF)
	focus := guix.ColorFromHex(0xFFC4D6FF)

	return &basic.Theme{
		DriverInfo:               driver,
		DefaultFontInfo:          defaultFont,
		DefaultMonospaceFontInfo: defaultMonospaceFont,
		WindowBackground:         guix.White,

		//                                   fontColor    brushColor   penColor
		BubbleOverlayStyle:        basic.CreateStyle(guix.Gray40, guix.Gray20, guix.Gray40, 1.0),
		ButtonDefaultStyle:        basic.CreateStyle(guix.Gray40, guix.White, guix.Gray40, 1.0),
		ButtonOverStyle:           basic.CreateStyle(guix.Gray40, guix.Gray90, guix.Gray40, 1.0),
		ButtonPressedStyle:        basic.CreateStyle(guix.Gray20, guix.Gray70, guix.Gray30, 1.0),
		CodeSuggestionListStyle:   basic.CreateStyle(guix.Gray40, guix.Gray20, guix.Gray10, 1.0),
		DropDownListDefaultStyle:  basic.CreateStyle(guix.Gray40, guix.White, guix.Gray20, 1.0),
		DropDownListOverStyle:     basic.CreateStyle(guix.Gray40, guix.Gray90, guix.Gray50, 1.0),
		FocusedStyle:              basic.CreateStyle(guix.Gray20, guix.Transparent, focus, 1.0),
		HighlightStyle:            basic.CreateStyle(guix.Gray40, guix.Transparent, neonBlue, 2.0),
		LabelStyle:                basic.CreateStyle(guix.Gray40, guix.Transparent, guix.Transparent, 0.0),
		PanelBackgroundStyle:      basic.CreateStyle(guix.Gray40, guix.White, guix.Gray15, 1.0),
		ScrollBarBarDefaultStyle:  basic.CreateStyle(guix.Gray40, guix.Gray30, guix.Gray40, 1.0),
		ScrollBarBarOverStyle:     basic.CreateStyle(guix.Gray40, guix.Gray50, guix.Gray60, 1.0),
		ScrollBarRailDefaultStyle: basic.CreateStyle(guix.Gray40, scrollBarRailDefaultBg, guix.Transparent, 1.0),
		ScrollBarRailOverStyle:    basic.CreateStyle(guix.Gray40, scrollBarRailOverBg, guix.Gray20, 1.0),
		SplitterBarDefaultStyle:   basic.CreateStyle(guix.Gray40, guix.Gray80, guix.Gray40, 1.0),
		SplitterBarOverStyle:      basic.CreateStyle(guix.Gray40, guix.Gray80, guix.Gray50, 1.0),
		TabActiveHighlightStyle:   basic.CreateStyle(guix.Gray30, neonBlue, neonBlue, 0.0),
		TabDefaultStyle:           basic.CreateStyle(guix.Gray40, guix.White, guix.Gray40, 1.0),
		TabOverStyle:              basic.CreateStyle(guix.Gray30, guix.Gray90, guix.Gray50, 1.0),
		TabPressedStyle:           basic.CreateStyle(guix.Gray20, guix.Gray70, guix.Gray30, 1.0),
		TextBoxDefaultStyle:       basic.CreateStyle(guix.Gray40, guix.White, guix.Gray20, 1.0),
		TextBoxOverStyle:          basic.CreateStyle(guix.Gray40, guix.White, guix.Gray50, 1.0),
	}
}
