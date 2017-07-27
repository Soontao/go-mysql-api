package swagger

import (
	"github.com/go-openapi/spec"
	"fmt"
	"github.com/Soontao/go-mysql-api/mysql"
)

func SwaggerPathsFromDatabaseMetadata(meta *mysql.DataBaseMetadata) (paths map[string]spec.PathItem) {
	paths = make(map[string]spec.PathItem)
	for _, t := range meta.Tables {
		AppendPathsFor(t, paths)
	}
	return
}

func AppendPathsFor(meta *mysql.TableMetadata, paths map[string]spec.PathItem) () {
	tName := meta.TableName
	apiPath := fmt.Sprintf("/api/%s", tName)
	paths[apiPath] = NewAPIPathItemForTable(tName)
	return
}
