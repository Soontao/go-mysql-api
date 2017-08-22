package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gopkg.in/doug-martin/goqu.v4"

	"reflect"

	"database/sql"

	"regexp"

	"github.com/labstack/echo"
	"github.com/mediocregopher/gojson"
	"strings"
	"github.com/Soontao/go-mysql-api/key"
	types    "github.com/Soontao/go-mysql-api/t"
	"github.com/labstack/echo/middleware"
)

// Message
type Message struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func loggerMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[REQ] ${time_rfc3339_nano} ${method} (HTTP${status}) ${uri} ${latency}ns\n",
	})
}

func goJSON(c echo.Context, code int, i interface{}) error {
	_, pretty := c.QueryParams()["pretty"]
	if pretty {
		return goJSONPretty(c, code, i, "  ")
	}
	b, err := gojson.Marshal(i)
	if err != nil {
		return err
	}
	return c.JSONBlob(code, b)
}

func JSONMessage(c echo.Context, m string, i interface{}) error {
	return goJSON(c, http.StatusOK, &Message{http.StatusOK, m, i})
}

func goJSONPretty(c echo.Context, code int, i interface{}, indent string) (err error) {
	b, err := gojson.MarshalIndent(i, "", indent)
	if err != nil {
		return
	}
	return c.JSONBlob(code, b)
}

func bodyMapOf(c echo.Context) (jsonMap map[string]interface{}, err error) {
	jsonMap = make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	return jsonMap, err
}

func bodySliceOf(c echo.Context) (jsonSlice []interface{}, err error) {
	jsonSlice = make([]interface{}, 0)
	err = json.NewDecoder(c.Request().Body).Decode(&jsonSlice)
	return
}

func customErrorHandler(err error, c echo.Context) {
	if reflect.TypeOf(err) == reflect.TypeOf(&echo.HTTPError{}) {
		httpError := err.(*echo.HTTPError)
		goJSON(c, httpError.Code, &Message{httpError.Code, httpError.Message.(string), nil})
	} else {
		goJSON(c, http.StatusInternalServerError, &Message{http.StatusInternalServerError, err.Error(), nil})
	}
}

func parseQueryParamsNew(c echo.Context) (option types.QueryOption) {
	option = types.QueryOption{}
	queryParam := c.QueryParams()
	option.Limit, option.Offset, option.Fields, option.Wheres, option.Links = parseQueryParams(c)
	if queryParam[key.KEY_QUERY_SEARCH] != nil {
		searchStrArray := queryParam[key.KEY_QUERY_SEARCH]
		if searchStrArray[0] != "" {
			option.Search = searchStrArray[0]
		}
	}
	return
}

func parseQueryParams(c echo.Context) (limit int, offset int, fields []string, wheres map[string]goqu.Op, links []string) {
	queryParam := c.QueryParams()
	limit, _ = strconv.Atoi(c.QueryParam(key.KEY_QUERY_LIMIT)) // _limit
	offset, _ = strconv.Atoi(c.QueryParam(key.KEY_QUERY_SKIP)) // _skip
	fields = make([]string, 0)
	if queryParam[key.KEY_QUERY_FIELDS] != nil { // _fields
		for _, sArrFields := range queryParam[key.KEY_QUERY_FIELDS] {
			fields = append(fields, strings.Split(sArrFields, ",")...)
		}
	}
	if queryParam[key.KEY_QUERY_FIELD] != nil { // _field
		for _, f := range queryParam[key.KEY_QUERY_FIELD] {
			fields = append(fields, f)
		}
	}
	if queryParam[key.KEY_QUERY_LINK] != nil { // _link
		links = make([]string, len(queryParam[key.KEY_QUERY_LINK]))
		for idx, f := range queryParam[key.KEY_QUERY_LINK] {
			links[idx] = f
		}
	}
	r := regexp.MustCompile("\\'(.*?)\\'\\.([\\w]+)\\((.*?)\\)")
	if queryParam[key.KEY_QUERY_WHERE] != nil {
		wheres = make(map[string]goqu.Op)
		for _, sWhere := range queryParam[key.KEY_QUERY_WHERE] {
			arr := r.FindStringSubmatch(sWhere)
			if len(arr) == 4 {
				switch arr[2] {
				case "in", "notIn":
					wheres[arr[1]] = goqu.Op{arr[2]: strings.Split(arr[3], ",")}
				case "like", "is", "neq", "isNot", "eq":
					wheres[arr[1]] = goqu.Op{arr[2]: arr[3]}
				}

			}
		}
	}
	return
}

func parseSQLResult(rs sql.Result) (rt map[string]interface{}, err error) {
	lastInsertID, err := rs.LastInsertId()
	rowesAffected, err := rs.RowsAffected()
	rt = map[string]interface{}{"lastInsertID": lastInsertID, "rowesAffected": rowesAffected}
	return
}
