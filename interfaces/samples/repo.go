import "github.com/Masterminds/squirrel"

// START OMIT
func (r *repo) List(a Applier) ([]Order, error) {
	q := squirrel.Select("...").From("orders").Where( /*...*/ )

	if err := pagination.MustApply(a, &q); err != nil {
		return nil, err
	}

	// do Select
}

// END OMIT
