package detail

import (
	"context"

	"github.com/Hayao0819/Abracadabra/notion/nautils"
	"github.com/Hayao0819/Abracadabra/ui/common"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jomei/notionapi"
	"github.com/londek/reactea"
)

type Component struct {
	reactea.BasicComponent               // It implements AfterUpdate() for us, so we don't have to care!
	reactea.BasicPropfulComponent[Props] // It implements props backend - UpdateProps() and Props()

	page *nautils.FullPage
}

type Props struct {
	common.ClientProps
	Page notionapi.PageID
}

func New() *Component {
	return &Component{}
}

func (c *Component) fetchPage() tea.Msg {
	page, err := c.Props().Client.FullPageFromID(context.Background(), c.Props().Page)
	if err != nil {
		return err
	}
	return page
}

func (c *Component) Init(props Props) tea.Cmd {
	c.UpdateProps(props)
	return nil
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return reactea.Destroy
		case "left":
			reactea.SetCurrentRoute("")
			return nil
		}
	}
	return nil
}
func (c *Component) Render(int, int) string {
	// return c.Props().Page.Title.Title[0].PlainText
	return c.page.Title.Title[0].PlainText
}
