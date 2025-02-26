package ncmd

import (
	"github.com/Hayao0819/Abracadabra/notion"
	"github.com/spf13/cobra"
)

func meCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "me",
		Short: "Retribute me",
		RunE: func(cmd *cobra.Command, args []string) error {

			nc := notion.ShouldGetClient()

			me, err := nc.RawClient().User.Me(cmd.Context())
			if err != nil {
				return err
			}

			cmd.Printf("ID: %s\n", me.ID)
			cmd.Printf("Name: %s\n", me.Name)
			// cmd.Printf("Email: %s\n", me.Person.Email)

			return nil
		},
	}

	return &cmd
}

func init() {
	reg.Add(meCmd())
}
