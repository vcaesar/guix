// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"github.com/vcaesar/guix"
)

type SuggestionAdapter struct {
	guix.FilteredListAdapter
}

func (a *SuggestionAdapter) SetSuggestions(suggestions []guix.CodeSuggestion) {
	items := make([]guix.FilteredListItem, len(suggestions))
	for i, s := range suggestions {
		items[i].Name = s.Name()
		items[i].Data = s
	}
	a.SetItems(items)
}

func (a *SuggestionAdapter) Suggestion(item guix.AdapterItem) guix.CodeSuggestion {
	return item.(guix.FilteredListItem).Data.(guix.CodeSuggestion)
}
