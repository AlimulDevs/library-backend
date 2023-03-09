package helper

import (
	middlewareInstructor "golang/app/middlewares/instructor"
	"golang/models/dto"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) dto.Instructor {
	claims := middlewareInstructor.GetUserInstructor(c)
	var userData dto.Instructor
	if claims != nil {
		userData = dto.Instructor{
			ID: claims.ID,
		}
	}

	return userData
}
