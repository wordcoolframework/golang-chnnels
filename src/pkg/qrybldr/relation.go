package qrybldr

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

func (q *Qrybldr) With(relation string) *Qrybldr {
	modelType := reflect.TypeOf(q.model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	// ---------- nested relation ----------
	if strings.Contains(relation, ".") {
		parts := strings.Split(relation, ".")
		first := parts[0]
		rest := strings.Join(parts[1:], ".")
		q.db = q.db.Preload(first, func(db *gorm.DB) *gorm.DB {
			return db.Preload(rest)
		})
		return q
	}

	// ---------- single-level relation ----------
	field, ok := modelType.FieldByName(relation)
	if !ok {
		panic(fmt.Sprintf("relation %s not found on model %s", relation, modelType.Name()))
	}

	fieldType := field.Type
	isSlice := false
	if fieldType.Kind() == reflect.Slice {
		isSlice = true
		fieldType = fieldType.Elem()
	}

	parentTable := TableNameFromType(fieldType)
	currentTable := q.tableName()
	tag := field.Tag.Get("gorm")

	// many2many
	if strings.Contains(tag, "many2many") {
		pivotTable := ""
		parts := strings.Split(tag, ":")
		if len(parts) > 1 {
			pivotTable = strings.Split(parts[1], ";")[0]
		}

		currentModel := Singularize(modelType.Name())
		parentModel := Singularize(fieldType.Name())

		joinQuery := fmt.Sprintf(
			"LEFT JOIN %s ON %s.id = %s.%s_id LEFT JOIN %s ON %s.%s_id = %s.id",
			pivotTable,
			currentTable, pivotTable, toSnakeCase(currentModel),
			parentTable,
			pivotTable, toSnakeCase(parentModel), parentTable,
		)

		q.db = q.db.Joins(joinQuery).Preload(relation)
		return q
	}

	// belongsTo / hasOne / hasMany
	foreignKey := ""
	if strings.Contains(tag, "foreignKey") {
		parts := strings.Split(tag, ":")
		if len(parts) > 1 {
			foreignKey = strings.TrimSpace(parts[1])
		}
	} else {
		candidate := Singularize(fieldType.Name()) + "ID"
		if _, found := modelType.FieldByName(candidate); found {
			foreignKey = candidate
		}
	}

	if foreignKey != "" {
		foreignKey = toSnakeCase(foreignKey)
		if isSlice {
			joinQuery := fmt.Sprintf(
				"LEFT JOIN %s ON %s.id = %s.%s",
				parentTable, currentTable, parentTable, foreignKey,
			)
			q.db = q.db.Joins(joinQuery).Preload(relation)
		} else {
			joinQuery := fmt.Sprintf(
				"LEFT JOIN %s ON %s.%s = %s.id",
				parentTable, currentTable, foreignKey, parentTable,
			)
			q.db = q.db.Joins(joinQuery).Preload(relation)
		}
		return q
	}

	// fallback
	q.db = q.db.Preload(relation)
	return q
}
