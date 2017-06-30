package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (server *MysqlAPIServer) endpointMetadata(c echo.Context) error {
	return c.JSON(http.StatusOK, server.api.GetDatabaseMetadata())
}
