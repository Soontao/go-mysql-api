package swagger

import (
	"github.com/go-openapi/spec"
	types    "github.com/Soontao/go-mysql-api/types"
)

func GenSwaggerFromDBMetadata(dbMetadata *types.DataBaseMetadata) (s *spec.Swagger) {
	s = &spec.Swagger{}
	s.SwaggerProps = spec.SwaggerProps{}
	s.Swagger = "2.0"
	s.Schemes = []string{"http"}
	s.Tags = GetTagsFromDBMetadata(dbMetadata)
	s.Info = NewSwaggerInfo(dbMetadata, "version 1")
	s.Definitions = SwaggerDefinationsFromDabaseMetadata(dbMetadata)
	s.Paths = &spec.Paths{Paths: SwaggerPathsFromDatabaseMetadata(dbMetadata)}
	return
}
