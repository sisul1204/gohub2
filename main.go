package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub2/bootstrap"
	btsConfig "gohub2/config"
	"gohub2/pkg/config"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	var env string
	flag.StringVar(&env, "env", "", "加载 .env文件，如 --env=testing 加载的是.env.testing文件")
	flag.Parse()
	config.InitConfig(env)

	router := gin.New()
	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))

	if err != nil {
		fmt.Println(err.Error())
	}
}
