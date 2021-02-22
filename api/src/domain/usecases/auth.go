package usecases

import domain "go-bb-test/src/domain/models"

type IAuth interface {
	Login(username string, password string)(domain.UserAuth, error)
}