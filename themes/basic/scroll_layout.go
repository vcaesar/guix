// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/mixins"
)

func CreateScrollLayout(theme *Theme) guix.ScrollLayout {
	l := &mixins.ScrollLayout{}
	l.Init(l, theme)
	return l
}
