package handler

import (
	"Aws-pgx/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(router *echo.Echo) *echo.Echo {

	api := router.Group("/api")
	{
		class := api.Group("/class")
		{
			class.POST("", h.CreateClass)
			class.GET("", h.GetAllClasses)

			classStudent := class.Group(":id/student")
			{
				classStudent.GET("", h.GetAllClasses)
			}
		}

		student := api.Group("/student")
		{
			student.GET("", h.GetAllStudents)
			student.POST("", h.CreateStudent)
		}
	}

	return router
}

