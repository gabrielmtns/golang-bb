package repository

import domain "go-bb-test/src/domain/models"

type IUsersRepository interface {
	All()([]domain.User, error)
	ById(id int)(domain.User, error)
	Create(user domain.User)(domain.User, error)
	Update(id int, user domain.User)(domain.User, error)
	Delete(id int) error
}
