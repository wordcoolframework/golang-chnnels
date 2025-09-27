package qrybldr

import (
	"fmt"
	"reflect"
	"strings"
)

type PaginationResult struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PerPage  int         `json:"per_page"`
	LastPage int         `json:"last_page"`
}

type CursorPaginationResult struct {
	Data    interface{} `json:"data"`
	HasNext bool        `json:"has_next"`
	Cursor  interface{} `json:"cursor,omitempty"`
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

func (q *Qrybldr) PaginateWithSort(dest interface{}, page, perPage int, sortBy, sortOrder string) (PaginationResult, error) {
	if sortBy != "" {
		order := sortBy
		if strings.ToUpper(sortOrder) == "DESC" {
			order += " DESC"
		} else {
			order += " ASC"
		}
		q.db = q.db.Order(order)
	}
	return q.Paginate(dest, page, perPage)
}

func (q *Qrybldr) CursorPaginate(dest interface{}, cursorField string, cursorValue interface{}, perPage int, reverse bool) (CursorPaginationResult, error) {
	var result CursorPaginationResult

	if reverse {
		q.db = q.db.Where(fmt.Sprintf("%s < ?", cursorField), cursorValue)
	} else {
		q.db = q.db.Where(fmt.Sprintf("%s > ?", cursorField), cursorValue)
	}

	q.db = q.db.Order(fmt.Sprintf("%s DESC", cursorField)).Limit(perPage + 1)

	err := q.db.Find(dest).Error
	if err != nil {
		return result, err
	}

	destSlice := reflect.ValueOf(dest)
	if destSlice.Kind() == reflect.Ptr {
		destSlice = destSlice.Elem()
	}

	if destSlice.Len() > perPage {
		result.HasNext = true
		result.Data = destSlice.Slice(0, perPage).Interface()
	} else {
		result.HasNext = false
		result.Data = dest
	}

	return result, nil
}
