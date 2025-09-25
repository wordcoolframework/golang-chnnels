package qrybldr

type PaginationResult struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PerPage  int         `json:"per_page"`
	LastPage int         `json:"last_page"`
}

func (q *Qrybldr) Paginate(dest interface{}, page int, perPage int) (PaginationResult, error) {
	var result PaginationResult
	var count int64

	q.db.Count(&count)
	result.Total = count
	result.Page = page
	result.PerPage = perPage
	if perPage > 0 {
		result.LastPage = int((count + int64(perPage) - 1) / int64(perPage))
	} else {
		result.LastPage = 1
	}

	offset := (page - 1) * perPage
	err := q.db.Offset(offset).Limit(perPage).Find(dest).Error
	result.Data = dest

	return result, err
}
