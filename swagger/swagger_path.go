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
	isView := meta.TableType == "VIEW"
	apiNoIDPath := fmt.Sprintf("/api/%s", tName)
	apiIDPath := fmt.Sprintf("/api/%s/{id}", tName)

	withoutPathItem := spec.PathItem{}
	withIDPathItem := spec.PathItem{}
	// /api/:table group
	withoutPathItem.Get = NewOperation(tName, fmt.Sprintf("get some %s records", tName), []spec.Parameter{}, NewDefinitionMessageWrap(tName, NewRefSchema(tName, "array")).SchemaProps)

	if !isView {
		// /api/:table group
		withoutPathItem.Put = NewOperation(tName, fmt.Sprintf("create a %s record", tName), []spec.Parameter{NewParamForDefinition(tName)}, NewDefinitionMessageWrap(tName, NewCUDOperationReturnMessage()).SchemaProps)
		withoutPathItem.Delete = NewOperation(tName, fmt.Sprintf("delete some %s records", tName), []spec.Parameter{}, NewDefinitionMessageWrap(tName, NewCUDOperationReturnMessage()).SchemaProps)
		// /api/:table/:id group
		withIDPathItem.Get = NewOperation(
			tName,
			fmt.Sprintf("get specific %s record", tName),
			[]spec.Parameter{
				NewPathIDParameter(meta),
			},
			NewDefinitionMessageWrap(tName, NewRefSchema(tName, "array")).SchemaProps,
		)
	}

	paths[apiNoIDPath] = withoutPathItem
	if !isView {
		paths[apiIDPath] = withIDPathItem
	}
	return
}
