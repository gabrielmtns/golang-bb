package usecases

import "go-bb-test/src/domain/models"

type IUsers interface {
	Get(id int)(domain.User, error)
	GetAll() ([]domain.User, error)
	Add(user domain.User) (domain.User, error)
	Delete(id int) error
	Update(id int, user domain.User) (domain.User, error)
}
