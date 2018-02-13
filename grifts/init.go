package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gustavo/emerald/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
