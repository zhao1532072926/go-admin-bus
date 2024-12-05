package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	_ "github.com/GoAdminGroup/themes/sword"                      // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"

	"go-admin-bus/models"
	"go-admin-bus/pages"
	"go-admin-bus/tables"

	"go-admin-bus/bus"
	busRouter "go-admin-bus/bus/router"
	"go-admin-bus/bus/utils"
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()


	if err := eng.AddConfigFromYAML("./config.yml").
		AddGenerators(tables.Generators).
		AddGenerators(bus.Generators).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/admin", pages.GetDashBoard)
	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	models.Init(eng.MysqlConnection())

	// 注册 bus 路由
	jwt := &utils.JWT{
		Secret:        "your-secret-key",
		TimeoutSecond: 86400, // 24小时
	}
	busRouter.InitRouter(r, models.GetDB(), jwt)

	_ = r.Run(":8000")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
