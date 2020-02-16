func MustApply(a Applier, q *sq.SelectBuilder) error {
	if a == nil {
		return ErrApplierIsNil
	}
	if q == nil {
		return ErrSelectBuilderIsNil
	}

	*q = a.Apply(*q)

	return nil
}

func FromParams(pl PageLimiter) PageLimitApplier {
	page := pl.GetPage()
	if page < 1 {
		page = 1
	}
	limit := pl.GetLimit()
	if limit < 1 {
		limit = 50
	}
	return New(limit, page)
}
