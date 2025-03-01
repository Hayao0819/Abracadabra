package ncmd

import (
	"github.com/Hayao0819/Abracadabra/notion"
	"github.com/Hayao0819/Abracadabra/notion/nautils"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list",
		Short: "List",
		RunE: func(cmd *cobra.Command, args []string) error {

			nc := notion.ShouldGetClient()

			res, err := nc.SearchPage(cmd.Context(), "")
			if err != nil {
				return err
			}

			for _, p := range res {
				t := nautils.PageTitle(p)
				cmd.Printf("%s %s\n", t.Title[0].PlainText, p.URL)
			}

			return nil
		},
	}

	return &cmd
}

func init() {
	reg.Add(listCmd())
}
