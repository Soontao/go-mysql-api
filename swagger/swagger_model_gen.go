package swagger

import (
	"github.com/go-openapi/spec"
	"github.com/Soontao/go-mysql-api/mysql"
	"github.com/go-openapi/jsonreference"
	"fmt"
)

func GenSwaggerFromDBMetadata(dbMetadata *mysql.DataBaseMetadata) (s *spec.Swagger) {
	s = &spec.Swagger{}
	s.SwaggerProps = spec.SwaggerProps{}
	s.Swagger = "2.0"
	s.Schemes = []string{"http"}
	s.Info = NewSwaggerInfo(fmt.Sprintf("Database %s API", dbMetadata.DatabaseName), "version 1")
	s.Definitions = SwaggerDefinationsFromDabaseMetadata(dbMetadata)

	paths := &spec.Paths{}
	paths.Paths = map[string]spec.PathItem{
		"/api/get": spec.PathItem{
			spec.Refable{}, spec.VendorExtensible{}, spec.PathItemProps{
				Get: &spec.Operation{
					spec.VendorExtensible{}, spec.OperationProps{
						Description: "test api",
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
												spec.SchemaProps{
													Type: spec.StringOrArray{"array"},
													Items: &spec.SchemaOrArray{
														&spec.Schema{
															spec.VendorExtensible{},
															spec.SchemaProps{
																Ref: getTableSwaggerRef("monitor"),
															},
															spec.SwaggerSchemaProps{},
															nil,
														},
														nil,
													},
												},
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
				},
			},
		},
	}
	s.SwaggerProps.Paths = paths
	return
}

func SwaggerDefinationsFromDabaseMetadata(dbMeta *mysql.DataBaseMetadata) (definations spec.Definitions) {
	definations = spec.Definitions{}
	for _, t := range dbMeta.Tables {
		schema := spec.Schema{}
		schema.Type = spec.StringOrArray{"object"}
		schema.Title = t.TableName
		schema.SchemaProps = SwaggerSchemaPropsFromTableMetadata(t)
		definations[t.TableName] = schema
	}
	return
}

func SwaggerPathsFromDatabaseMetadata(meta *mysql.DataBaseMetadata) (paths map[string]spec.PathItem) {
	paths = make(map[string]spec.PathItem)
	for _, t := range meta.Tables {
		AppendPathsFor(t, paths)
	}
	return
}

func AppendPathsFor(meta *mysql.TableMetadata, paths map[string]spec.PathItem) () {
	tName := meta.TableName
	definationRef := getTableSwaggerRef(tName)

	return
}

func NewSwaggerInfo(title, version string) (info *spec.Info) {
	info = &spec.Info{spec.VendorExtensible{}, spec.InfoProps{
		Title:   title,
		Version: version,
	}}
	return
}

func NewAPIPathItem() (pathItem spec.PathItem) {
	pathItem = spec.PathItem{
		spec.Refable{}, spec.VendorExtensible{}, spec.PathItemProps{
			Get: &spec.Operation{
				spec.VendorExtensible{}, spec.OperationProps{
					Description: "test api",
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
											spec.SchemaProps{
												Type: spec.StringOrArray{"array"},
												Items: &spec.SchemaOrArray{
													&spec.Schema{
														spec.VendorExtensible{},
														spec.SchemaProps{
															Ref: getTableSwaggerRef("user"),
														},
														spec.SwaggerSchemaProps{},
														nil,
													},
													nil,
												},
											},
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
			},
		},
	}
	return
}

func NewOperation(opDescribetion string, params []spec.Parameter, respSchemaProps spec.SchemaProps) (op *spec.Operation) {
	op = &spec.Operation{
		spec.VendorExtensible{}, spec.OperationProps{
			Description: opDescribetion,
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

func SwaggerSchemaPropsFromTableMetadata(tMeta *mysql.TableMetadata) (schemaProp spec.SchemaProps) {
	schemaProp = spec.SchemaProps{}
	schemaProp.Properties = map[string]spec.Schema{}
	for _, col := range tMeta.Columns {
		schemaProp.Properties[col.ColumnName] = spec.Schema{
			spec.VendorExtensible{},
			spec.SchemaProps{
				Type:        spec.StringOrArray{dbTypeToSchemaType(col.DataType)},
				Description: col.Comment,
				Title:       col.ColumnName,
			},
			spec.SwaggerSchemaProps{},
			nil,
		}
	}
	return
}

func getTableSwaggerRef(t string) (ref spec.Ref) {
	ref = spec.Ref{}
	ref.Ref, _ = jsonreference.New(fmt.Sprintf("#/definitions/%s", t))
	return
}

func dbTypeToSchemaType(t string) (rt_t string) {
	switch t {
	default:
		rt_t = "string"
	case "int", "integer", "bigint", "tinyint", "smallint", "mediumint":
		rt_t = "integer"
	case "float", "double", "decimal":
		rt_t = "number"
	case "bool", "boolean":
		rt_t = "boolean"
	}
	return
}
