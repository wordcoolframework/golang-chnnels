package pkg

import (
	"reflect"
	"strings"
)

func GetTableNameWithModel(model any) string {
	if model == nil {
		return ""
	}

	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	m, ok := reflect.New(t).Type().MethodByName("TableName")
	if ok && m.Type.NumOut() == 1 && m.Type.Out(0).Kind() == reflect.String {
		v := reflect.New(t)
		out := v.MethodByName("TableName").Call(nil)
		if len(out) == 1 {
			return out[0].String()
		}
	}

	name := t.Name()
	return strings.ToLower(name) + "s"
}
