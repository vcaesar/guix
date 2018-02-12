// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import (
	"github.com/vcaesar/guix"
)

type Style struct {
	FontColor guix.Color
	Brush     guix.Brush
	Pen       guix.Pen
}

func CreateStyle(fontColor, brushColor, penColor guix.Color, penWidth float32) Style {
	return Style{
		FontColor: fontColor,
		Pen:       guix.CreatePen(penWidth, penColor),
		Brush:     guix.CreateBrush(brushColor),
	}
}
