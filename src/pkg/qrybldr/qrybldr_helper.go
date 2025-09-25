package qrybldr

import (
	"reflect"
	"strings"
	"unicode"
)

func TableNameFromType(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	method, ok := reflect.New(t).Interface().(interface {
		TableName() string
	})
	if ok {
		return method.TableName()
	}

	return strings.ToLower(t.Name()) + "s"
}

func Singularize(word string) string {
	if len(word) == 0 {
		return word
	}

	lower := strings.ToLower(word)

	if strings.HasSuffix(lower, "ies") && len(lower) > 3 {
		return lower[:len(lower)-3] + "y" // categories -> category
	}

	if strings.HasSuffix(lower, "s") && len(lower) > 1 {
		return lower[:len(lower)-1] // products -> product
	}

	return lower
}

func toSnakeCase(str string) string {
	if strings.HasSuffix(str, "ID") {
		str = str[:len(str)-2] + "Id"
	}

	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}

	res := string(result)
	res = strings.ReplaceAll(res, "_i_d", "_id")
	return res
}
