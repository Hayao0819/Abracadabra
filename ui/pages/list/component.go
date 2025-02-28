package list

import (
	"context"
	"fmt"
	"strings"

	"github.com/Hayao0819/Abracadabra/notion/nautils"
	"github.com/Hayao0819/Abracadabra/ui/common"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/londek/reactea"
	"github.com/samber/lo"
)

type Component struct {
	reactea.BasicComponent               // It implements AfterUpdate() for us, so we don't have to care!
	reactea.BasicPropfulComponent[Props] // It implements props backend - UpdateProps() and Props()

	pages    []*nautils.FullPage
	err      error
	selected int
}

type Props struct {
	common.ClientProps
}

func New() *Component {
	return &Component{
		pages:    nil,
		selected: 0,
	}
}

type updatePagesMsg struct {
	pages []*nautils.FullPage
	err   error
}

func (c *Component) updatePages() tea.Msg {
	ps, err := c.Props().Client.SearchFullPage(context.Background(), "")
	return updatePagesMsg{
		pages: ps,
		err:   err,
	}
}

func (c *Component) Init(props Props) tea.Cmd {
	c.UpdateProps(props)

	pages, err := props.Client.SearchFullPage(context.Background(), "")
	if err != nil {
		c.err = err
		return nil
	}
	c.pages = pages
	return c.updatePages
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return reactea.Destroy
		case "down":
			if c.selected < len(c.pages)-1 {
				c.selected++
			}
			return nil
		case "up":
			if c.selected > 0 {
				c.selected--
			}
			return nil
		case "enter":
			reactea.SetCurrentRoute("/detail/" + c.pages[c.selected].Page.ID.String())
			return nil
		}
	}
	return nil
}

func (c *Component) Render(int, int) string {
	titles := lo.Map(c.pages, func(p *nautils.FullPage, i int) string {
		if i == c.selected {
			return fmt.Sprintf(" > %s", p.Title.Title[0].PlainText)
		}
		return fmt.Sprintf("   %s", p.Title.Title[0].PlainText)
	})
	return strings.Join(titles, "\n") + "\n"
}
