package main

import (
	"net/http"
	"time"
	"todomvc-app-template-golang/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(cors())
	initRouter(engine)
	server := initServer(engine)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func initRouter(engin *gin.Engine) {
	r := engin.Group("api")
	{
		r.POST("add", handler.Add)
		r.POST("del", handler.Del)
		r.POST("update", handler.Update)
		r.POST("find", handler.Find)
	}
}

func initServer(engine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           "8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Orgin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin,Content-Type")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		} else {
			ctx.Next()
		}

	}
}
