package ncmd

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

var reg = cobrautils.Registory{}

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "ncmd",
		Short: "Notion Playground",
	}

	reg.Bind(&cmd)

	return &cmd
}
