package server

import (
	"net/http"
	"github.com/Soontao/go-mysql-api/swagger"
	"github.com/labstack/echo"
	"github.com/Soontao/go-mysql-api/static"
	"github.com/Soontao/go-mysql-api/inter"
)

// mountEndpoints to echo server
func mountEndpoints(s *echo.Echo, api inter.IDatabaseAPI) {
	s.GET("/api/metadata", endpointMetadata(api)).Name = "Database Metadata"
	s.POST("/api/echo", endpointEcho).Name = "Echo API"
	s.GET("/api/endpoints", endpointServerEndpoints(s)).Name = "Server Endpoints"
	s.GET("/api/updatemetadata", endpointUpdateMetadata(api)).Name = "Update DB Metadata"
	s.GET("/api/swagger.json", endpointSwaggerJSON(api)).Name = "Swagger Infomation"
	s.GET("/api/swagger-ui.html", endpointSwaggerUI).Name = "Swagger UI"

	s.GET("/api/:table", endpointTableGet(api)).Name = "Retrive Some Records"
	s.PUT("/api/:table", endpointTableCreate(api)).Name = "Create Single Record"
	s.DELETE("/api/:table", endpointTableDelete(api)).Name = "Remove Some Records"

	s.GET("/api/:table/:id", endpointTableGetSpecific(api)).Name = "Retrive Record By ID"
	s.DELETE("/api/:table/:id", endpointTableDeleteSpecific(api)).Name = "Delete Record By ID"
	s.POST("/api/:table/:id", endpointTableUpdateSpecific(api)).Name = "Update Record By ID"

	s.PUT("/api/batch/:table", endpointBatchCreate(api)).Name = "Batch Create Records"
}

func endpointSwaggerUI(c echo.Context) error {
	return c.HTML(http.StatusOK, static.SWAGGER_UI_HTML)
}

func endpointSwaggerJSON(api inter.IDatabaseAPI) func(c echo.Context) error {
	return func(c echo.Context) error {
		s := swagger.GenSwaggerFromDBMetadata(api.GetDatabaseMetadata())
		s.Host = c.Request().Host
		s.Schemes = []string{c.Scheme()}
		return c.JSON(http.StatusOK, s)
	}
}

func endpointMetadata(api inter.IDatabaseAPI) func(c echo.Context) error {
	return func(c echo.Context) error {
		if c.QueryParam("simple") == "true" {
			return goJSON(c, http.StatusOK, api.GetDatabaseMetadata().GetSimpleMetadata())
		}
		return goJSON(c, http.StatusOK, api.GetDatabaseMetadata())
	}
}

func endpointEcho(c echo.Context) (err error) {
	bodyM, err := bodyMapOf(c)
	if err != nil {
		return err
	}
	return JSONMessage(c, "echo api", bodyM)
}

func endpointUpdateMetadata(api inter.IDatabaseAPI) func(c echo.Context) error {
	return func(c echo.Context) error {
		api.UpdateAPIMetadata()
		return JSONMessage(c, "metadata refreshed", nil)
	}
}

func endpointTableGet(api inter.IDatabaseAPI) func(c echo.Context) error {
	return func(c echo.Context) error {
		tableName := c.Param("table")
		option := parseQueryParamsNew(c)
		option.Table = tableName
		rs, err := api.Select(option)
		if err != nil {
			return err
		}
		return JSONMessage(c, "get table", rs)
	}
}

func endpointTableGetSpecific(api inter.IDatabaseAPI) func(c echo.Context) error {
	return func(c echo.Context) error {
		tableName := c.Param("table")
		id := c.Param("id")
		option := parseQueryParamsNew(c)
		option.Table = tableName
		option.Id = id
		rs, err := api.Select(option)
		if err != nil {
			return err
		}
		return JSONMessage(c, "get table by id", rs)
	}
}

func endpointTableCreate(api inter.IDatabaseAPI) func(c echo.Context) error {
	return func(c echo.Context) error {
		payload, err := bodyMapOf(c)
		tableName := c.Param("table")
		if err != nil {
			return err
		}
		rs, err := api.Create(tableName, payload)
		if err != nil {
			return err
		}
		msg, err := parseSQLResult(rs)
		if err != nil {
			return err
		}
		return JSONMessage(c, "create record", msg)
	}
}

func endpointTableUpdateSpecific(api inter.IDatabaseAPI) func(c echo.Context) error {
	return func(c echo.Context) error {
		payload, err := bodyMapOf(c)
		tableName := c.Param("table")
		id := c.Param("id")
		if err != nil {
			return err
		}
		rs, err := api.Update(tableName, id, payload)
		if err != nil {
			return err
		}
		msg, err := parseSQLResult(rs)
		if err != nil {
			return err
		}
		return JSONMessage(c, "update record", msg)
	}
}

func endpointTableDelete(api inter.IDatabaseAPI) func(c echo.Context) error {
	return func(c echo.Context) error {
		payload, err := bodyMapOf(c)
		tableName := c.Param("table")
		if err != nil {
			return err
		}
		rs, err := api.Delete(tableName, nil, payload)
		if err != nil {
			return err
		}
		msg, err := parseSQLResult(rs)
		if err != nil {
			return err
		}
		return JSONMessage(c, "delete record", msg)
	}
}

func endpointTableDeleteSpecific(api inter.IDatabaseAPI) func(c echo.Context) error {
	return func(c echo.Context) error {
		tableName := c.Param("table")
		id := c.Param("id")
		rs, err := api.Delete(tableName, id, nil)
		if err != nil {
			return err
		}
		msg, err := parseSQLResult(rs)
		if err != nil {
			return err
		}
		return JSONMessage(c, "delete record by id", msg)
	}
}

func endpointBatchCreate(api inter.IDatabaseAPI) func(c echo.Context) error {
	return func(c echo.Context) error {
		payload, err := bodySliceOf(c)
		msg := make([]map[string]interface{}, 0)
		tableName := c.Param("table")
		if err != nil {
			return err
		}
		for _, record := range payload {
			rs, err := api.Create(tableName, record.(map[string]interface{}))
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
}

func endpointServerEndpoints(e *echo.Echo) func(c echo.Context) error {
	return func(c echo.Context) error {
		return JSONMessage(c, "server endpoints", e.Routes())
	}
}
