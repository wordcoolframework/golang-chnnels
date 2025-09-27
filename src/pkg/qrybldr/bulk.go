package qrybldr

func (q *Qrybldr) BulkInsert(data []interface{}) error {
	if len(data) == 0 {
		return nil
	}
	return q.db.CreateInBatches(data, 1000).Error
}

func (q *Qrybldr) BulkUpdate(updates map[string]interface{}, ids []uint) error {
	if len(ids) == 0 {
		return nil
	}
	return q.db.Where("id IN (?)", ids).Updates(updates).Error
}

func (q *Qrybldr) BulkDelete(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}
	return q.db.Where("id IN (?)", ids).Delete(q.model).Error
}
