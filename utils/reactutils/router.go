package reactutils

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/londek/reactea"
	"github.com/londek/reactea/router"
)




func Route[TProps any](c reactea.Component[TProps], props TProps) func(p router.Params) (reactea.SomeComponent, tea.Cmd) {
	return func(p router.Params) (reactea.SomeComponent, tea.Cmd) {
		return c , c.Init(props)
	}
}
