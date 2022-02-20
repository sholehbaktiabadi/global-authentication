package handler

import (
	"global-auth/api/response"
	"global-auth/helper"
	"global-auth/middleware"
	"global-auth/repository/admin"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminHandler struct {
	adminRepository admin.AdminRepository
}

func (r AdminHandler) Create(ctx *gin.Context) {
	var admin admin.Admin
	err := ctx.ShouldBindJSON(&admin)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErrValidate(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	if len(admin.Email) < 1 {
		ctx.JSON(http.StatusForbidden, response.ResErrValidate(http.StatusForbidden, "email tidak boleh kosong"))
		return
	}
	if len(admin.Name) < 1 {
		ctx.JSON(http.StatusForbidden, response.ResErrValidate(http.StatusForbidden, "name tidak boleh kosong"))
		return
	}
	if len(admin.Password) < 1 {
		ctx.JSON(http.StatusForbidden, response.ResErrValidate(http.StatusForbidden, "password tidak boleh kosong"))
		return
	}
	if len(admin.Role) < 1 {
		ctx.JSON(http.StatusForbidden, response.ResErrValidate(http.StatusForbidden, "role tidak boleh kosong"))
		return
	}

	admin.ID = primitive.NewObjectID()
	hashPassword, err := helper.GeneratePassword(admin.Password)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResErr(http.StatusNotFound, "admin tidak terdaftar"))
		return
	}
	admin.Password = string(hashPassword)
	admin.CreatedAt = time.Now()
	res, err := r.adminRepository.Create(ctx, admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusAccepted, response.ResOK("success created", res))
	return
}

func (r AdminHandler) Login(ctx *gin.Context) {
	m := middleware.Middleware{}
	var admin admin.Admin
	err := ctx.ShouldBindJSON(&admin)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	result, err := r.adminRepository.FinOne(ctx, admin)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResErr(http.StatusNotFound, err.Error()))
		return
	}
	fail := helper.CompareHashAndPassword([]byte(result.Password), admin.Password)
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

func (r AdminHandler) AdminRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/register", r.Create)
	routerGroup.POST("/login", r.Login)
}

func NewAdminHandler(repository admin.AdminRepository) AdminHandler {
	return AdminHandler{
		adminRepository: repository,
	}
}
