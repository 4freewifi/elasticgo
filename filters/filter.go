package filters

import (
	"encoding/json"
)

type Filter interface {
	Cache(bool) Filter
}

type filterImpl struct {
	name string
	dsl  map[string]map[string]interface{}
}

func createFilter(filter string, body map[string]interface{}) filterImpl {
	return filterImpl{
		name: filter,
		dsl: map[string]map[string]interface{}{
			filter: body,
		},
	}
}

func (f filterImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.dsl)
}

func (f filterImpl) Cache(cache bool) Filter {
	f.dsl[f.name]["_cache"] = cache
	return f
}
