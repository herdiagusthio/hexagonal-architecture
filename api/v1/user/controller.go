package user

import (
	"hexagonalArchitecture/business/user"

	"github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service user.Service
}

//NewController Construct item API controller
func NewController(service user.Service) *Controller {
	return &Controller{
		service,
	}
}

//GetItemByID Get item by ID echo handler
func (controller *Controller) FindUserByID(c echo.Context) error {
	return nil
}
