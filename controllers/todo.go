package controllers

import (
	"final-project/config"
	"final-project/models"
	"net/http"

	"github.com/labstack/echo"
)

// GetTodo godoc
// @Summary Get details of all todo
// @Description Get details of all todo
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {array} Todo
// @Router /todos [get]
func GetTodo(ctx echo.Context) error {
	db, err := config.SetupDB()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	rows, err := db.Query("SELECT title, description, due_date, user.name, status.status_txt FROM todo JOIN user ON person_in_charge = user.user_id JOIN status ON status = status.status_id")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	var result []models.Todo

	for rows.Next() {
		each := models.Todo{}
		err := rows.Scan(&each.Title, &each.Description, &each.DueDate, &each.PersonInCharge, &each.Status)

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

// GetTodoByID godoc
// @Summary Get details of todo by id
// @Description Get details of todo by id
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {object} Todo
// @Router /todos/{id} [get]
func GetTodoByID(ctx echo.Context) error {
	db, err := config.SetupDB()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	id := ctx.Param("id")

	result := models.Todo{}

	err = db.QueryRow("SELECT title, description, due_date, user.name, status.status_txt FROM todo JOIN user ON person_in_charge = user.user_id JOIN status ON status = status.status_id WHERE id = ?", id).
		Scan(&result.Title, &result.Description, &result.DueDate, &result.PersonInCharge, &result.Status)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Data not found")
	}

	return ctx.JSON(http.StatusOK, result)
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo with the input payload
// @Tags todo
// @Accept json
// @Produce json
// @Param todo body Todo true "Create todo"
// @Success 200 {object} Todo
// @Router /todos [post]
func CreateTodo(ctx echo.Context) error {
	db, err := config.SetupDB()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	title := ctx.FormValue("title")
	description := ctx.FormValue("description")
	due_date := ctx.FormValue("due_date")
	person_in_charge := ctx.FormValue("person_in_charge")
	status := ctx.FormValue("status")

	dataUser := models.User{}
	err = db.QueryRow("SELECT user_id, name FROM user WHERE name = ?", person_in_charge).
		Scan(&dataUser.UserID, &dataUser.Name)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "User not found")
	}

	dataStatus := models.Status{}
	err = db.QueryRow("SELECT status_id, status_txt FROM status WHERE status_txt = ?", status).
		Scan(&dataStatus.StatusID, &dataStatus.StatusTxt)
	if err != nil {
		dataStatus.StatusID = 1
		status = "New"
	}

	rows, err := db.Exec("INSERT INTO todo (title, description, due_date, person_in_charge, status) VALUES (?, ?, ?, ?, ?)", title, description, due_date, dataUser.UserID, dataStatus.StatusID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	_, err = rows.LastInsertId()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	result := models.Todo{
		Title:          title,
		Description:    description,
		DueDate:        due_date,
		PersonInCharge: person_in_charge,
		Status:         status,
	}

	return ctx.JSON(http.StatusOK, result)
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update a todo with the input payload
// @Tags todo
// @Accept json
// @Produce json
// @Param todo body Todo true "Update todo"
// @Success 200 {object} Todo
// @Router /todos/{id} [put]
func UpdateTodo(ctx echo.Context) error {
	db, err := config.SetupDB()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	id := ctx.Param("id")

	title := ctx.FormValue("title")
	description := ctx.FormValue("description")
	due_date := ctx.FormValue("due_date")
	person_in_charge := ctx.FormValue("person_in_charge")
	status := ctx.FormValue("status")

	dataTodo := models.Todo{}
	err = db.QueryRow("SELECT status.status_txt FROM todo JOIN status ON status = status.status_id WHERE id = ?", id).
		Scan(&dataTodo.Status)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Data not found")
	}

	if dataTodo.Status == "Done" || dataTodo.Status == "Deleted" {
		return ctx.JSON(http.StatusNotFound, "Status has done or deleted")
	}

	dataUser := models.User{}
	err = db.QueryRow("SELECT user_id, name FROM user WHERE name = ?", person_in_charge).
		Scan(&dataUser.UserID, &dataUser.Name)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "User not found")
	}

	dataStatus := models.Status{}
	err = db.QueryRow("SELECT status_id, status_txt FROM status WHERE status_txt = ?", status).
		Scan(&dataStatus.StatusID, &dataStatus.StatusTxt)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Status not found")
	}

	rows, err := db.Exec("UPDATE todo SET title = ?, description = ?, due_date = ?, person_in_charge = ?, status = ? WHERE id = ?", title, description, due_date, dataUser.UserID, dataStatus.StatusID, id)
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

	result := models.Todo{
		Title:          title,
		Description:    description,
		DueDate:        due_date,
		PersonInCharge: person_in_charge,
		Status:         status,
	}

	return ctx.JSON(http.StatusOK, result)
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo
// @Tags todo
// @Accept json
// @Produce json
// @Param todo body Todo true "Delete todo"
// @Success 200 {object} Todo
// @Router /todos/{id} [delete]
func DeleteTodo(ctx echo.Context) error {
	db, err := config.SetupDB()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	id := ctx.Param("id")

	rows, err := db.Exec("DELETE FROM todo WHERE id = ?", id)
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
