package usecases

import "go-bb-test/src/domain/models"

type IUsers interface {
	Get(id int64)(domain.User, error)
	GetAll() ([]domain.User, error)
	Add(user domain.User) (domain.User, error)
	Delete(id int64) error
	Update(id int64, user domain.User) (domain.User, error)
}
