package qrybldr

import "fmt"

func (q *Qrybldr) Select(columns ...string) *Qrybldr {
	q.db = q.db.Select(columns)
	return q
}

func (q *Qrybldr) Where(query string, args ...interface{}) *Qrybldr {
	q.db = q.db.Where(query, args...)
	return q
}

func (q *Qrybldr) WhereLike(col string, value any) *Qrybldr {
	likeValue := fmt.Sprintf("%%%v%%", value)
	q.db = q.db.Where(fmt.Sprintf("%s LIKE ?", col), likeValue)
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

func (q *Qrybldr) Join(table string, condition string) *Qrybldr {
	q.db = q.db.Joins("JOIN " + table + " ON " + condition)
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

func (q *Qrybldr) Count() (int64, error) {
	var count int64
	err := q.db.Count(&count).Error
	return count, err
}

func (q *Qrybldr) Exists() (bool, error) {
	c, err := q.Count()
	return c > 0, err
}
