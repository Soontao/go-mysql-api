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
	apiBatchPath := fmt.Sprintf("/api/batch/%s", tName)

	withoutIDPathItem := spec.PathItem{}
	withIDPathItem := spec.PathItem{}
	withoutIDBatchPathItem := spec.PathItem{}
	// /api/:table group
	withoutIDPathItem.Get = NewOperation(
		tName,
		fmt.Sprintf("get some %s records", tName),
		NewQueryParametersForMySQLAPI(),
		NewDefinitionMessageWrap(tName, NewRefSchema(tName, "array")).SchemaProps,
	)

	if !isView {
		// /api/:table group
		withoutIDPathItem.Put = NewOperation(
			tName,
			fmt.Sprintf("create a %s record", tName),
			[]spec.Parameter{NewParamForDefinition(tName)},
			NewDefinitionMessageWrap(tName, NewCUDOperationReturnMessage()).SchemaProps,
		)
		withoutIDPathItem.Delete = NewOperation(
			tName,
			fmt.Sprintf("delete some %s records", tName),
			[]spec.Parameter{NewParamForDefinition(tName)},
			NewDefinitionMessageWrap(tName, NewCUDOperationReturnMessage()).SchemaProps,
		)
		// /api/:table/:id group
		withIDPathItem.Get = NewOperation(
			tName,
			fmt.Sprintf("get specific %s record", tName),
			append(NewQueryParametersForMySQLAPI(), NewPathIDParameter(meta)),
			NewDefinitionMessageWrap(tName, NewRefSchema(tName, "array")).SchemaProps,
		)
		withIDPathItem.Post = NewOperation(
			tName,
			fmt.Sprintf("update specific %s record", tName),
			append([]spec.Parameter{NewParamForDefinition(tName)}, NewPathIDParameter(meta)),
			NewDefinitionMessageWrap(tName, NewCUDOperationReturnMessage()).SchemaProps,
		)
		withIDPathItem.Delete = NewOperation(
			tName,
			fmt.Sprintf("delete specific %s record", tName),
			append([]spec.Parameter{}, NewPathIDParameter(meta)),
			NewDefinitionMessageWrap(tName, NewCUDOperationReturnMessage()).SchemaProps,
		)
		withoutIDBatchPathItem.Put = NewOperation(
			tName,
			fmt.Sprintf("Batch create %s records", tName),
			[]spec.Parameter{NewParamForArrayDefinition(tName)},
			NewDefinitionMessageWrap(tName, NewCUDOperationReturnArrayMessage()).SchemaProps,
		)
	}

	paths[apiNoIDPath] = withoutIDPathItem
	if !isView {
		paths[apiIDPath] = withIDPathItem
		paths[apiBatchPath] = withoutIDBatchPathItem
	}
	return
}
