package bus

import (
	"go-admin-bus/bus/config"
	"go-admin-bus/bus/models"
	"go-admin-bus/bus/router"
	"go-admin-bus/bus/utils"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Bus 业务模块实例
type Bus struct {
	db  *gorm.DB
	jwt *utils.JWT
}

// Init 初始化业务模块
func Init(eng *engine.Engine, r *gin.Engine) error {
	// 1. 加载维度服务配置
	if err := config.LoadConfig("./config/service.yml"); err != nil {
		return err
	}

	// 2. 初始化认证组件
	jwtConfig := config.GetJWTConfig()
	jwt := &utils.JWT{
		Secret:        jwtConfig.Secret,
		TimeoutSecond: jwtConfig.TimeoutSecond,
	}

	// 3. 初始化数据库连接
	dbConn := eng.MysqlConnection()
	db, err := gorm.Open("mysql", dbConn.GetDB("default"))
	if err != nil {
		return err
	}

	// 4. 生成数据表
	if err := models.GenerateModels(db); err != nil {
		return err
	}

	// 5. 注册路由
	router.InitRouter(r, db, jwt)

	// 6. 初始化 Generators (需要在配置加载后)
	generators := initGenerators()
	eng.AddGenerators(generators)

	return nil
} 