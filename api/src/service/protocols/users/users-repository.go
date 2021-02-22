package repository

import domain "go-bb-test/src/domain/models"

type IUsersRepository interface {
	All()([]domain.User, error)
	ById(id int64)(domain.User, error)
	Create(user domain.User)(domain.User, error)
	Update(id int64, user domain.User)(domain.User, error)
	Delete(id int64) error
	UsernamePasswordCompare(username string, password string)(domain.User, error)
	SaltByUsername(username string)(string, error)
}
