package server

import (
	"github.com/Soontao/go-mysql-api/mysql"
	"github.com/labstack/echo"
)

// MysqlAPIServer is a http server could access mysql api
type MysqlAPIServer struct {
	e   *echo.Echo
	api *mysql.MysqlAPI
}

// NewMysqlAPIServer create a new MysqlAPIServer instance
func NewMysqlAPIServer(dbURI string) *MysqlAPIServer {
	server := &MysqlAPIServer{}
	server.e = echo.New()
	server.api = mysql.NewMysqlAPI(dbURI)
	return server
}

// Start server
func (server *MysqlAPIServer) Start(address string) *MysqlAPIServer {
	server.e.GET("/metadata", server.endpointMetadata)
	server.e.Logger.Fatal(server.e.Start(address))
	return server
}
