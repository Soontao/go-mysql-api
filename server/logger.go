package server

import (
	"os"

	"github.com/labstack/gommon/log"
)

var L *log.Logger = NewLogger()

func NewLogger() (l *log.Logger) {
	l = log.New("mysql-api-server")
	l.SetHeader(`[${level}] ${time_rfc3339_nano}`)
	l.SetLevel(log.DEBUG)
	l.SetOutput(os.Stdout)
	return
}
