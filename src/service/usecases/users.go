package service

import (
	"crypto/sha256"
	"encoding/hex"
	domain "go-bb-test/src/domain/models"
	repository "go-bb-test/src/service/protocols/users"
	"go-bb-test/src/service/utils"
)

type Users struct {
	UserRepository 	repository.IUsersRepository
}


func (u Users) Add(user domain.User) (domain.User, error){
	user.Salt = utils.RandomString(10)
	var hash = sha256.New()

	for i:= 0; i<3; i++{
		hash.Write([]byte(user.Password + user.Salt))
		user.Password = hex.EncodeToString(hash.Sum(nil))
	}

	return u.UserRepository.Create(user)
}

func (u Users) Get(id int)(domain.User, error) {
	return u.UserRepository.ById(id)
}

func (u Users) GetAll() ([]domain.User, error) {
	return u.UserRepository.All()
}

func (u Users) 	Update(id int, user domain.User) (domain.User, error) {
	return u.UserRepository.Update(id, user)
}

func (u Users) Delete(id int) error {
	return u.UserRepository.Delete(id)
}