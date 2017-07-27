package swagger

import (
	"github.com/go-openapi/spec"
	"github.com/Soontao/go-mysql-api/mysql"
	"fmt"
)

func GenSwaggerFromDBMetadata(dbMetadata *mysql.DataBaseMetadata) (s *spec.Swagger) {
	s = &spec.Swagger{}
	s.SwaggerProps = spec.SwaggerProps{}
	s.Swagger = "2.0"
	s.Schemes = []string{"http"}
	s.Tags = GetTagsFromDBMetadata(dbMetadata)
	s.Info = NewSwaggerInfo(fmt.Sprintf("Database %s API", dbMetadata.DatabaseName), "version 1")
	s.Definitions = SwaggerDefinationsFromDabaseMetadata(dbMetadata)
	s.Paths = &spec.Paths{Paths:SwaggerPathsFromDatabaseMetadata(dbMetadata)}
	return
}
