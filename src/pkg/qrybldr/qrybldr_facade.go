package qrybldr

import "gorm.io/gorm"

type qrybldrFacade struct {
	db *gorm.DB
}

var qrybldr *qrybldrFacade

func Init(db *gorm.DB) {
	qrybldr = &qrybldrFacade{db: db}
}

func Orm() *qrybldrFacade {
	return qrybldr
}

func (q *qrybldrFacade) Query() *Qrybldr {
	return Instance(q.db)
}
