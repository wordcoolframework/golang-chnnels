package qrybldr

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func (q *Qrybldr) Select(columns ...string) *Qrybldr {
	q.db = q.db.Select(columns)
	return q
}

func (q *Qrybldr) Where(query string, args ...interface{}) *Qrybldr {
	q.db = q.db.Where(query, args...)
	return q
}

func (q *Qrybldr) WhereIn(column string, values []interface{}) *Qrybldr {
	if len(values) == 0 {
		return q.Where("1 = 0")
	}
	placeholders := strings.Repeat("?,", len(values))
	placeholders = placeholders[:len(placeholders)-1]
	return q.Where(fmt.Sprintf("%s IN (%s)", column, placeholders), values...)
}

func (q *Qrybldr) WhereNotIn(column string, values []interface{}) *Qrybldr {
	if len(values) == 0 {
		return q
	}
	placeholders := strings.Repeat("?,", len(values))
	placeholders = placeholders[:len(placeholders)-1]
	return q.Where(fmt.Sprintf("%s NOT IN (%s)", column, placeholders), values...)
}

func (q *Qrybldr) WhereBetween(column string, min, max interface{}) *Qrybldr {
	return q.Where(fmt.Sprintf("%s BETWEEN ? AND ?", column), min, max)
}

func (q *Qrybldr) WhereNotBetween(column string, min, max interface{}) *Qrybldr {
	return q.Where(fmt.Sprintf("%s NOT BETWEEN ? AND ?", column), min, max)
}

func (q *Qrybldr) WhereNull(column string) *Qrybldr {
	return q.Where(fmt.Sprintf("%s IS NULL", column))
}

func (q *Qrybldr) WhereNotNull(column string) *Qrybldr {
	return q.Where(fmt.Sprintf("%s IS NOT NULL", column))
}

func (q *Qrybldr) WhereDate(column string, date string) *Qrybldr {
	return q.Where(fmt.Sprintf("DATE(%s) = ?", column), date)
}

func (q *Qrybldr) WhereMonth(column string, month int) *Qrybldr {
	return q.Where(fmt.Sprintf("MONTH(%s) = ?", column), month)
}

func (q *Qrybldr) WhereYear(column string, year int) *Qrybldr {
	return q.Where(fmt.Sprintf("YEAR(%s) = ?", column), year)
}

func (q *Qrybldr) WhereLike(col string, value any) *Qrybldr {
	return q.WhereLikeTable(col, value, q.tableName())
}

func (q *Qrybldr) WhereLikeTable(col string, value any, table interface{}) *Qrybldr {
	var tableName string

	switch t := table.(type) {
	case string:
		tableName = t
	default:
		modelType := reflect.TypeOf(table)
		if modelType.Kind() == reflect.Ptr {
			modelType = modelType.Elem()
		}
		tableName = TableNameFromType(modelType)
	}
	likeValue := fmt.Sprintf("%%%v%%", value)
	q.db = q.db.Where(fmt.Sprintf("%s.%s LIKE ?", tableName, col), likeValue)
	return q
}

func (q *Qrybldr) OrderBy(order string) *Qrybldr {
	q.db = q.db.Order(order)
	return q
}

func (q *Qrybldr) Limit(limit int) *Qrybldr {
	q.db = q.db.Limit(limit)
	return q
}

func (q *Qrybldr) Offset(offset int) *Qrybldr {
	q.db = q.db.Offset(offset)
	return q
}

func (q *Qrybldr) GroupBy(columns ...string) *Qrybldr {
	q.db = q.db.Group(columns[0])
	for _, col := range columns[1:] {
		q.db = q.db.Group(col)
	}
	return q
}

func (q *Qrybldr) Having(cond string, args ...interface{}) *Qrybldr {
	q.db = q.db.Having(cond, args...)
	return q
}

func (q *Qrybldr) Distinct() *Qrybldr {
	q.db = q.db.Distinct()
	return q
}

func (q *Qrybldr) FindByID(dest interface{}, id interface{}) error {
	return q.db.First(dest, id).Error
}

func (q *Qrybldr) Get(dest interface{}) error {
	return q.db.Find(dest).Error
}

func (q *Qrybldr) Pluck(dest interface{}, column string) error {
	return q.db.Pluck(column, dest).Error
}

func (q *Qrybldr) Count(model interface{}) (int64, error) {
	if err := q.checkInitialized(); err != nil {
		return 0, err
	}

	var count int64
	err := q.db.Model(model).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (q *Qrybldr) Exists(model interface{}) (bool, error) {
	if err := q.checkInitialized(); err != nil {
		return false, err
	}

	if model == nil {
		return false, errors.New("model cannot be nil")
	}

	count, err := q.Count(model)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
