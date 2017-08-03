package main

import (
	"github.com/Soontao/go-mysql-api/server"
	"github.com/mkideal/cli"
)

type cliArgs struct {
	cli.Helper
	ConnectionStr      string `cli:"*c,*conn" usage:"mysql connection str" dft:"$API_CONN_STR"`
	ListenAddress      string `cli:"*l,*listen" usage:"listen host and port" dft:"$API_HOST_LS"`
	NoInfomationSchema bool `cli:"n,noinfo" usage:"dont use mysql information shcema" dft:"$API_NO_USE_INFO"`
}

func main() {
	cli.Run(new(cliArgs), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*cliArgs)
		server.NewMysqlAPIServer(argv.ConnectionStr, !argv.NoInfomationSchema).Start(argv.ListenAddress)
		return nil
	})

}
