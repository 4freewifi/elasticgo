package queries

type matchAllQuery struct {
	queryImpl
}

func MatchAll() matchAllQuery {
	return matchAllQuery{
		createQuery("match_all", map[string]interface{}{}),
	}
}

// TODO: boost
