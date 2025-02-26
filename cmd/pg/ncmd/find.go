package ncmd

import (
	"github.com/Hayao0819/Abracadabra/notion"
	"github.com/Hayao0819/Abracadabra/notion/nautils"
	"github.com/spf13/cobra"
)

func findCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "find",
		Short: "Find",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			nc := notion.ShouldGetClient()
			res, err := nc.SearchPage(cmd.Context(), args[0])
			if err != nil {
				return err
			}

			for _, p := range res {
				t := nautils.PageTitle(p)
				cmd.Println(t.Title[0].PlainText)
			}

			return nil
		},
	}

	return &cmd
}

func init() {
	reg.Add(findCmd())
}
