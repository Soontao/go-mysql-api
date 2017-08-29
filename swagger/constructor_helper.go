package swagger

import (
	"github.com/go-openapi/spec"
	"fmt"
	"github.com/Soontao/go-mysql-api/key"
	. "github.com/Soontao/go-mysql-api/types"
)

func NewRefSchema(refDefinationName, reftype string) (s spec.Schema) {
	s = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: spec.StringOrArray{reftype},
			Items: &spec.SchemaOrArray{
				&spec.Schema{
					spec.VendorExtensible{},
					spec.SchemaProps{
						Ref: getTableSwaggerRef(refDefinationName),
					},
					spec.SwaggerSchemaProps{},
					nil,
				},
				nil,
			},
		},
	}
	return
}

func NewField(sName, sType string, iExample interface{}) (s spec.Schema) {
	s = spec.Schema{
		spec.VendorExtensible{},
		spec.SchemaProps{
			Type:  spec.StringOrArray{sType},
			Title: sName,
		},
		spec.SwaggerSchemaProps{
			Example: iExample,
		},
		nil,
	}
	return
}

func NewCUDOperationReturnMessage() (s spec.Schema) {
	s = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: spec.StringOrArray{"object"},
			Properties: map[string]spec.Schema{
				"lastInsertID":  NewField("lastInsertID", "integer", 0),
				"rowesAffected": NewField("rowesAffected", "integer", 1),
			},
		},
	}
	return
}

func NewCUDOperationReturnArrayMessage() (s spec.Schema) {
	s = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: spec.StringOrArray{"array"},
			Items: &spec.SchemaOrArray{
				Schema: &spec.Schema{
					SchemaProps: spec.SchemaProps{
						Properties: map[string]spec.Schema{
							"lastInsertID":  NewField("lastInsertID", "integer", 0),
							"rowesAffected": NewField("rowesAffected", "integer", 1),
						},
					},
				},
			},
		},
	}
	return
}

func NewDefinitionMessageWrap(definitionName string, data spec.Schema) (sWrap *spec.Schema) {

	sWrap = &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: spec.StringOrArray{"object"},
			Properties: map[string]spec.Schema{
				"status":  NewField("status", "integer", 200),
				"message": NewField("message", "string", nil),
				"data":    data,
			},
		},
		SwaggerSchemaProps: spec.SwaggerSchemaProps{},
	}
	return
}

func NewSwaggerInfo(meta *DataBaseMetadata, version string) (info *spec.Info) {
	info = &spec.Info{spec.VendorExtensible{}, spec.InfoProps{
		Title:       fmt.Sprintf("Database %s API", meta.DatabaseName),
		Version:     version,
		Description: "To the time to life",
	}}
	return
}

func GetParametersFromDbMetadata(meta *DataBaseMetadata) (params map[string]spec.Parameter) {
	params = make(map[string]spec.Parameter)
	for _, t := range meta.Tables {
		for _, col := range t.Columns {
			params[col.ColumnName] = spec.Parameter{
				ParamProps: spec.ParamProps{
					In:          "body",
					Description: col.Comment,
					Name:        col.ColumnName,
					Required:    col.NullAble == "true",
				},
			}
		}
	}
	return
}

func NewQueryParametersForMySQLAPI() (ps []spec.Parameter) {
	ps = []spec.Parameter{
		NewQueryParameter(key.KEY_QUERY_FIELD, "include a field", "string", false),
		NewQueryParameter(key.KEY_QUERY_FIELDS, "include some fields, split with comma", "string", false),
		NewQueryParameter(key.KEY_QUERY_LIMIT, "limit max records num", "integer", false),
		NewQueryParameter(key.KEY_QUERY_SKIP, "The Number will be skiped at table start", "integer", false),
		NewQueryParameter(key.KEY_QUERY_WHERE, "Filter with field name and value", "string", false),
		NewQueryParameter(key.KEY_QUERY_LINK, "Auto join a table", "string", false),
		NewQueryParameter(key.KEY_QUERY_SEARCH, "Full table search with str", "string", false),
	}
	return
}

func NewQueryParameter(paramName, paramDescription, paramType string, required bool) (p spec.Parameter) {
	p = spec.Parameter{
		SimpleSchema: spec.SimpleSchema{
			Type: paramType,
		},
		ParamProps: spec.ParamProps{
			In:          "query",
			Name:        paramName,
			Required:    required,
			Description: paramDescription,
		},
	}
	return
}

func NewPathIDParameter(tMeta *TableMetadata) (p spec.Parameter) {
	p = spec.Parameter{
		SimpleSchema: spec.SimpleSchema{
			Type: "string",
		},
		ParamProps: spec.ParamProps{
			In:          "path",
			Name:        "id",
			Required:    true,
			Description: fmt.Sprintf("%s %s", tMeta.TableName, tMeta.GetPrimaryColumn().ColumnName),
		},
	}
	return
}

func NewParamForArrayDefinition(tName string) (p spec.Parameter) {
	s := NewRefSchema(tName, "array")
	p = spec.Parameter{
		ParamProps: spec.ParamProps{
			In:     "body",
			Name:   tName,
			Schema: &s,
		},
	}
	return
}

func NewParamForDefinition(tName string) (p spec.Parameter) {
	p = spec.Parameter{
		ParamProps: spec.ParamProps{
			In:     "body",
			Name:   tName,
			Schema: getTableSwaggerRefSchema(tName),
		},
	}
	return
}

func NewOperation(tName, opDescribetion string, params []spec.Parameter, respSchemaProps spec.SchemaProps) (op *spec.Operation) {
	op = &spec.Operation{
		spec.VendorExtensible{}, spec.OperationProps{
			Description: opDescribetion,
			Tags:        []string{tName},
			Parameters:  params,
			Responses: &spec.Responses{
				spec.VendorExtensible{},
				spec.ResponsesProps{
					nil,
					map[int]spec.Response{
						200: spec.Response{
							spec.Refable{},
							spec.ResponseProps{
								Description: "success",
								Schema: &spec.Schema{
									spec.VendorExtensible{},
									respSchemaProps,
									spec.SwaggerSchemaProps{},
									nil,
								},
							},
							spec.VendorExtensible{},
						},
					},
				},
			},
		},
	}
	return
}

func NewTag(t string) (tag spec.Tag) {
	tag = spec.Tag{TagProps: spec.TagProps{Name: t}}
	return
}

func NewTagsForOne(t string) (tags []spec.Tag) {
	tags = []spec.Tag{NewTag(t)}
	return
}
