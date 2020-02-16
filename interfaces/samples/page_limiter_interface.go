type PageLimiter interface {
	GetLimit() int32
	GetPage() int32
}

type Applier interface {
	Apply(squirrel.SelectBuilder) squirrel.SelectBuilder
}

type PageLimiterApplier interface {
	PageLimiter
	Applier
}

