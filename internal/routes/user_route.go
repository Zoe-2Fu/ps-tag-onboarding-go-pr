package routes

import (
	handler "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/handlers"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo, handler handler.UserHandler) {
	e.GET("/user/:id", handler.Find)
	e.POST("/user", handler.Save)
}
