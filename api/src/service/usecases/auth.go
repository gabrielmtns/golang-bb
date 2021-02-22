package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	domain "go-bb-test/src/domain/models"
	repository "go-bb-test/src/service/protocols/users"
	"go-bb-test/src/service/utils"
	"os"
	"time"
)
type TokenStandard struct {
	User domain.User
	jwt.StandardClaims
}

type Auth struct {
	UserRepository 	repository.IUsersRepository
}
func (a Auth) Login(username string, password string) (domain.UserAuth, error){
	salt, err := a.UserRepository.SaltByUsername(username)

	if err != nil {
		fmt.Print(err)
		return domain.UserAuth{}, err
	}
	fmt.Print("TEXTO DE saida", password + salt)


	hashPassword := utils.TextHasher(password + salt)

	result, err := a.UserRepository.UsernamePasswordCompare(username, hashPassword)

	fmt.Println("hash ::: ", hashPassword)


	if err != nil {
		return domain.UserAuth{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenStandard{
		User: result,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 5).Unix(),
		},
	})

	tk, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		fmt.Println(os.Getenv("JWT_KEY"), err)
		return domain.UserAuth{}, err
	}

	return domain.UserAuth{
		User: result,
		Token: tk,
	}, nil
}