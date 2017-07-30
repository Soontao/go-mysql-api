package server

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/Soontao/go-mysql-api/swagger"
	"github.com/labstack/echo"
)

func (m *MysqlAPIServer) getStaticEndPoint() echo.HandlerFunc {
	assests := http.FileServer(rice.MustFindBox("../static").HTTPBox())
	return echo.WrapHandler(http.StripPrefix("/static/", assests))
}

func (m *MysqlAPIServer) endpointSwaggerJSON(c echo.Context) error {
	s := swagger.GenSwaggerFromDBMetadata(m.api.GetDatabaseMetadata())
	s.Host = c.Request().Host
	s.Schemes = []string{c.Scheme()}
	return c.JSON(http.StatusOK, s)
}

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
	return JSONMessage(c, "echo api", bodyM)
}

func (m *MysqlAPIServer) endpointUpdateMetadata(c echo.Context) error {
	m.api.UpdateAPIMetadata()
	return JSONMessage(c, "metadata refreshed", nil)
}

func (m *MysqlAPIServer) endpointTableGet(c echo.Context) (err error) {
	tableName := c.Param("table")
	limit, offset, fields, wheres, links := parseQueryParams(c)
	rs, err := m.api.Select(tableName, nil, limit, offset, fields, wheres, links)
	if err != nil {
		return err
	}
	return JSONMessage(c, "get table", rs)
}

func (m *MysqlAPIServer) endpointTableGetSpecific(c echo.Context) (err error) {
	tableName := c.Param("table")
	id := c.Param("id")
	limit, offset, fields, wheres, links := parseQueryParams(c)
	rs, err := m.api.Select(tableName, id, limit, offset, fields, wheres, links)
	if err != nil {
		return err
	}
	return JSONMessage(c, "get table by id", rs)
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
	return JSONMessage(c, "create record", msg)
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
	return JSONMessage(c, "update record", msg)
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
	return JSONMessage(c, "delete record", msg)
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
	return JSONMessage(c, "delete record by id", msg)
}

func (m *MysqlAPIServer) endpointBatchCreate(c echo.Context) (err error) {
	payload, err := bodySliceOf(c)
	msg := make([]map[string]interface{}, 0)
	tableName := c.Param("table")
	if err != nil {
		return
	}
	for _, record := range payload {
		rs, err := m.api.Create(tableName, record.(map[string]interface{}))
		var r_msg map[string]interface{}
		if err != nil {
			r_msg = map[string]interface{}{"error": err}
		} else {
			r_msg, _ = parseSQLResult(rs)
		}
		msg = append(msg, r_msg)
	}
	return JSONMessage(c, "batch create record", msg)
}

func (m *MysqlAPIServer) endpointServerEndpoints(c echo.Context) (err error) {
	return JSONMessage(c, "server endpoints", m.e.Routes())
}
