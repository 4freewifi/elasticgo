package filters

type rangeFilter struct {
	filterImpl
	fieldBody map[string]interface{}
}

func Range(field string) rangeFilter {
	fieldBody := map[string]interface{}{}
	return rangeFilter{
		filterImpl: createFilter("range", map[string]interface{}{
			field: fieldBody,
		}),
		fieldBody: fieldBody,
	}
}

func (f rangeFilter) Gte(value interface{}) rangeFilter {
	f.fieldBody["gte"] = value
	return f
}

func (f rangeFilter) Gt(value interface{}) rangeFilter {
	f.fieldBody["gt"] = value
	return f
}

func (f rangeFilter) Lte(value interface{}) rangeFilter {
	f.fieldBody["lte"] = value
	return f
}

func (f rangeFilter) Lt(value interface{}) rangeFilter {
	f.fieldBody["lt"] = value
	return f
}

// TODO: execution
