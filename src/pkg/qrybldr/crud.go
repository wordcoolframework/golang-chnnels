package qrybldr

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (q *Qrybldr) First(dest interface{}) error {
	return q.db.First(dest).Error
}

func (q *Qrybldr) FirstOrFail(dest interface{}) error {
	err := q.db.First(dest).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	}
	return err
}

func (q *Qrybldr) Create(dest interface{}) error {
	if q.db == nil {
		return fmt.Errorf("gorm.DB is nil in Qrybldr")
	}
	return q.db.Create(dest).Error
}

func (q *Qrybldr) FirstOrCreate(dest interface{}, defaults map[string]interface{}) error {
	return q.db.FirstOrCreate(dest, defaults).Error
}

func (q *Qrybldr) Update(values map[string]interface{}) error {
	return q.db.Updates(values).Error
}

func (q *Qrybldr) Delete(model interface{}) error {
	return q.db.Delete(model).Error
}
