// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/math"
	"github.com/vcaesar/guix/mixins"
)

func CreateLabel(theme *Theme) guix.Label {
	l := &mixins.Label{}
	l.Init(l, theme, theme.DefaultFont(), theme.LabelStyle.FontColor)
	l.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	return l
}
