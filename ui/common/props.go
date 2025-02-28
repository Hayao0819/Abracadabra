package common

import (
	"github.com/Hayao0819/Abracadabra/notion/nautils"
	"github.com/spf13/cobra"
)

type CobraProps struct {
	Cmd  *cobra.Command
	Args *[]string
}

type ClientProps struct {
	Client *nautils.Client
}
