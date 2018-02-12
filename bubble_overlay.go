// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package guix

import (
	"github.com/vcaesar/guix/math"
)

type BubbleOverlay interface {
	Control
	Show(control Control, target math.Point)
	Hide()
}
