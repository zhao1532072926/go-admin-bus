package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-admin-bus/bus/models"
	"go-admin-bus/bus/utils"
	"net/http"
)

type Controller struct {
	orm *gorm.DB
	jwt *utils.JWT
}

type LoginRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewController(orm *gorm.DB, jwt *utils.JWT) *Controller {
	return &Controller{
		orm: orm,
		jwt: jwt,
	}
}

func (c *Controller) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.FirstUsers
	if err := c.orm.Where("phone = ? AND password = ?", req.Phone, req.Password).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// 生成 JWT token
	claims := map[string]interface{}{
		"phone": user.Phone,
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

func (c *Controller) CreateFirstDdDetail(ctx *gin.Context) {
	var detail models.FirstDdDetail
	if err := ctx.ShouldBindJSON(&detail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	if err := c.orm.Create(&detail).Error; err != nil {
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
