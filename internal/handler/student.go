package handler

import (
	"Aws-pgx/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) TestStudent(c echo.Context) error {
	return c.String(http.StatusOK, "Test")
}

func (h *Handler) GetAllStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, "students")
}

func (h *Handler) CreateStudent(c echo.Context) error {
	var student model.Student

	if err := c.Bind(&student); err != nil {
		return err
	}

	id, err := h.services.CreateStudent(student)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
