package qrybldr

import (
	"fmt"
	"strings"
)

func (q *Qrybldr) Search(searchTerm string, columns ...string) *Qrybldr {
	if searchTerm == "" || len(columns) == 0 {
		return q
	}

	var conditions []string
	var args []interface{}

	for _, col := range columns {
		conditions = append(conditions, fmt.Sprintf("%s LIKE ?", col))
		args = append(args, fmt.Sprintf("%%%s%%", searchTerm))
	}

	whereClause := strings.Join(conditions, " OR ")
	return q.Where(whereClause, args...)
}

func (q *Qrybldr) FullTextSearch(columns []string, searchTerm string) *Qrybldr {
	if len(columns) == 0 || searchTerm == "" {
		return q
	}
	matchClause := fmt.Sprintf("MATCH(%s) AGAINST(? IN BOOLEAN MODE)", strings.Join(columns, ","))
	return q.Where(matchClause, searchTerm)
}
