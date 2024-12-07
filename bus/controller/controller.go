package controller

import (
	"go-admin-bus/bus/config"
	"go-admin-bus/bus/models"
	"go-admin-bus/bus/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	orm          *gorm.DB
	jwt          *utils.JWT
	serviceNames config.ServiceNames
}

type LoginRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewController(orm *gorm.DB, jwt *utils.JWT, serviceNames config.ServiceNames) *Controller {
	return &Controller{
		orm:          orm,
		jwt:          jwt,
		serviceNames: serviceNames,
	}
}

func (c *Controller) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.Users
	if err := c.orm.Table(c.serviceNames.Users).Where("phone = ? AND password = ?", req.Phone, req.Password).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	claims := map[string]interface{}{
		"phone": user.Phone,
		"service": c.serviceNames.Name,
	}
	
	token, err := c.jwt.GenerateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (c *Controller) CreateDetail(ctx *gin.Context) {
	var detail models.Details
	if err := ctx.ShouldBindJSON(&detail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 从路径中获取维度和子服务类型
	// 路径格式: /api/first/dd/detail
	path := ctx.Request.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的请求路径",
		})
		return
	}
	
	dimension := parts[2]    // first
	subService := parts[3]   // dd
	
	// 使用统一的表名生成函数
	tableName := config.GetTableName(dimension, subService).Details

	// 验证该表是否属于当前服务
	if !c.serviceNames.ValidateTableName(tableName) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "无效的服务类型",
		})
		return
	}

	if err := c.orm.Table(tableName).Create(&detail).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建记录失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": detail,
	})
}
