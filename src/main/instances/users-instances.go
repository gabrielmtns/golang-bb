package instances

import (
	"database/sql"
	"go-bb-test/src/app/controllers"
	"go-bb-test/src/infra/database"
	"go-bb-test/src/infra/repositories"
	"go-bb-test/src/service/usecases"
)


func UserControllerFactory(dbRef *sql.DB) *controllers.UserController {

	var databaseHandler = database.DbHandler{
		Connection: dbRef,
	}

	var userRepository = repositories.NewUserRepository(databaseHandler)

	var userService = service.Users{
		UserRepository: userRepository,
	}

	return controllers.NewUserController(userService)
}
