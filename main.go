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

	"go-admin-bus/pages"
	"go-admin-bus/tables"

	"go-admin-bus/bus"
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
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/admin", pages.GetDashBoard)
	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	// 初始化 bus
	if err := bus.Init(eng, r); err != nil {
		panic(err)
	}

	_ = r.Run(":8000")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
