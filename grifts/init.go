package grifts

import (
	"github.com/crolly/gorm_test/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
