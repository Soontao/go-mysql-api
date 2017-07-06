package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"reflect"

	"github.com/labstack/echo"
	"github.com/mediocregopher/gojson"
)

// Message
type Message struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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

func goJSONMessage(c echo.Context, m string, i interface{}) error {
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

func customErrorHandler(err error, c echo.Context) {
	if reflect.TypeOf(err) == reflect.TypeOf(&echo.HTTPError{}) {
		httpError := err.(*echo.HTTPError)
		goJSON(c, httpError.Code, &Message{httpError.Code, httpError.Message.(string), nil})
	} else {
		goJSON(c, http.StatusInternalServerError, &Message{http.StatusInternalServerError, err.Error(), nil})
	}
}

func parseQueryParams(c echo.Context) (limit int, offset int, fields []interface{}) {
	queryParam := c.QueryParams()
	limit, _ = strconv.Atoi(c.QueryParam("_limit"))
	offset, _ = strconv.Atoi(c.QueryParam("_skip"))
	if queryParam["_field"] != nil {
		fields = make([]interface{}, len(queryParam["_field"]))
		for idx, f := range queryParam["_field"] {
			fields[idx] = f
		}
	}
	return
}
