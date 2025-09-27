package qrybldr

import "fmt"

func (q *Qrybldr) Join(table string, condition string) *Qrybldr {
	q.db = q.db.Joins("JOIN " + table + " ON " + condition)
	return q
}

func (q *Qrybldr) LeftJoin(table, condition string) *Qrybldr {
	q.db = q.db.Joins("LEFT JOIN " + table + " ON " + condition)
	return q
}

func (q *Qrybldr) RightJoin(table, condition string) *Qrybldr {
	q.db = q.db.Joins("RIGHT JOIN " + table + " ON " + condition)
	return q
}

func (q *Qrybldr) InnerJoin(table, condition string) *Qrybldr {
	q.db = q.db.Joins("INNER JOIN " + table + " ON " + condition)
	return q
}

func (q *Qrybldr) JoinWithSelect(joinType, table, condition string, selectColumns []string) *Qrybldr {
	q.db = q.db.Joins(fmt.Sprintf("%s JOIN %s ON %s", joinType, table, condition))
	if len(selectColumns) > 0 {
		q.db = q.db.Select(selectColumns)
	}
	return q
}
