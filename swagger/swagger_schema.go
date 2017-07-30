package swagger

import (
	"regexp"
	"strings"

	"github.com/Soontao/go-mysql-api/mysql"
	"github.com/go-openapi/spec"
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

func getEnumIfItIs(c *mysql.ColumnMetadata) (enum []interface{}) {
	enum = make([]interface{}, 0)
	re := regexp.MustCompile("\\'([\\w]+)\\'")
	if strings.HasPrefix(c.ColumnType, "enum") {
		enumStrArr := re.FindAllString(c.ColumnType, -1)
		for _, enumItem := range enumStrArr {
			enum = append(enum, enumItem)
		}
	}
	return
}

func ColumnSchema(col *mysql.ColumnMetadata) (s *spec.Schema) {
	s = &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: spec.StringOrArray{dbTypeToSchemaType(col.DataType)},
		},
	}
	return
}

func SchemaPropsFromTbmeta(tMeta *mysql.TableMetadata) (schemaProp spec.SchemaProps) {
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
				Enum:        getEnumIfItIs(col),
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
