package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (server *MysqlAPIServer) endpointMetadata(c echo.Context) error {
	return c.JSON(http.StatusOK, server.api.GetDatabaseMetadata())
}

func (server *MysqlAPIServer) endpointQuery(c echo.Context) error {
	// need to query by type, consturct sql by metadata
	return c.String(http.StatusOK, "Not implement")
}
