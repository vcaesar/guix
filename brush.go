// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package guix

var (
	WhiteBrush       = CreateBrush(White)
	TransparentBrush = CreateBrush(Transparent)
	BlackBrush       = CreateBrush(Black)
	DefaultBrush     = WhiteBrush
)

type Brush struct {
	Color Color
}

func CreateBrush(color Color) Brush {
	return Brush{color}
}
