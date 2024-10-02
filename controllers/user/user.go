package user

import (
	"backend-mental-guardians/controllers/user/request"
	"backend-mental-guardians/controllers/user/response"
	userEntities "backend-mental-guardians/entities/user"
	"backend-mental-guardians/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase userEntities.UseCaseInterface
}

func NewUserController(userUseCase userEntities.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (userController *UserController) Register(c echo.Context) error {
	var userFromRequest request.UserRegister

	c.Bind(&userFromRequest)

	userEntities := userEntities.User{
		Email:     userFromRequest.Email,
		Password:  userFromRequest.Password,
		FirstName: userFromRequest.FirstName,
		LastName:  userFromRequest.LastName,
	}

	newUser, err := userController.userUseCase.Register(userEntities)

	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.UserToken{
		Token: newUser.Token,
	}
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Register", userResponse))
}

func (userController *UserController) Login(c echo.Context) error {
	var userFromRequest request.UserLogin

	c.Bind(&userFromRequest)

	userEntities := userEntities.User{
		Email:    userFromRequest.Email,
		Password: userFromRequest.Password,
	}

	userResponse, err := userController.userUseCase.Login(userEntities)

	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	response := response.UserToken{
		Token: userResponse.Token,
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", response))
}