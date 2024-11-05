package controller

import (
	"net/http"
	"sturdy-winner-api/internal/api/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

// @Tags    User
// @Summary Get user list
// @Produce json
// @Success 200 {array} dto.UserResponse "OK"
// @Router  /api/users [get]
func (ctrl UserController) GetUserList(c *fiber.Ctx) error {
	resp, err := ctrl.userService.Users()
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(resp)
}
