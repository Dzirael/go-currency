package api

import (
	"fmt"
	"log"

	"github.com/Dzirael/go-curenncy/internal/api/router"
	"github.com/Dzirael/go-curenncy/internal/pkg/config"
	"github.com/Dzirael/go-curenncy/internal/pkg/db"
	"github.com/gin-gonic/gin"
)

func setConfiguration() {
	config.Setup()
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run() {
	setConfiguration()
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	if err := web.Run(":" + conf.Server.Port); err != nil {
		log.Fatalf("Rest client failed %s", err.Error())
	}
}
