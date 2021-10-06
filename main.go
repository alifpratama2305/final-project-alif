package main

import (
	"final-project/controllers"

	"github.com/labstack/echo"
)

// @title Todos API
// @version 1.0
// @description This is a sample service for managing todos
// @termsOfService http://swagger.io/terms/
// @contact.name Dwi Ahmad Hisyam
// @contact.email dwihisyam3@gmail.com
// @license.name Todos v1.0
// @license.url http://localhost:8080/todos
// @host localhost:8080
// @BasePath /
func main() {
	r := echo.New()

	r.GET("/todos", controllers.GetTodo)
	r.POST("/todos", controllers.CreateTodo)
	r.GET("/todos/:id", controllers.GetTodoByID)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	r.GET("/users", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUserByID)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Start(":8080")
}
