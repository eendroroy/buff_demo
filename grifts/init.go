package grifts

import (
	"github.com/eendroroy/buff_demo/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
