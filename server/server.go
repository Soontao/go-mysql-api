package server

import (
	"github.com/Soontao/go-mysql-api/server/lib"
	"github.com/labstack/echo"
	"github.com/Soontao/go-mysql-api/adapter"
)

// MysqlAPIServer is a http server could access mysql api
type MysqlAPIServer struct {
	*echo.Echo               // echo web server
	api adapter.IDatabaseAPI // database api adapter
}

// New create a new MysqlAPIServer instance
func New(api adapter.IDatabaseAPI) *MysqlAPIServer {
	server := &MysqlAPIServer{}
	server.Echo = echo.New()
	server.HTTPErrorHandler = customErrorHandler
	server.HideBanner = true
	server.Logger = lib.Logger
	server.Use(loggerMiddleware())
	server.api = api
	mountEndpoints(server.Echo, server.api)
	return server
}

// Start server
func (server *MysqlAPIServer) Start(address string) *MysqlAPIServer {
	server.StartMetadataRefreshCron()
	server.Logger.Infof("server start at %s", address)
	server.Logger.Fatal(server.Echo.Start(address))
	return server
}
