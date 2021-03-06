package basic

import (
	"github.com/vcaesar/guix"
	"github.com/vcaesar/guix/mixins"
)

func CreateTableLayout(theme *Theme) guix.TableLayout {
	l := &mixins.TableLayout{}
	l.Init(l, theme)
	return l
}
