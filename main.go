package main

import (
	"github.com/Soontao/go-mysql-api/server"
	"github.com/mkideal/cli"
)

type cliArgs struct {
	cli.Helper
	ConnectionStr string `cli:"*c,*conn" usage:"mysql connection str" dft:"$CONN_STR"`
	ListenAddress string `cli:"*l,*listen" usage:"listen host and port" dft:"0.0.0.0:1323"`
}

func main() {
	cli.Run(new(cliArgs), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*cliArgs)
		server.
			NewMysqlAPIServer(argv.ConnectionStr).
			Start(argv.ListenAddress)
		return nil
	})

}
