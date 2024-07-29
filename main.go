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
