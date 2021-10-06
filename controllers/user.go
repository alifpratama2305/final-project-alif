package controllers

import (
	"final-project/config"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// GetUser godoc
// @Summary Get details of all user
// @Description Get details of all user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func GetUser(ctx echo.Context) error {
	db, err := config.SetupDB()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	rows, err := db.Query("SELECT user_id, name FROM user")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	var result []models.User

	for rows.Next() {
		each := models.User{}
		err := rows.Scan(&each.UserID, &each.Name)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, "Internal server error")
		}

		result = append(result, each)
	}

	if len(result) == 0 {
		return ctx.JSON(http.StatusNotFound, "Data not found")
	}

	return ctx.JSON(http.StatusOK, result)
}

// GetUserByID godoc
// @Summary Get details of user by id
// @Description Get details of user by id
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /users/{userId} [get]
func GetUserByID(ctx echo.Context) error {
	db, err := config.SetupDB()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	id := ctx.Param("id")

	result := models.User{}

	err = db.QueryRow("SELECT user_id, name FROM user WHERE user_id = ?", id).
		Scan(&result.UserID, &result.Name)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Data not found")
	}

	return ctx.JSON(http.StatusOK, result)

}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags user
// @Accept json
// @Produce json
// @Param user body User true "Create user"
// @Success 200 {object} User
// @Router /users [post]
func CreateUser(ctx echo.Context) error {
	db, err := config.SetupDB()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	name := ctx.FormValue("name")

	rows, err := db.Exec("INSERT INTO user (name) VALUES (?)", name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	userID, err := rows.LastInsertId()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	result := models.User{
		UserID: int8(userID),
		Name:   name,
	}

	return ctx.JSON(http.StatusOK, result)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user with the input payload
// @Tags user
// @Accept json
// @Produce json
// @Param user body User true "Update user"
// @Success 200 {object} User
// @Router /users/{userId} [put]
func UpdateUser(ctx echo.Context) error {
	db, err := config.SetupDB()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	id := ctx.Param("id")

	name := ctx.FormValue("name")

	rows, err := db.Exec("UPDATE user SET name = ? WHERE user_id = ?", name, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	if rowsAffected == 0 {
		return ctx.JSON(http.StatusNotFound, "Data not found")
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}
	result := models.User{
		UserID: int8(userID),
		Name:   name,
	}

	return ctx.JSON(http.StatusOK, result)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body User true "Delete user"
// @Success 200 {object} User
// @Router /users/{userId} [delete]
func DeleteUser(ctx echo.Context) error {
	db, err := config.SetupDB()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	id := ctx.Param("id")

	rows, err := db.Exec("DELETE FROM user WHERE user_id = ?", id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	if rowsAffected == 0 {
		return ctx.JSON(http.StatusNotFound, "Data not found")
	}

	return ctx.JSON(http.StatusOK, "Delete data success")
}
