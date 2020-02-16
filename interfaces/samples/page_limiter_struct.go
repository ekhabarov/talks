type pageLimiter struct {
	limit, page int32
}

func New(limit, page int32) PageLimitApplier {
	return &pageLimiter{limit: limit, page: page}
}

func (pl *pageLimiter) GetLimit() int32 {
	return pl.limit
}

func (pl *pageLimiter) GetPage() int32 {
	return pl.page
}

func (pl *pageLimiter) Apply(q sq.SelectBuilder) sq.SelectBuilder {
	if pl.GetLimit() < 1 {
		return q
	}
	return q.
		Limit(uint64(pl.GetLimit())).
		Offset((uint64(pl.GetPage()) - 1) * uint64(pl.GetLimit()))
}
