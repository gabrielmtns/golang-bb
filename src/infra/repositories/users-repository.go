package repositories

import (
	"database/sql"
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
		`

	_, err = u.DatabaseHandler.Exec(query, user.Name, user.Lastname, user.Username, user.Password, user.Salt)


	if err != nil {
		log.Printf("Error on create user %v", err)
		return user, err
	}

	return user, nil
}

func (u *UserRepository) All()(users []domain.User, err error){
	const query = "SELECT id, name, lastname, username FROM users"

	rows, err := u.DatabaseHandler.Query(query)

	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id int
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


func(u *UserRepository) ById (id int)(domain.User, error) {
	const query = "SELECT id, name, lastname, username WHERE id = ?"
	row, err := u.DatabaseHandler.Query(query, id)

	defer row.Close()


	if err != nil {
		return domain.User{}, nil
	}


	var name sql.NullString
	var lastname sql.NullString
	var username sql.NullString

	row.Next()
	if err = row.Scan(&id, &name, &lastname, &username); err != nil {
		return domain.User{}, err
	}


	return domain.User{
		Id: id,
		Lastname: lastname.String,
		Username: username.String,
	}, nil
}

func(u *UserRepository) Update(id int, user domain.User)(domain.User, error){
	return domain.User{}, nil

}
func(u *UserRepository) Delete(id int) error {
	return nil
}