package ncmd

import (
	"github.com/Hayao0819/Abracadabra/notion"
	"github.com/Hayao0819/Abracadabra/notion/nautils"
	"github.com/spf13/cobra"
)

func toMdCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "toMd",
		Short: "Convert to markdown",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			nc := notion.ShouldGetClient()

			res, err := nc.SearchPage(cmd.Context(), args[0])
			if err != nil {
				return err
			}

			targetPage := res[0]

			p, err := nc.PageBlocks(cmd.Context(), targetPage)
			if err != nil {
				return err
			}

			converter := nautils.HTMLConverter{}

			cmd.Println(converter.ToHTML(p))

			return nil
		},
	}

	return &cmd
}

func init() {
	reg.Add(toMdCmd())
}
