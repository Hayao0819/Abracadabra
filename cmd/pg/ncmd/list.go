package ncmd

import (
	"github.com/Hayao0819/Abracadabra/notion"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list",
		Short: "List",
		RunE: func(cmd *cobra.Command, args []string) error {

			nc := notion.ShouldGetClient()

			me, err := nc.User.Me(cmd.Context())
			if err != nil {
				return err
			}

			cmd.Println(me)

			return nil
		},
	}

	return &cmd
}

func init() {
	reg.Add(listCmd())
}
