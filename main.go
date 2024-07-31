package main

import (
	"database/sql"
	"fmt"
	"user-management/handler"
	"user-management/repository"
	"user-management/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Whobay123@"
	dbname   = "user_management"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {

	db := connectDB()
	defer db.Close()

	// User
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	route := gin.Default()
	v1 := route.Group("/v1")
	v1.GET("/users", userHandler.ViewUsers)
	v1.GET("/users/:id", userHandler.ViewUserById)
	v1.POST("/users", userHandler.CreateUser)
	v1.PUT("/users/:id", userHandler.UpdateUser)
	v1.DELETE("/users/:id", userHandler.DeleteUser)

	// Role
	roleRepository := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepository)
	roleHandler := handler.NewRoleHandler(roleService)

	v1.GET("/roles", roleHandler.ViewRoles)
	v1.GET("/roles/:id", roleHandler.ViewRolesById)
	v1.POST("/roles", roleHandler.CreateRole)
	v1.PUT("/roles/:id", roleHandler.UpdateRole)
	v1.DELETE("/roles/:id", roleHandler.DeleteRole)

	// User Role
	userRoleRepository := repository.NewUserRoleRepository(db)
	userRoleService := service.NewUserRoleService(userRoleRepository)
	userRolehandler := handler.NewUserRoleHandler(userRoleService)

	v1.GET("/user-roles", userRolehandler.ViewUserRoles)
	v1.GET("/user-roles/:id", userRolehandler.ViewUserRolesById)
	v1.POST("/user-roles", userRolehandler.CreateUserRoles)
	v1.PUT("/user-roles/:id", userRolehandler.UpdateUserRoles)
	v1.DELETE("/user-roles/:id", userRolehandler.DeleteUserRoles)

	route.Run()

}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
	return db
}
