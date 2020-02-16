package service

type User interface {
	Get(id int) model.User
}
