package router

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Dzirael/go-curenncy/internal/api/controllers"
	"github.com/Dzirael/go-curenncy/internal/api/middlewares"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("data/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Middlewares
	// app.Use(gin.Logger())
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	// Cache store
	store := persistence.NewInMemoryStore(time.Second)

	// Routes
	// Rate
	app.GET("/rate", cache.CachePage(store, 5*time.Second, controllers.GetRate))
	// Subscription
	app.POST("/subscribe", controllers.CreateSubscribe)
	app.POST("/sendEmails", controllers.GetRate)
	// SWAGGER
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return app
}
