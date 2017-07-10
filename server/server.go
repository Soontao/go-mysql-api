package server

import (
	"github.com/Soontao/go-mysql-api/lib"
	"github.com/Soontao/go-mysql-api/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	server.e.HideBanner = true
	server.e.Logger = lib.Logger
	server.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[REQ] ${time_rfc3339_nano} ${method} (HTTP${status}) ${uri} ${latency}ns\n",
	}))
	server.api = mysql.NewMysqlAPI(dbURI)
	return server
}

// Start server
func (server *MysqlAPIServer) Start(address string) *MysqlAPIServer {
	server.e.GET("/api/metadata", server.endpointMetadata) // metadata
	server.e.POST("/api/echo", server.endpointEcho)        // echo api

	server.e.GET("/api/:table", server.endpointTableGet)       // Retrive
	server.e.PUT("/api/:table", server.endpointTableCreate)    // Create
	server.e.DELETE("/api/:table", server.endpointTableDelete) // Remove

	server.e.GET("/api/:table/:id", server.endpointTableGetSpecific)       // Retrive
	server.e.DELETE("/api/:table/:id", server.endpointTableDeleteSpecific) // Delete
	server.e.POST("/api/:table/:id", server.endpointTableUpdateSpecific)   // Update

	server.e.Logger.Infof("server start at %s", address)
	server.e.Logger.Fatal(server.e.Start(address))
	return server
}
