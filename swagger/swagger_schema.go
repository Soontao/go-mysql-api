package swagger

import (
	"github.com/go-openapi/spec"
	"github.com/Soontao/go-mysql-api/mysql"
)

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

func SwaggerSchemaPropsFromTableMetadata(tMeta *mysql.TableMetadata) (schemaProp spec.SchemaProps) {
	schemaProp = spec.SchemaProps{}
	schemaProp.Required = []string{}
	schemaProp.Properties = map[string]spec.Schema{}
	for _, col := range tMeta.Columns {
		if col.NullAble == "NO" {
			schemaProp.Required = append(schemaProp.Required, col.ColumnName)
		}
		schemaProp.Properties[col.ColumnName] = spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        spec.StringOrArray{dbTypeToSchemaType(col.DataType)},
				Description: col.Comment,
				Title:       col.ColumnName,
				Default:     col.DefaultValue,
			},
		}
	}
	return
}

func getTableSwaggerRefSchema(t string) (s *spec.Schema) {
	s = &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Ref: getTableSwaggerRef(t),
		},
	}
	return
}
