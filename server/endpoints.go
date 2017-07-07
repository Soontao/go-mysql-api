package server

import (
	"net/http"
	"strconv"

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
	tableName := c.Param("table")
	limit, offset, fields := parseQueryParams(c)
	if rs, err := m.api.Select(tableName, nil, nil, limit, offset, fields); err != nil {
		return err
	} else {
		return goJSONMessage(c, "get table", rs)
	}
}

func (m *MysqlAPIServer) endpointTableGetSpecific(c echo.Context) (err error) {
	tableName := c.Param("table")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	limit, offset, fields := parseQueryParams(c)
	if rs, err := m.api.Select(tableName, id, nil, limit, offset, fields); err != nil {
		return err
	} else {
		return goJSONMessage(c, "get table by id", rs)
	}
}

func (m *MysqlAPIServer) endpointTableCreate(c echo.Context) (err error) {
	payload, err := bodyMapOf(c)
	tableName := c.Param("table")
	if err != nil {
		return
	}
	if rs, err := m.api.Create(tableName, payload); err != nil {
		return err
	} else {
		msg, err := parseSQLResult(rs)
		if err != nil {
			return err
		}
		return goJSONMessage(c, "create record", msg)
	}
}

func (m *MysqlAPIServer) endpointTableUpdate(c echo.Context) (err error) {
	payload, err := bodyMapOf(c)
	tableName := c.Param("table")
	if err != nil {
		return
	}
	if rs, err := m.api.Update(tableName, payload); err != nil {
		return err
	} else {
		msg, err := parseSQLResult(rs)
		if err != nil {
			return err
		}
		return goJSONMessage(c, "update record", msg)
	}
}

func (m *MysqlAPIServer) endpointTableDelete(c echo.Context) (err error) {
	payload, err := bodyMapOf(c)
	tableName := c.Param("table")
	if err != nil {
		return
	}
	if rs, err := m.api.Delete(tableName, nil, payload); err != nil {
		return err
	} else {
		msg, err := parseSQLResult(rs)
		if err != nil {
			return err
		}
		return goJSONMessage(c, "delete record", msg)
	}
}

func (m *MysqlAPIServer) endpointTableDeleteSpecific(c echo.Context) (err error) {
	tableName := c.Param("table")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	if rs, err := m.api.Delete(tableName, id, nil); err != nil {
		return err
	} else {
		msg, err := parseSQLResult(rs)
		if err != nil {
			return err
		}
		return goJSONMessage(c, "delete record by id", msg)
	}
}
