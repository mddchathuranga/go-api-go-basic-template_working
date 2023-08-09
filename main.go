package main

import "com/adl/et/telco/dte/template/baseapp/router" /*
	"com/adl/et/telco/dte/template/baseapp/handlers"

	"github.com/gin-gonic/gin"
	"github.com/mddchathuranga/DTEAlarmingPluginGoLang/alarmer"

	"github.com/mddchathuranga/DTELoggingPluginGoLang/logging"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"*/

func main() {
	router.InitRouter()
	/*	router := gin.Default()
		logging.InitializeLogger()
		logger := logging.GetLogger()
		logger.Info("logger plugin initialized")
		alarmer.InitializeAlarm()
		alarmer.CreateAlarm("alarm plugin initialized", "WARN")
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.POST("/action", handlers.IntergrationHandler)
		// Start the server
		router.Run(":8080")*/
}
