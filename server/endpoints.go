package server

import (
	"net/http"

	"github.com/labstack/echo"
	// mysql dialect
	_ "gopkg.in/doug-martin/goqu.v4/adapters/mysql"
)

func (server *MysqlAPIServer) endpointMetadata(c echo.Context) error {
	return goJSON(c, http.StatusOK, server.api.GetDatabaseMetadata())
}

func (server *MysqlAPIServer) endpointQuery(c echo.Context) error {
	// need to query by type, consturct sql by metadata
	return c.String(http.StatusOK, "Not implement")
}

func (server *MysqlAPIServer) endpointTableGet(c echo.Context) error {
	tableName := c.Param("table")
	if sql, _, err := server.sql.From(tableName).ToSql(); err != nil {
		return err
	} else if rs, err := server.api.Query(sql); err != nil {
		return err
	} else {
		return goJSON(c, http.StatusOK, rs)
	}
}
