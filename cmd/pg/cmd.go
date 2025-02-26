package pg

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

var reg = cobrautils.Registory{}

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "pg",
		Short: "Playground",
	}

	reg.Bind(&cmd)

	return &cmd
}
