package api

import (
	"hexagonalArchitecture/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

//RegisterPath Register all API with routing path
func RegisterPath(e *echo.Echo, userController *user.Controller) {
	if userController == nil {
		panic("Controller parameter cannot be nil")
	}

	//user with Versioning endpoint
	userV1 := e.Group("v1/users")
	userV1.GET("/:id", userController.FindUserByID)

}
