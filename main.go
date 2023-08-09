package main

import (
	"com/adl/et/telco/dte/template/baseapp/handlers"

	"github.com/gin-gonic/gin"
	"github.com/mddchathuranga/DTEAlarmingPluginGoLang/alarmer"

	"github.com/mddchathuranga/DTELoggingPluginGoLang/logging"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	logging.InitializeLogger()
	logging.Info("this is info Log")
	alarmer.InitializeAlarm()
	alarmer.CreateAlarmEx("this is an example alarm with critical")
	alarmer.CreateAlarm("Another alarm message", "WARN")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/action", handlers.IntergrationHandler)
	// Start the server
	router.Run(":8080")
}
