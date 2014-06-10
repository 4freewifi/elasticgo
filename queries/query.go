package queries

import (
	"encoding/json"
)

type Query interface {
}

type queryImpl struct {
	name string
	dsl  map[string]map[string]interface{}
}

func createQuery(query string, body map[string]interface{}) queryImpl {
	return queryImpl{
		name: query,
		dsl: map[string]map[string]interface{}{
			query: body,
		},
	}
}

func (f queryImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.dsl)
}
