package aggregations

type termsAggregation struct {
	aggregationImpl
}

func Terms(name string, field string) termsAggregation {
	return termsAggregation{
		createAggregation("terms", name, map[string]interface{}{
			"field": field,
		}),
	}
}
