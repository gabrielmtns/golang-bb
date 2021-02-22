package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	domain "go-bb-test/src/domain/models"
	"go-bb-test/src/infra/database"
	repository "go-bb-test/src/service/protocols/users"
	"log"
)

type UserRepository struct {
	DatabaseHandler database.DbHandler
}

func NewUserRepository(handler database.DbHandler) repository.IUsersRepository {
	return &UserRepository{
		DatabaseHandler: handler,
	}
}

func (u *UserRepository) Create(user domain.User)(users domain.User, err error){

	const query = `
		INSERT INTO
			users(name, lastname, username, password, salt)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING id
		`

	result, err := u.DatabaseHandler.QueryRow(query, user.Name, user.Lastname, user.Username, user.Password, user.Salt)


	if err != nil {
		log.Printf("Error on create user %v", err)
		return user, err
	}

	result.Scan(&user.Id)

	return user, nil
}


func (u *UserRepository) All()(users []domain.User, err error){
	const query = "SELECT id, name, lastname, username FROM users ORDER BY 2"

	rows, err := u.DatabaseHandler.Query(query)

	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var name sql.NullString
		var lastname sql.NullString
		var username sql.NullString
		if err = rows.Scan(&id, &name, &lastname, &username); err != nil {
			log.Println(err)
			return
		}
		var user = domain.User{
			Id: id,
			Name: name.String,
			Lastname: lastname.String,
			Username: username.String,
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil{
		return
	}

	return
}

func (u *UserRepository) UsernamePasswordCompare (username string, password string)(domain.User, error){
	const query = "SELECT  id, name, lastname, username FROM users WHERE username = $1 and password = $2"

	row, err := u.DatabaseHandler.Query(query, username, password)

	defer row.Close()



	if err != nil {
		return domain.User{}, err
	}

	var name sql.NullString
	var lastname sql.NullString
	var id int64

	row.Next()
	if err = row.Scan(&id, &name, &lastname, &username); err != nil {
		fmt.Print(err)
		return domain.User{}, err
	}

	return domain.User{
		Id: id,
		Name: name.String,
		Lastname: lastname.String,
		Username: username,
	}, nil

}


func (u *UserRepository) SaltByUsername (username string)(string, error){
	const query = "SELECT salt FROM users WHERE username = $1"

	row, err := u.DatabaseHandler.Query(query, username)

	defer row.Close()

	var salt string


	if err != nil {
		return salt, err
	}

	result := row.Next()

	if !result {
		return salt, errors.New("Login ou senha invalidos")
	}

	if err = row.Scan(&salt); err != nil{
		return salt, err
	}

	return salt, nil

}

func(u *UserRepository) ById (id int64)(domain.User, error) {
	const query = "SELECT id, name, lastname, username FROM users WHERE id = $1"
	row, err := u.DatabaseHandler.Query(query, id)

	defer row.Close()


	if err != nil {

		return domain.User{}, err
	}


	var name sql.NullString
	var lastname sql.NullString
	var username sql.NullString

	row.Next()
	if err = row.Scan(&id, &name, &lastname, &username); err != nil {
		fmt.Print(err)
		return domain.User{}, err
	}

	return domain.User{
		Id: id,
		Name: name.String,
		Lastname: lastname.String,
		Username: username.String,
	}, nil
}
func(u *UserRepository) Update(id int64, user domain.User)(domain.User, error){
	const query = "UPDATE users SET name = $1, lastName = $2, username = $3  WHERE id = $4"

	_, err := u.DatabaseHandler.Exec(query, user.Name, user.Lastname, user.Username, id)

	fmt.Print(err)
	if err != nil {
		return user, err
	}

	return user, nil

}
func(u *UserRepository) Delete(id int64) error {
	const query = "DELETE FROM users WHERE id = $1"
	_, err := u.DatabaseHandler.Exec(query, id)
	return err
}