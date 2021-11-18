package handler

import (
	"Aws-pgx/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Handler) TestStudent(c echo.Context) error {
	return c.String(http.StatusOK, "Test")
}

func (s *Handler) GetAllStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, "students")
}

func (s *Handler) CreateStudent(c echo.Context) error {
	student := &model.Student{}
	if err := c.Bind(student); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, student)
}
