// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

import (
	"github.com/vcaesar/guix"

	"github.com/goxjs/glfw"
)

func translateKeyboardKey(in glfw.Key) guix.KeyboardKey {
	switch in {
	case glfw.KeySpace:
		return guix.KeySpace
	case glfw.KeyApostrophe:
		return guix.KeyApostrophe
	case glfw.KeyComma:
		return guix.KeyComma
	case glfw.KeyMinus:
		return guix.KeyMinus
	case glfw.KeyPeriod:
		return guix.KeyPeriod
	case glfw.KeySlash:
		return guix.KeySlash
	case glfw.Key0:
		return guix.Key0
	case glfw.Key1:
		return guix.Key1
	case glfw.Key2:
		return guix.Key2
	case glfw.Key3:
		return guix.Key3
	case glfw.Key4:
		return guix.Key4
	case glfw.Key5:
		return guix.Key5
	case glfw.Key6:
		return guix.Key6
	case glfw.Key7:
		return guix.Key7
	case glfw.Key8:
		return guix.Key8
	case glfw.Key9:
		return guix.Key9
	case glfw.KeySemicolon:
		return guix.KeySemicolon
	case glfw.KeyEqual:
		return guix.KeyEqual
	case glfw.KeyA:
		return guix.KeyA
	case glfw.KeyB:
		return guix.KeyB
	case glfw.KeyC:
		return guix.KeyC
	case glfw.KeyD:
		return guix.KeyD
	case glfw.KeyE:
		return guix.KeyE
	case glfw.KeyF:
		return guix.KeyF
	case glfw.KeyG:
		return guix.KeyG
	case glfw.KeyH:
		return guix.KeyH
	case glfw.KeyI:
		return guix.KeyI
	case glfw.KeyJ:
		return guix.KeyJ
	case glfw.KeyK:
		return guix.KeyK
	case glfw.KeyL:
		return guix.KeyL
	case glfw.KeyM:
		return guix.KeyM
	case glfw.KeyN:
		return guix.KeyN
	case glfw.KeyO:
		return guix.KeyO
	case glfw.KeyP:
		return guix.KeyP
	case glfw.KeyQ:
		return guix.KeyQ
	case glfw.KeyR:
		return guix.KeyR
	case glfw.KeyS:
		return guix.KeyS
	case glfw.KeyT:
		return guix.KeyT
	case glfw.KeyU:
		return guix.KeyU
	case glfw.KeyV:
		return guix.KeyV
	case glfw.KeyW:
		return guix.KeyW
	case glfw.KeyX:
		return guix.KeyX
	case glfw.KeyY:
		return guix.KeyY
	case glfw.KeyZ:
		return guix.KeyZ
	case glfw.KeyLeftBracket:
		return guix.KeyLeftBracket
	case glfw.KeyBackslash:
		return guix.KeyBackslash
	case glfw.KeyRightBracket:
		return guix.KeyRightBracket
	case glfw.KeyGraveAccent:
		return guix.KeyGraveAccent
	case glfw.KeyWorld1:
		return guix.KeyWorld1
	case glfw.KeyWorld2:
		return guix.KeyWorld2
	case glfw.KeyEscape:
		return guix.KeyEscape
	case glfw.KeyEnter:
		return guix.KeyEnter
	case glfw.KeyTab:
		return guix.KeyTab
	case glfw.KeyBackspace:
		return guix.KeyBackspace
	case glfw.KeyInsert:
		return guix.KeyInsert
	case glfw.KeyDelete:
		return guix.KeyDelete
	case glfw.KeyRight:
		return guix.KeyRight
	case glfw.KeyLeft:
		return guix.KeyLeft
	case glfw.KeyDown:
		return guix.KeyDown
	case glfw.KeyUp:
		return guix.KeyUp
	case glfw.KeyPageUp:
		return guix.KeyPageUp
	case glfw.KeyPageDown:
		return guix.KeyPageDown
	case glfw.KeyHome:
		return guix.KeyHome
	case glfw.KeyEnd:
		return guix.KeyEnd
	case glfw.KeyCapsLock:
		return guix.KeyCapsLock
	case glfw.KeyScrollLock:
		return guix.KeyScrollLock
	case glfw.KeyNumLock:
		return guix.KeyNumLock
	case glfw.KeyPrintScreen:
		return guix.KeyPrintScreen
	case glfw.KeyPause:
		return guix.KeyPause
	case glfw.KeyF1:
		return guix.KeyF1
	case glfw.KeyF2:
		return guix.KeyF2
	case glfw.KeyF3:
		return guix.KeyF3
	case glfw.KeyF4:
		return guix.KeyF4
	case glfw.KeyF5:
		return guix.KeyF5
	case glfw.KeyF6:
		return guix.KeyF6
	case glfw.KeyF7:
		return guix.KeyF7
	case glfw.KeyF8:
		return guix.KeyF8
	case glfw.KeyF9:
		return guix.KeyF9
	case glfw.KeyF10:
		return guix.KeyF10
	case glfw.KeyF11:
		return guix.KeyF11
	case glfw.KeyF12:
		return guix.KeyF12
	case glfw.KeyKP0:
		return guix.KeyKp0
	case glfw.KeyKP1:
		return guix.KeyKp1
	case glfw.KeyKP2:
		return guix.KeyKp2
	case glfw.KeyKP3:
		return guix.KeyKp3
	case glfw.KeyKP4:
		return guix.KeyKp4
	case glfw.KeyKP5:
		return guix.KeyKp5
	case glfw.KeyKP6:
		return guix.KeyKp6
	case glfw.KeyKP7:
		return guix.KeyKp7
	case glfw.KeyKP8:
		return guix.KeyKp8
	case glfw.KeyKP9:
		return guix.KeyKp9
	case glfw.KeyKPDecimal:
		return guix.KeyKpDecimal
	case glfw.KeyKPDivide:
		return guix.KeyKpDivide
	case glfw.KeyKPMultiply:
		return guix.KeyKpMultiply
	case glfw.KeyKPSubtract:
		return guix.KeyKpSubtract
	case glfw.KeyKPAdd:
		return guix.KeyKpAdd
	case glfw.KeyKPEnter:
		return guix.KeyKpEnter
	case glfw.KeyKPEqual:
		return guix.KeyKpEqual
	case glfw.KeyLeftShift:
		return guix.KeyLeftShift
	case glfw.KeyLeftControl:
		return guix.KeyLeftControl
	case glfw.KeyLeftAlt:
		return guix.KeyLeftAlt
	case glfw.KeyLeftSuper:
		return guix.KeyLeftSuper
	case glfw.KeyRightShift:
		return guix.KeyRightShift
	case glfw.KeyRightControl:
		return guix.KeyRightControl
	case glfw.KeyRightAlt:
		return guix.KeyRightAlt
	case glfw.KeyRightSuper:
		return guix.KeyRightSuper
	case glfw.KeyMenu:
		return guix.KeyMenu
	default:
		return guix.KeyUnknown
	}
}

func translateKeyboardModifier(in glfw.ModifierKey) guix.KeyboardModifier {
	out := guix.ModNone
	if in&glfw.ModShift != 0 {
		out |= guix.ModShift
	}
	if in&glfw.ModControl != 0 {
		out |= guix.ModControl
	}
	if in&glfw.ModAlt != 0 {
		out |= guix.ModAlt
	}
	if in&glfw.ModSuper != 0 {
		out |= guix.ModSuper
	}
	return out
}
