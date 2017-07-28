package server

import (
	"github.com/Soontao/go-mysql-api/lib"
	"github.com/Soontao/go-mysql-api/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/go-openapi/spec"
)

// MysqlAPIServer is a http server could access mysql api
type MysqlAPIServer struct {
	e       *echo.Echo
	api     *mysql.MysqlAPI
	swagger *spec.Swagger
}

// NewMysqlAPIServer create a new MysqlAPIServer instance
func NewMysqlAPIServer(dbURI string, useInformationSchema bool) *MysqlAPIServer {
	server := &MysqlAPIServer{}
	server.e = echo.New()
	server.e.HTTPErrorHandler = customErrorHandler
	server.e.HideBanner = true
	server.e.Logger = lib.Logger
	server.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[REQ] ${time_rfc3339_nano} ${method} (HTTP${status}) ${uri} ${latency}ns\n",
	}))
	server.api = mysql.NewMysqlAPI(dbURI, useInformationSchema)
	return server
}

// Start server
func (server *MysqlAPIServer) Start(address string) *MysqlAPIServer {
	server.e.GET("/api/metadata", server.endpointMetadata).Name = "Database Metadata"
	server.e.POST("/api/echo", server.endpointEcho).Name = "Echo API"
	server.e.GET("/api/endpoints", server.endpointServerEndpoints).Name = "Server Endpoints"
	server.e.GET("/api/updatemetadata", server.endpointUpdateMetadata).Name = "Update DB Metadata"
	server.e.GET("/api/swagger-ui.html", server.endpointSwaggerHTML).Name = "Swagger UI Page"
	server.e.GET("/api/swagger.json", server.endpointSwaggerJSON).Name = "Swagger Infomation"

	server.e.GET("/api/:table", server.endpointTableGet).Name = "Retrive Some Records"
	server.e.PUT("/api/:table", server.endpointTableCreate).Name = "Create Single Record"
	server.e.DELETE("/api/:table", server.endpointTableDelete).Name = "Remove Some Records"

	server.e.GET("/api/:table/:id", server.endpointTableGetSpecific).Name = "Retrive Record By ID"
	server.e.DELETE("/api/:table/:id", server.endpointTableDeleteSpecific).Name = "Delete Record By ID"
	server.e.POST("/api/:table/:id", server.endpointTableUpdateSpecific).Name = "Update Record By ID"

	server.e.PUT("/api/batch/:table", server.endpointBatchCreate).Name = "Batch Create Record"

	server.e.Logger.Infof("server start at %s", address)
	server.e.Logger.Fatal(server.e.Start(address))
	return server
}
