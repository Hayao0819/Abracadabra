package cmd

import (
	"github.com/Hayao0819/Abracadabra/conf"
	"github.com/Hayao0819/Abracadabra/notion"
	"github.com/Hayao0819/Abracadabra/ui"
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/londek/reactea"
	"github.com/spf13/cobra"
)

var reg = cobrautils.Registory{}

func rootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:          "abr",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

			if err := conf.Init(); err != nil {
				return err
			}

			if err := notion.Init(); err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := notion.GetClient()
			if err != nil {
				return err
			}

			app := reactea.NewProgram(ui.New(client))
			_, err = app.Run()
			if err != nil {
				return err
			}
			return nil
		},
	}

	reg.Bind(&cmd)

	return &cmd
}
