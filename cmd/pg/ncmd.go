package pg

import "github.com/Hayao0819/Abracadabra/cmd/pg/ncmd"

func init() {
	reg.Add(ncmd.Cmd())
}
