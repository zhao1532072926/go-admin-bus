package router

import (
	"go-admin-bus/bus/controller"
	"go-admin-bus/bus/middleware"
	"go-admin-bus/bus/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRouter(r gin.IRouter, orm *gorm.DB, jwt *utils.JWT) {
	ctrl := controller.NewController(orm, jwt)

	r = r.Group("/api/bus")

	r.POST("/login", ctrl.Login)

	auth := middleware.JWTAuth(jwt)

	r.POST("/first_dd_details", auth, ctrl.CreateFirstDdDetail)
}
