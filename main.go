package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"shopping/api"
	_ "shopping/docs"
	"shopping/utils/graceful"
	"time"
)

// @title 电商项目
// @description 电商项目
// @version 1.0
// @contact.name golang技术栈
// @contact.url https://www.golang-tech-stack.com

// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()
	registerMiddlers(r)
	api.RegisterHandlers(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Printf("lidten:%s\n", err)
		}
	}()
	graceful.ShowdownGin(srv, time.Second*5)
}

func registerMiddlers(r *gin.Engine) {
	r.Use(
		gin.LoggerWithFormatter(
			func(params gin.LogFormatterParams) string {
				return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s %s\"\n",
					params.ClientIP,
					params.TimeStamp.Format(time.RFC3339),
					params.Method,
					params.Path,
					params.Request.Proto,
					params.StatusCode,
					params.Latency,
					params.ErrorMessage,
				)
			}))
	r.Use(gin.Recovery())
}
