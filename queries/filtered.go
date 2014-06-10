package queries

import (
	. "github.com/4FreeWifi/elasticgo/filters"
)

type filteredQuery struct {
	queryImpl
}

func Filtered(query Query, filter Filter) filteredQuery {
	return filteredQuery{
		createQuery("filtered", map[string]interface{}{
			"query":  query,
			"filter": filter,
		}),
	}
}

// TODO: strategy
