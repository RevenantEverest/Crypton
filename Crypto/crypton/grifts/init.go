package grifts

import (
	"github.com/ahmede7th/crypton/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
