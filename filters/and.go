package filters

type andFilter struct {
	filterImpl
}

func And(filters ...Filter) andFilter {
	return andFilter{
		createFilter("and", map[string]interface{}{
			"filters": filters,
		}),
	}
}
