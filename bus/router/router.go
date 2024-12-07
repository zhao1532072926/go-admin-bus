package router

import (
	"fmt"
	"go-admin-bus/bus/config"
	"go-admin-bus/bus/controller"
	"go-admin-bus/bus/middleware"
	"go-admin-bus/bus/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRouter(r gin.IRouter, orm *gorm.DB, jwt *utils.JWT) {
	services := config.GetAllServiceNames()
	
	for _, service := range services {
		ctrl := controller.NewController(orm, jwt, service)
		
		// 为每个第一维度创建路由组 /api/first
		group := r.Group(fmt.Sprintf("/api/%s", service.Name))
		
		// 登录接口 /api/first/login
		group.POST("/login", ctrl.Login)
		
		// 需要认证的接口
		auth := middleware.JWTAuth(jwt, service)
		
		// 为每个子服务创建详情接口 /api/first/dd/detail
		for _, subService := range service.SubServices {
			subGroup := group.Group(fmt.Sprintf("/%s", subService))
			subGroup.POST("/details", auth, ctrl.CreateDetail)
		}
	}
}
