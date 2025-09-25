package qrybldr

import (
	"gorm.io/gorm"
)

type Qrybldr struct {
	db    *gorm.DB
	model interface{}
}

func Instance(db *gorm.DB) *Qrybldr {
	return &Qrybldr{db: db}
}

func (q *Qrybldr) Model(model interface{}) *Qrybldr {
	q.model = model
	q.db = q.db.Model(model)
	return q
}

func (q *Qrybldr) Table(table string) *Qrybldr {
	q.db = q.db.Table(table)
	return q
}

func (q *Qrybldr) Gorm() *gorm.DB {
	return q.db
}
