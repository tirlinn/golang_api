package handler

import (
	"Aws-pgx/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) CreateClass(c echo.Context) error {
	var input model.Class

	if err := c.Bind(&input); err != nil {
		return err
	}

	id, err := h.services.CreateClass(input)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func(h *Handler) GetAllClasses(c echo.Context) error {
	classes, err := h.services.GetAllClasses()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, classes)
}
