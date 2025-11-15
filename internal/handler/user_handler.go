package handler

import (
	"net/http"
	"runsystem-test/internal/dto"
	"runsystem-test/internal/helper"
	"runsystem-test/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	payLoad := dto.UserRequest{}
	if err := ctx.ShouldBindJSON(&payLoad); err != nil {
		errBind := helper.NewUnprocessableEntityError(err.Error())
		response := helper.ApiResponse(errBind.Status(), errBind.Message(), "error", nil)
		ctx.JSON(errBind.Status(), response)
		return
	}

	usr, err := uh.UserService.CreateUser(&payLoad)
	if err != nil {
		response := helper.ApiResponse(err.Status(), err.Message(), "error", nil)
		ctx.JSON(err.Status(), response)
		return
	}

	successResponse := helper.ApiResponse(http.StatusCreated, "User Created", "success", usr)
	ctx.JSON(http.StatusCreated, successResponse)
}

func (uh *UserHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		errConv := helper.NewUnprocessableEntityError("Invalid user ID format")
		response := helper.ApiResponse(errConv.Status(), errConv.Message(), "error", nil)
		ctx.JSON(errConv.Status(), response)
		return
	}

	usr, err2 := uh.UserService.GetUserByID(intID)
	if err2 != nil {
		response := helper.ApiResponse(err2.Status(), err2.Message(), "error", nil)
		ctx.JSON(err2.Status(), response)
		return
	}

	successResponse := helper.ApiResponse(http.StatusOK, "User Retrieved", "success", usr)
	ctx.JSON(http.StatusOK, successResponse)
}
