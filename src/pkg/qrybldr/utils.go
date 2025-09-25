package qrybldr

import (
	"reflect"
	"strings"
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
