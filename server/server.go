package server

import (
	"github.com/Soontao/go-mysql-api/mysql"
	"github.com/labstack/echo"
	"gopkg.in/doug-martin/goqu.v4"
)

// MysqlAPIServer is a http server could access mysql api
type MysqlAPIServer struct {
	e   *echo.Echo
	api *mysql.MysqlAPI
	sql *goqu.Database
}

// NewMysqlAPIServer create a new MysqlAPIServer instance
func NewMysqlAPIServer(dbURI string) *MysqlAPIServer {
	server := &MysqlAPIServer{}
	server.e = echo.New()
	server.api = mysql.NewMysqlAPI(dbURI)
	server.sql = goqu.New("mysql", server.api.Connection())
	return server
}

// Start server
func (server *MysqlAPIServer) Start(address string) *MysqlAPIServer {
	server.e.GET("/api/metadata", server.endpointMetadata)
	server.e.GET("/api/:table", server.endpointTableGet)
	server.e.Logger.Fatal(server.e.Start(address))
	return server
}
