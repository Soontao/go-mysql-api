package swagger

import (
	"fmt"
	"testing"

	"os"

	"github.com/Soontao/go-mysql-api/adapter/mysql"
	"github.com/Soontao/go-mysql-api/server/swaggerswagger"
)

var connectionStr = os.Getenv("API_CONN_STR")

func TestGenerateSwaggerConfig(t *testing.T) {
	api := mysql.NewMysqlAPI(connectionStr, true)
	defer api.Stop()
	s := swagger.GenSwaggerFromDBMetadata(api.GetDatabaseMetadata())
	j, _ := s.MarshalJSON()
	fmt.Println(string(j))
}
