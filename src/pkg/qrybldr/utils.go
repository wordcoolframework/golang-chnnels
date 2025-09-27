package qrybldr

import (
	"reflect"
	"strings"

	"gorm.io/gorm"
)

func (q *Qrybldr) tableName() string {
	if q.model == nil {
		return ""
	}

	method := reflect.ValueOf(q.model).MethodByName("TableName")
	if method.IsValid() {
		results := method.Call(nil)
		if len(results) == 1 {
			if name, ok := results[0].Interface().(string); ok {
				return name
			}
		}
	}

	modelType := reflect.TypeOf(q.model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	return strings.ToLower(modelType.Name()) + "s"
}

func (q *Qrybldr) Clone() *Qrybldr {
	return &Qrybldr{
		db:    q.db.Session(&gorm.Session{}),
		model: q.model,
	}
}

func (q *Qrybldr) Debug() *Qrybldr {
	q.db = q.db.Debug()
	return q
}

func (q *Qrybldr) WithTrashed() *Qrybldr {
	q.db = q.db.Unscoped()
	return q
}

func (q *Qrybldr) OnlyTrashed() *Qrybldr {
	q.db = q.db.Unscoped().Where("deleted_at IS NOT NULL")
	return q
}
