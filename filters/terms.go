package filters

type termsFilter struct {
	filterImpl
}

func Terms(field string, values ...interface{}) termsFilter {
	return termsFilter{
		createFilter("terms", map[string]interface{}{
			field: values,
		}),
	}
}

func (f termsFilter) Add(field string, values ...interface{}) termsFilter {
	f.dsl["terms"][field] = values
	return f
}

// TODO: execution
