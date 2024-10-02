package user

import (
	"backend-mental-guardians/controllers/user/request"
	"backend-mental-guardians/controllers/user/response"
	userEntities "backend-mental-guardians/entities/user"
	"backend-mental-guardians/utilities"
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

func (userController *UserController) GetProfileByID(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	userId, err := utilities.GetUserIdFromToken(token)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse, err := userController.userUseCase.GetProfileByID(uint(userId))
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	res := response.UserResponse{
		ID:        userResponse.ID,
		Email:     userResponse.Email,
		FirstName: userResponse.FirstName,
		LastName:  userResponse.LastName,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get User By ID", res))
}