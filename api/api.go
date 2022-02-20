package api

import (
	"global-auth/api/handler"
	"global-auth/repository/admin"
	"global-auth/repository/user"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouterApp(db mongo.Database) *gin.Engine {
	var (
		v1              = "v1"
		adminPrefix     = v1 + "/admin"
		userPrefix      = v1 + "/user"
		router          = gin.New()
		adminRepository = admin.NewAdminRepository(db)
		userRepository  = user.NewUserRepository(db)
		userHandler     = handler.NewUserHandler(userRepository)
		adminHandler    = handler.NewAdminHandler(adminRepository)
	)

	adminHandler.AdminRoutes(router.Group(adminPrefix + "/auth"))
	userHandler.UserRoutes(router.Group(userPrefix + "/auth"))
	return router
}
