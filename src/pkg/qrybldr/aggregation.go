package qrybldr

import "fmt"

func (q *Qrybldr) Sum(column string) (float64, error) {
	var result float64
	err := q.db.Select(fmt.Sprintf("SUM(%s) as total", column)).Scan(&result).Error
	return result, err
}

func (q *Qrybldr) Avg(column string) (float64, error) {
	var result float64
	err := q.db.Select(fmt.Sprintf("AVG(%s) as average", column)).Scan(&result).Error
	return result, err
}

func (q *Qrybldr) Min(column string) (interface{}, error) {
	var result interface{}
	err := q.db.Select(fmt.Sprintf("MIN(%s) as minimum", column)).Scan(&result).Error
	return result, err
}

func (q *Qrybldr) Max(column string) (interface{}, error) {
	var result interface{}
	err := q.db.Select(fmt.Sprintf("MAX(%s) as maximum", column)).Scan(&result).Error
	return result, err
}

func (q *Qrybldr) GroupByWithHaving(groupColumns []string, havingCond string, havingArgs ...interface{}) *Qrybldr {
	q = q.GroupBy(groupColumns...)
	return q.Having(havingCond, havingArgs...)
}
