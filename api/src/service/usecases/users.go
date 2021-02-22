package service

import (
	domain "go-bb-test/src/domain/models"
	repository "go-bb-test/src/service/protocols/users"
	"go-bb-test/src/service/utils"
)

type Users struct {
	UserRepository 	repository.IUsersRepository
}


func (u Users) Add(user domain.User) (domain.User, error){
	user.Salt = utils.RandomString(10)
	user.Password = utils.TextHasher(user.Password + user.Salt)
	return u.UserRepository.Create(user)
}

func (u Users) Get(id int64)(domain.User, error) {
	return u.UserRepository.ById(id)
}

func (u Users) GetAll() ([]domain.User, error) {
	return u.UserRepository.All()
}

func (u Users) 	Update(id int64, user domain.User) (domain.User, error) {
	return u.UserRepository.Update(id, user)
}

func (u Users) Delete(id int64) error {
	return u.UserRepository.Delete(id)
}