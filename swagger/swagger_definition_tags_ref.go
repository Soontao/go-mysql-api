package swagger

import (
	"github.com/go-openapi/spec"
	"github.com/Soontao/go-mysql-api/mysql"
	"github.com/go-openapi/jsonreference"
	"fmt"
)

func SwaggerDefinationsFromDabaseMetadata(dbMeta *mysql.DataBaseMetadata) (definations spec.Definitions) {
	definations = spec.Definitions{}
	for _, t := range dbMeta.Tables {
		schema := spec.Schema{}
		schema.Type = spec.StringOrArray{"object"}
		schema.Title = t.TableName
		schema.SchemaProps = SchemaPropsFromTbmeta(t)
		definations[t.TableName] = schema
	}
	return
}

func getTableSwaggerRef(t string) (ref spec.Ref) {
	ref = spec.Ref{}
	ref.Ref, _ = jsonreference.New(fmt.Sprintf("#/definitions/%s", t))
	return
}

func getTableSwaggerRefAble(t string) (refable spec.Refable) {
	refable = spec.Refable{getTableSwaggerRef(t)}
	return
}

func GetTagsFromDBMetadata(meta *mysql.DataBaseMetadata) (tags []spec.Tag) {
	tags = make([]spec.Tag, 0)
	for _, t := range meta.Tables {
		tags = append(tags, spec.Tag{TagProps: spec.TagProps{Name: t.TableName, Description: t.Comment}})
	}
	return
}
