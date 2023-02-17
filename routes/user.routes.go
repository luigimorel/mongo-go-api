package routes

import (
	"github.com/morelmiles/mongo-auth/controllers"
	"github.com/morelmiles/mongo-auth/middleware"
	"github.com/morelmiles/mongo-auth/services"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {

	router := rg.Group("users")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/profile", uc.userController.GetProfile)
}
