package swagger

import (
	"github.com/go-openapi/spec"
	"fmt"
	"github.com/Soontao/go-mysql-api/mysql"
)

func NewRefSchema(refDefinationName, reftype string) (s spec.Schema) {
	s = spec.Schema{
		spec.VendorExtensible{},
		spec.SchemaProps{
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
		spec.SwaggerSchemaProps{},
		nil,
	}
	return
}

func NewField(sName, sType string) (s spec.Schema) {
	s = spec.Schema{
		spec.VendorExtensible{},
		spec.SchemaProps{
			Type:  spec.StringOrArray{sType},
			Title: sName,
		},
		spec.SwaggerSchemaProps{},
		nil,
	}
	return
}

func NewDefinitionMessageWrap(definitionName string) (sWrap *spec.Schema) {
	sWrap = &spec.Schema{
		spec.VendorExtensible{},
		spec.SchemaProps{
			Type: spec.StringOrArray{"object"},
			Properties: map[string]spec.Schema{
				"status":  NewField("status", "integer"),
				"message": NewField("message", "string"),
				"data":    NewRefSchema(definitionName, "array"),
			},
		},
		spec.SwaggerSchemaProps{},
		nil,
	}
	return
}

func NewSwaggerInfo(title, version string) (info *spec.Info) {
	info = &spec.Info{spec.VendorExtensible{}, spec.InfoProps{
		Title:   title,
		Version: version,
	}}
	return
}

func NewAPIPathItemForTable(tName string) (pathItem spec.PathItem) {
	pathItem = spec.PathItem{
		spec.Refable{}, spec.VendorExtensible{}, spec.PathItemProps{
			Get:  NewOperation(tName, fmt.Sprintf("get some %s records", tName), []spec.Parameter{}, NewDefinitionMessageWrap(tName).SchemaProps),
			Post: NewOperation(tName, fmt.Sprintf("create a %s record", tName), []spec.Parameter{NewParamForDefinition(tName)}, NewDefinitionMessageWrap(tName).SchemaProps),
		},
	}
	return
}

func GetParametersFromDbMetadata(meta *mysql.DataBaseMetadata) (params map[string]spec.Parameter) {
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
