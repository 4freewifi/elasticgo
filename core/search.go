package core

import (
	. "github.com/4FreeWifi/elasticgo/aggregations"
	. "github.com/4FreeWifi/elasticgo/queries"
	"encoding/json"
)

type Search interface {
	Size(int) Search
	Query(Query) Search
	Aggregate(Aggregation) Search
}

type searchImpl struct {
	dsl map[string]interface{}
}

func NewSearch() Search {
	return searchImpl{
		dsl: map[string]interface{}{},
	}
}

func (s searchImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.dsl)
}

func (s searchImpl) Size(size int) Search {
	s.dsl["size"] = size
	return s
}

func (s searchImpl) Query(q Query) Search {
	s.dsl["query"] = q
	return s
}

func (s searchImpl) Aggregate(a Aggregation) Search {
	s.dsl["aggregations"] = a
	return s
}
