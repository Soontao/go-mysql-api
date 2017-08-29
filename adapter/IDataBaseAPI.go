package adapter

import (
	"database/sql"
	. "github.com/Soontao/go-mysql-api/types"
)

type IDatabaseAPI interface {
	Create(table string, obj map[string]interface{}) (rs sql.Result, err error)
	Update(table string, id interface{}, obj map[string]interface{}) (rs sql.Result, err error)
	Delete(table string, id interface{}, obj map[string]interface{}) (rs sql.Result, err error)
	Select(option QueryOption) (rs []map[string]interface{}, err error)
	GetDatabaseMetadata() *DataBaseMetadata
	UpdateAPIMetadata() (api IDatabaseAPI)
}



