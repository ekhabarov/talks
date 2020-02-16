func (s *service) List(pla PageLimitApplier) ([]Order, error) {
	log("limit", pla.GetLimit(), "page", pla.GetPage())

	orders, err := s.repo.List(pla)
	// do something with categories
}
