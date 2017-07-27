package swagger_test

import (
	"testing"

	"github.com/Soontao/go-mysql-api/swagger"
	"os"
	"github.com/Soontao/go-mysql-api/mysql"
)

var connectionStr = os.Getenv("API_CONN_STR")

func TestGenerateSwaggerConfig(t *testing.T) {
	api := mysql.NewMysqlAPI(connectionStr, true)
	defer api.Stop()
	s := swagger.GenSwaggerFromDBMetadata(api.GetDatabaseMetadata())
	j, _ := s.MarshalJSON()
	println(string(j))
}
