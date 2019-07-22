package grifts

import (
	"github.com/dtmkeng/framework-exmaple/coke/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
