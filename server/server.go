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
	server.e.HTTPErrorHandler = customErrorHandler
	server.api = mysql.NewMysqlAPI(dbURI)
	return server
}

// Start server
func (server *MysqlAPIServer) Start(address string) *MysqlAPIServer {
	server.e.GET("/api/metadata", server.endpointMetadata)
	server.e.POST("/api/echo", server.endpointEcho)
	server.e.GET("/api/:table", server.endpointTableGet)
	server.e.GET("/api/:table/:id", server.endpointTableGetSpecific)
	server.e.Logger.Fatal(server.e.Start(address))
	return server
}
