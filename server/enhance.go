package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mediocregopher/gojson"
)

// Message
type Message struct {
	Status  int
	Message string
	Data    interface{}
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

func goJSONPretty(c echo.Context, code int, i interface{}, indent string) (err error) {
	b, err := gojson.MarshalIndent(i, "", indent)
	if err != nil {
		return
	}
	return c.JSONBlob(code, b)
}

func processError(c echo.Context, err error) error {
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Message{Status: 500, Message: err.Error()})
	}
	return nil
}
