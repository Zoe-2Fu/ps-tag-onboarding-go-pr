package handler

import (
	"net/http"

	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go-pr/internal/constants"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				errMsg := errs.ErrorMessage{
					Error:   "Status Not Found",
					Details: []string{errs.ResponseUserNotFound},
				}
				return c.JSON(http.StatusNotFound, errMsg)
			}

			if errMsg, ok := c.Get("validationError").(*errs.ErrorMessage); ok {
				return c.JSON(http.StatusBadRequest, errMsg)
			} else if httpErr, ok := err.(*echo.HTTPError); ok {
				return c.JSON(httpErr.Code, httpErr.Message)
			}
		}
		return nil
	}
}
