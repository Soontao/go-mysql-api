package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (m *MysqlAPIServer) endpointMetadata(c echo.Context) error {
	return goJSON(c, http.StatusOK, m.api.GetDatabaseMetadata())
}

func (m *MysqlAPIServer) endpointEcho(c echo.Context) (err error) {
	if m, err := bodyMapOf(c); err != nil {
		return err
	} else {
		return goJSONMessage(c, "echo api", m)
	}
}

func (m *MysqlAPIServer) endpointTableGet(c echo.Context) (err error) {
	if sql, err := m.api.SQL().GetByTable(c.Param("table"), c.QueryParams()); err != nil {
		return err
	} else if rs, err := m.api.Query(sql); err != nil {
		return err
	} else {
		return goJSONMessage(c, "get table", rs)
	}
}

func (m *MysqlAPIServer) endpointTableGetSpecific(c echo.Context) (err error) {
	if sql, err := m.api.SQL().GetByTableAndID(c.Param("table"), c.Param("id"), c.QueryParams()); err != nil {
		return err
	} else if rs, err := m.api.Query(sql); err != nil {
		return err
	} else {
		return goJSONMessage(c, "get table by id", rs)
	}
}

func (m *MysqlAPIServer) endpointTableDelete(c echo.Context) (err error) {
	payload, err := bodyMapOf(c)
	if sql, err := m.api.SQL().DeleteByTableWhere(c.Param("table"), payload); err != nil {
		return err
	} else if rs, err := m.api.Query(sql); err != nil {
		return err
	} else {
		return goJSONMessage(c, "get table by id", rs)
	}
}
