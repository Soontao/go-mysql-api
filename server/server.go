package server

import (
	"github.com/Soontao/go-mysql-api/lib"
	"github.com/Soontao/go-mysql-api/mysql"
	"github.com/go-openapi/spec"
	"github.com/labstack/echo"
	"github.com/Soontao/go-mysql-api/inter"
)

// MysqlAPIServer is a http server could access mysql api
type MysqlAPIServer struct {
	e       *echo.Echo
	api     inter.IDatabaseAPI
	swagger *spec.Swagger
}

// NewMysqlAPIServer create a new MysqlAPIServer instance
func NewMysqlAPIServer(dbURI string, useInformationSchema bool) *MysqlAPIServer {
	server := &MysqlAPIServer{}
	server.e = echo.New()
	server.e.HTTPErrorHandler = customErrorHandler
	server.e.HideBanner = true
	server.e.Logger = lib.Logger
	server.e.Use(loggerMiddleware())
	server.api = mysql.NewMysqlAPI(dbURI, useInformationSchema)
	mountEndpoints(server.e, server.api)
	return server
}

// Start server
func (server *MysqlAPIServer) Start(address string) *MysqlAPIServer {
	server.StartMetadataRefreshCron()
	server.e.Logger.Infof("server start at %s", address)
	server.e.Logger.Fatal(server.e.Start(address))
	return server
}
