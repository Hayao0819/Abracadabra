package list

// import (
// 	"context"
// 	"fmt"
// 	"strings"

// 	"github.com/Hayao0819/Abracadabra/notion/nautils"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/samber/lo"
// 	"github.com/spf13/cobra"
// )

// type listModel struct {
// 	cmd      *cobra.Command
// 	args     []string
// 	err      error
// 	client   *nautils.Client
// 	selected int
// 	pages    []*nautils.FullPage
// }

// type updatePagesMsg struct {
// 	pages []*nautils.FullPage
// 	err   error
// }

// func (m listModel) updatePages() tea.Msg {
// 	ps, err := m.client.SearchFullPage(context.Background(), "")
// 	return updatePagesMsg{
// 		pages: ps,
// 		err:   err,
// 	}
// }

// func newListModel(cmd *cobra.Command, args []string, client *nautils.Client) *listModel {
// 	return &listModel{
// 		cmd:      cmd,
// 		args:     args,
// 		err:      nil,
// 		client:   client,
// 		selected: 0,
// 	}
// }

// func (m listModel) Init() tea.Cmd {
// 	return m.updatePages
// }

// func (m listModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := message.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "q", "esc", "ctrl+c":
// 			return m, tea.Quit
// 		case "down":
// 			if m.selected < len(m.pages)-1 {
// 				m.selected++
// 			}
// 			return m, nil
// 		case "up":
// 			if m.selected > 0 {
// 				m.selected--
// 			}
// 			return m, nil
// 		}
	
// 	case updatePagesMsg:
// 		m.err = msg.err
// 		m.pages = msg.pages
// 		return m, nil
// 	}

// 	return m, nil
// }

// func (m listModel) View() string {
// 	titles := lo.Map(m.pages, func(p *nautils.FullPage, i int) string {
// 		if i == m.selected {
// 			return fmt.Sprintf(" > %s", p.Title.Title[0].PlainText)
// 		}
// 		return fmt.Sprintf("   %s", p.Title.Title[0].PlainText)
// 	})
// 	return strings.Join(titles, "\n") + "\n"
// }
