package handler

import (
	"global-auth/api/response"
	"global-auth/helper"
	"global-auth/middleware"
	"global-auth/repository/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userRepository user.UserRepository
}

func (r UserHandler) Create(ctx *gin.Context) {
	var user user.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErrValidate(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	if len(user.Email) < 1 {
		ctx.JSON(http.StatusForbidden, response.ResErrValidate(http.StatusForbidden, "email tidak boleh kosong"))
		return
	}
	if len(user.Name) < 1 {
		ctx.JSON(http.StatusForbidden, response.ResErrValidate(http.StatusForbidden, "name tidak boleh kosong"))
		return
	}
	if len(user.Password) < 1 {
		ctx.JSON(http.StatusForbidden, response.ResErrValidate(http.StatusForbidden, "password tidak boleh kosong"))
		return
	}

	user.ID = primitive.NewObjectID()
	hashPassword, err := helper.GeneratePassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResErr(http.StatusNotFound, err.Error()))
		return
	}
	user.Password = string(hashPassword)
	user.CreatedAt = time.Now()
	res, err := r.userRepository.Create(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusAccepted, response.ResOK("success created", res))
	return
}

func (r UserHandler) Login(ctx *gin.Context) {
	m := middleware.Middleware{}
	var user user.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	result, err := r.userRepository.FinOne(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResErr(http.StatusNotFound, "user tidak terdaftar"))
		return
	}
	fail := helper.CompareHashAndPassword([]byte(result.Password), user.Password)
	if fail != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, fail.Error()))
		return
	}
	getToken, err := m.JwtSign(result.ID, false)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	ctx.JSON(http.StatusAccepted, response.ResOK("success", getToken))
	return
}

func (r UserHandler) UserRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/register", r.Create)
	routerGroup.POST("/login", r.Login)
}

func NewUserHandler(repository user.UserRepository) UserHandler {
	return UserHandler{
		userRepository: repository,
	}
}
