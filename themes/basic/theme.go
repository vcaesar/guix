// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import (
	"github.com/vcaesar/guix"
)

type Theme struct {
	DriverInfo               guix.Driver
	DefaultFontInfo          guix.Font
	DefaultMonospaceFontInfo guix.Font

	WindowBackground guix.Color

	BubbleOverlayStyle        Style
	ButtonDefaultStyle        Style
	ButtonOverStyle           Style
	ButtonPressedStyle        Style
	CodeSuggestionListStyle   Style
	DropDownListDefaultStyle  Style
	DropDownListOverStyle     Style
	FocusedStyle              Style
	HighlightStyle            Style
	LabelStyle                Style
	PanelBackgroundStyle      Style
	ScrollBarBarDefaultStyle  Style
	ScrollBarBarOverStyle     Style
	ScrollBarRailDefaultStyle Style
	ScrollBarRailOverStyle    Style
	SplitterBarDefaultStyle   Style
	SplitterBarOverStyle      Style
	TabActiveHighlightStyle   Style
	TabDefaultStyle           Style
	TabOverStyle              Style
	TabPressedStyle           Style
	TextBoxDefaultStyle       Style
	TextBoxOverStyle          Style
}

// guix.Theme compliance
func (t *Theme) Driver() guix.Driver {
	return t.DriverInfo
}

func (t *Theme) DefaultFont() guix.Font {
	return t.DefaultFontInfo
}

func (t *Theme) SetDefaultFont(f guix.Font) {
	t.DefaultFontInfo = f
}

func (t *Theme) DefaultMonospaceFont() guix.Font {
	return t.DefaultMonospaceFontInfo
}

func (t *Theme) SetDefaultMonospaceFont(f guix.Font) {
	t.DefaultMonospaceFontInfo = f
}

func (t *Theme) CreateBubbleOverlay() guix.BubbleOverlay {
	return CreateBubbleOverlay(t)
}

func (t *Theme) CreateButton() guix.Button {
	return CreateButton(t)
}

func (t *Theme) CreateCodeEditor() guix.CodeEditor {
	return CreateCodeEditor(t)
}

func (t *Theme) CreateDropDownList() guix.DropDownList {
	return CreateDropDownList(t)
}

func (t *Theme) CreateImage() guix.Image {
	return CreateImage(t)
}

func (t *Theme) CreateLabel() guix.Label {
	return CreateLabel(t)
}

func (t *Theme) CreateLinearLayout() guix.LinearLayout {
	return CreateLinearLayout(t)
}

func (t *Theme) CreateList() guix.List {
	return CreateList(t)
}

func (t *Theme) CreatePanelHolder() guix.PanelHolder {
	return CreatePanelHolder(t)
}

func (t *Theme) CreateProgressBar() guix.ProgressBar {
	return CreateProgressBar(t)
}

func (t *Theme) CreateScrollBar() guix.ScrollBar {
	return CreateScrollBar(t)
}

func (t *Theme) CreateScrollLayout() guix.ScrollLayout {
	return CreateScrollLayout(t)
}

func (t *Theme) CreateSplitterLayout() guix.SplitterLayout {
	return CreateSplitterLayout(t)
}

func (t *Theme) CreateTableLayout() guix.TableLayout {
	return CreateTableLayout(t)
}

func (t *Theme) CreateTextBox() guix.TextBox {
	return CreateTextBox(t)
}

func (t *Theme) CreateTree() guix.Tree {
	return CreateTree(t)
}

func (t *Theme) CreateWindow(width, height int, title string) guix.Window {
	return CreateWindow(t, width, height, title)
}
