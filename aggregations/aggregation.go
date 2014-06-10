package aggregations

import (
	"encoding/json"
)

type Aggregation interface {
	MarshalJSON() ([]byte, error)
	Size(int) Aggregation
	Aggregate(Aggregation) Aggregation
}

type aggregationImpl struct {
	name string
	dsl  map[string]map[string]interface{}
}

func createAggregation(aggs string, name string,
	body map[string]interface{}) aggregationImpl {
	return aggregationImpl{
		name: name,
		dsl: map[string]map[string]interface{}{
			name: map[string]interface{}{
				aggs: body,
			},
		},
	}
}

func (a aggregationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.dsl)
}

func (a aggregationImpl) Aggregate(child Aggregation) Aggregation {
	a.dsl[a.name]["aggregations"] = child
	return a
}

func (a aggregationImpl) Size(size int) Aggregation {
	a.dsl[a.name]["size"] = size
	return a
}
