package cmd

import "github.com/Hayao0819/Abracadabra/cmd/pg"

func init() {
	reg.Add(pg.Cmd())
}
