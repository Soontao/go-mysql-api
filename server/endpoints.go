package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (m *MysqlAPIServer) endpointMetadata(c echo.Context) error {
	if c.QueryParam("simple") == "true" {
		return goJSON(c, http.StatusOK, m.api.GetDatabaseMetadata().GetSimpleMetadata())
	}
	return goJSON(c, http.StatusOK, m.api.GetDatabaseMetadata())
}

func (m *MysqlAPIServer) endpointEcho(c echo.Context) (err error) {
	bodyM, err := bodyMapOf(c)
	if err != nil {
		return err
	}
	return goJSONMessage(c, "echo api", bodyM)
}

func (m *MysqlAPIServer) endpointUpdateMetadata(c echo.Context) error {
	m.api.UpdateAPIMetadata()
	return goJSONMessage(c, "metadata refreshed", nil)
}

func (m *MysqlAPIServer) endpointTableGet(c echo.Context) (err error) {
	tableName := c.Param("table")
	limit, offset, fields, wheres, links := parseQueryParams(c)
	rs, err := m.api.Select(tableName, nil, limit, offset, fields, wheres, links)
	if err != nil {
		return err
	}
	return goJSONMessage(c, "get table", rs)
}

func (m *MysqlAPIServer) endpointTableGetSpecific(c echo.Context) (err error) {
	tableName := c.Param("table")
	id := c.Param("id")
	limit, offset, fields, wheres, links := parseQueryParams(c)
	rs, err := m.api.Select(tableName, id, limit, offset, fields, wheres, links)
	if err != nil {
		return err
	}
	return goJSONMessage(c, "get table by id", rs)
}

func (m *MysqlAPIServer) endpointTableCreate(c echo.Context) (err error) {
	payload, err := bodyMapOf(c)
	tableName := c.Param("table")
	if err != nil {
		return
	}
	rs, err := m.api.Create(tableName, payload)
	if err != nil {
		return err
	}
	msg, err := parseSQLResult(rs)
	if err != nil {
		return err
	}
	return goJSONMessage(c, "create record", msg)
}

func (m *MysqlAPIServer) endpointTableUpdateSpecific(c echo.Context) (err error) {
	payload, err := bodyMapOf(c)
	tableName := c.Param("table")
	id := c.Param("id")
	if err != nil {
		return
	}
	rs, err := m.api.Update(tableName, id, payload)
	if err != nil {
		return err
	}
	msg, err := parseSQLResult(rs)
	if err != nil {
		return err
	}
	return goJSONMessage(c, "update record", msg)
}

func (m *MysqlAPIServer) endpointTableDelete(c echo.Context) (err error) {
	payload, err := bodyMapOf(c)
	tableName := c.Param("table")
	if err != nil {
		return
	}
	rs, err := m.api.Delete(tableName, nil, payload)
	if err != nil {
		return err
	}
	msg, err := parseSQLResult(rs)
	if err != nil {
		return err
	}
	return goJSONMessage(c, "delete record", msg)
}

func (m *MysqlAPIServer) endpointTableDeleteSpecific(c echo.Context) (err error) {
	tableName := c.Param("table")
	id := c.Param("id")
	rs, err := m.api.Delete(tableName, id, nil)
	if err != nil {
		return err
	}
	msg, err := parseSQLResult(rs)
	if err != nil {
		return err
	}
	return goJSONMessage(c, "delete record by id", msg)
}
