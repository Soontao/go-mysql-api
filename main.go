package main

import "github.com/Soontao/go-mysql-api/server"

func main() {
	server.
		NewMysqlAPIServer("monitor:yn0Mbx1mPcZWlvzb@tcp(stu.ecs.fornever.org:3306)/monitor").
		Start(":1323")
}
