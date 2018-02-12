// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"fmt"
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"strings"
)

type CodeEditorOuter interface {
	TextBoxOuter
	CreateSuggestionList() guix.List
}

type CodeEditor struct {
	TextBox
	outer              CodeEditorOuter
	layers             guix.CodeSyntaxLayers
	suggestionAdapter  *SuggestionAdapter
	suggestionList     guix.List
	suggestionProvider guix.CodeSuggestionProvider
	tabWidth           int
	theme              guix.Theme
}

func (t *CodeEditor) updateSpans(edits []guix.TextBoxEdit) {
	runeCount := len(t.controller.TextRunes())
	for _, l := range t.layers {
		l.UpdateSpans(runeCount, edits)
	}
}

func (t *CodeEditor) Init(outer CodeEditorOuter, driver guix.Driver, theme guix.Theme, font guix.Font) {
	t.outer = outer
	t.tabWidth = 2
	t.theme = theme

	t.suggestionAdapter = &SuggestionAdapter{}
	t.suggestionList = t.outer.CreateSuggestionList()
	t.suggestionList.SetAdapter(t.suggestionAdapter)

	t.TextBox.Init(outer, driver, theme, font)
	t.controller.OnTextChanged(t.updateSpans)

	// Interface compliance test
	_ = guix.CodeEditor(t)
}

func (t *CodeEditor) ItemSize(theme guix.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: t.font.GlyphMaxSize().H}
}

func (t *CodeEditor) CreateSuggestionList() guix.List {
	l := t.theme.CreateList()
	l.SetBackgroundBrush(guix.DefaultBrush)
	l.SetBorderPen(guix.DefaultPen)
	return l
}

func (t *CodeEditor) SyntaxLayers() guix.CodeSyntaxLayers {
	return t.layers
}

func (t *CodeEditor) SetSyntaxLayers(layers guix.CodeSyntaxLayers) {
	t.layers = layers
	t.onRedrawLines.Fire()
}

func (t *CodeEditor) TabWidth() int {
	return t.tabWidth
}

func (t *CodeEditor) SetTabWidth(tabWidth int) {
	t.tabWidth = tabWidth
}

func (t *CodeEditor) SuggestionProvider() guix.CodeSuggestionProvider {
	return t.suggestionProvider
}

func (t *CodeEditor) SetSuggestionProvider(provider guix.CodeSuggestionProvider) {
	if t.suggestionProvider != provider {
		t.suggestionProvider = provider
		if t.IsSuggestionListShowing() {
			t.ShowSuggestionList() // Update list
		}
	}
}

func (t *CodeEditor) IsSuggestionListShowing() bool {
	return t.outer.Children().Find(t.suggestionList) != nil
}

func (t *CodeEditor) SortSuggestionList() {
	caret := t.controller.LastCaret()
	partial := t.controller.TextRange(t.controller.WordAt(caret))
	t.suggestionAdapter.Sort(partial)
}

func (t *CodeEditor) ShowSuggestionList() {
	if t.suggestionProvider == nil || t.IsSuggestionListShowing() {
		return
	}

	caret := t.controller.LastCaret()
	s, _ := t.controller.WordAt(caret)

	suggestions := t.suggestionProvider.SuggestionsAt(s)
	if len(suggestions) == 0 {
		t.HideSuggestionList()
		return
	}

	t.suggestionAdapter.SetSuggestions(suggestions)
	t.SortSuggestionList()
	child := t.AddChild(t.suggestionList)

	// Position the suggestion list below the last caret
	lineIdx := t.controller.LineIndex(caret)
	// TODO: What if the last caret is not visible?
	bounds := t.Size().Rect().Contract(t.Padding())
	line := t.Line(lineIdx)
	lineOffset := guix.ChildToParent(math.ZeroPoint, line, t.outer)
	target := line.PositionAt(caret).Add(lineOffset)
	cs := t.suggestionList.DesiredSize(math.ZeroSize, bounds.Size())
	t.suggestionList.Select(t.suggestionList.Adapter().ItemAt(0))
	t.suggestionList.SetSize(cs)
	child.Layout(cs.Rect().Offset(target).Intersect(bounds))
}

func (t *CodeEditor) HideSuggestionList() {
	if t.IsSuggestionListShowing() {
		t.RemoveChild(t.suggestionList)
	}
}

func (t *CodeEditor) Line(idx int) TextBoxLine {
	return guix.FindControl(t.ItemControl(idx).(guix.Parent), func(c guix.Control) bool {
		_, b := c.(TextBoxLine)
		return b
	}).(TextBoxLine)
}

// mixins.List overrides
func (t *CodeEditor) Click(ev guix.MouseEvent) (consume bool) {
	t.HideSuggestionList()
	return t.TextBox.Click(ev)
}

func (t *CodeEditor) KeyPress(ev guix.KeyboardEvent) (consume bool) {
	switch ev.Key {
	case guix.KeyTab:
		replace := true
		for _, sel := range t.controller.Selections() {
			s, e := sel.Range()
			if t.controller.LineIndex(s) != t.controller.LineIndex(e) {
				replace = false
				break
			}
		}
		switch {
		case replace:
			t.controller.ReplaceAll(strings.Repeat(" ", t.tabWidth))
			t.controller.Deselect(false)
		case ev.Modifier.Shift():
			t.controller.UnindentSelection(t.tabWidth)
		default:
			t.controller.IndentSelection(t.tabWidth)
		}
		return true
	case guix.KeySpace:
		if ev.Modifier.Control() {
			t.ShowSuggestionList()
			return
		}
	case guix.KeyUp:
		fallthrough
	case guix.KeyDown:
		if t.IsSuggestionListShowing() {
			return t.suggestionList.KeyPress(ev)
		}
	case guix.KeyLeft:
		t.HideSuggestionList()
	case guix.KeyRight:
		t.HideSuggestionList()
	case guix.KeyEnter:
		controller := t.controller
		if t.IsSuggestionListShowing() {
			text := t.suggestionAdapter.Suggestion(t.suggestionList.Selected()).Code()
			s, e := controller.WordAt(t.controller.LastCaret())
			controller.SetSelection(guix.CreateTextSelection(s, e, false))
			controller.ReplaceAll(text)
			controller.Deselect(false)
			t.HideSuggestionList()
		} else {
			t.controller.ReplaceWithNewlineKeepIndent()
		}
		return true
	case guix.KeyEscape:
		if t.IsSuggestionListShowing() {
			t.HideSuggestionList()
			return true
		}
	}
	return t.TextBox.KeyPress(ev)
}

func (t *CodeEditor) KeyStroke(ev guix.KeyStrokeEvent) (consume bool) {
	consume = t.TextBox.KeyStroke(ev)
	if t.IsSuggestionListShowing() {
		t.SortSuggestionList()
	}
	return
}

// mixins.TextBox overrides
func (t *CodeEditor) CreateLine(theme guix.Theme, index int) (TextBoxLine, guix.Control) {
	lineNumber := theme.CreateLabel()
	lineNumber.SetText(fmt.Sprintf("%.4d", index+1)) // Displayed lines start at 1

	line := &CodeEditorLine{}
	line.Init(line, theme, t, index)

	layout := theme.CreateLinearLayout()
	layout.SetDirection(guix.LeftToRight)
	layout.AddChild(lineNumber)
	layout.AddChild(line)

	return line, layout
}
