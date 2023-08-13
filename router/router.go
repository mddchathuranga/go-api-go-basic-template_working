package router

import (
	"com/adl/et/telco/dte/template/baseapp/alarm"
	"com/adl/et/telco/dte/template/baseapp/handlers"
	"com/adl/et/telco/dte/template/baseapp/log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	router := gin.Default()
	log.InitializeLogger()
	logger := log.GetLogger()
	logger.Info("logger plugin initialized")
	alarm.InitializeAlarm()
	alarm.CreateAlarm("alarm plugin initialized", "WARN")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/action", handlers.IntergrationHandler)
	router.Run(":8080")
}
