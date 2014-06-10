package aggregations

type dateHistogramAggregation struct {
	aggregationImpl
	body map[string]interface{}
}

func DateHistogram(name string, field string,
	interval string) dateHistogramAggregation {

	body := map[string]interface{}{
		"field":    field,
		"interval": interval,
	}
	return dateHistogramAggregation{
		createAggregation("date_histogram", name, body),
		body,
	}
}

func (a dateHistogramAggregation) PreZone(
	zone string) dateHistogramAggregation {
	a.body["pre_zone"] = zone
	return a
}

func (a dateHistogramAggregation) PreZoneAdjustLargeInterval(
	enable bool) dateHistogramAggregation {
	a.body["pre_zone_adjust_large_interval"] = enable
	return a
}

//TODO: timezone
