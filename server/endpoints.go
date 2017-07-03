package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (server *MysqlAPIServer) endpointMetadata(c echo.Context) error {
	return goJSON(c, http.StatusOK, server.api.GetDatabaseMetadata())
}

func (server *MysqlAPIServer) endpointEcho(c echo.Context) (err error) {
	if m, err := bodyMapOf(c); err != nil {
		return err
	} else {
		return goJSONMessage(c, "echo api", m)
	}
}

func (this *MysqlAPIServer) endpointTableGet(c echo.Context) (err error) {
	if sql, err := this.api.SQL().GetByTable(c.Param("table"), c.QueryParams()); err != nil {
		return err
	} else if rs, err := this.api.Query(sql); err != nil {
		return err
	} else {
		return goJSONMessage(c, "get table", rs)
	}
}

func (this *MysqlAPIServer) endpointTableGetSpecific(c echo.Context) (err error) {
	if sql, err := this.api.SQL().GetByTableAndID(c.Param("table"), c.Param("id"), c.QueryParams()); err != nil {
		return err
	} else if rs, err := this.api.Query(sql); err != nil {
		return err
	} else {
		return goJSONMessage(c, "get table by id", rs)
	}
}
