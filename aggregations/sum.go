package aggregations

type sumAggregation struct {
	aggregationImpl
}

func Sum(name string, field string) sumAggregation {
	return sumAggregation{
		createAggregation("sum", name, map[string]interface{}{
			"field": field,
		}),
	}
}
