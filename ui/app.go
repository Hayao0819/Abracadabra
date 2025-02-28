package ui

import (
	"fmt"

	"github.com/Hayao0819/Abracadabra/notion/nautils"
	"github.com/Hayao0819/Abracadabra/ui/common"
	"github.com/Hayao0819/Abracadabra/ui/pages/list"
	"github.com/Hayao0819/Abracadabra/utils/reactutils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/londek/reactea"

	"github.com/londek/reactea/router"
)

type Component struct {
	reactea.BasicComponent
	reactea.BasicPropfulComponent[reactea.NoProps]

	mainRouter reactea.Component[router.Props]
	client     *nautils.Client
}

func New(client *nautils.Client) *Component {
	return &Component{
		mainRouter: router.New(),
		client:     client,
	}
}

func (c *Component) Init(reactea.NoProps) tea.Cmd {
	return c.mainRouter.Init(map[string]router.RouteInitializer{
		"default": reactutils.Route(list.New(), list.Props{
			ClientProps: common.ClientProps{
				Client: c.client,
			},
		}),
		"/detail/:id": func(p router.Params) (reactea.SomeComponent, tea.Cmd) {
			// c := detail.New()
			// p := detail.Props{

			// }
			fmt.Println("detail")
			return nil, nil
		},
	})
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return reactea.Destroy
		}
	}

	return c.mainRouter.Update(msg)
}

func (c *Component) Render(width, height int) string {
	return c.mainRouter.Render(width, height)
}
