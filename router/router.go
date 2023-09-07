package router

import (
	"com/adl/et/telco/dte/template/baseapp/alarm"
	"com/adl/et/telco/dte/template/baseapp/handlers"
	"com/adl/et/telco/dte/template/baseapp/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	viper.SetConfigName("config") // Name of the configuration file without extension
	viper.AddConfigPath(".")      // Search the current directory for the configuration file
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	host := viper.GetString("Host")
	port := viper.GetString("Port")

	router := gin.Default()
	log.InitializeLogger()
	logger := log.GetLogger()
	logger.Info("logger plugin initialized")
	alarm.InitializeAlarm()
	alarm.CreateAlarm("alarm plugin initialized", "WARN")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/action", handlers.IntergrationHandler)
	router.Run(host + ":" + port)
}
