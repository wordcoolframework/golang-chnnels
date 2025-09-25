package qrybldr

func (q *Qrybldr) Begin() *Qrybldr {
	q.db = q.db.Begin()
	return q
}

func (q *Qrybldr) Commit() *Qrybldr {
	q.db.Commit()
	return q
}

func (q *Qrybldr) Rollback() *Qrybldr {
	q.db.Rollback()
	return q
}
