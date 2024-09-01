package main

import (
	"abc/api"
	"abc/config"
	schema "abc/db"
	"abc/db/inmemory"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func setupRouter() *gin.Engine {
	if schema.DB == nil {
		schema.DB = inmemory.Get()
	}
	r := gin.Default()
	routes := r.Group("/api")
	{
		routes.GET("/classes", api.GetClasses)
		routes.POST("/classes", api.CreateClasses)
	}
	{
		routes.POST("/booking/:user_id", api.BookClass)
		routes.GET("/booking", api.Bookings)
	}
	return r
}

func main() {
	// TODO: add sqlite support
	// schema.DB = sqlite.Get()
	r := setupRouter()
	err := r.Run(fmt.Sprintf(":%d", config.PORT))
	if err != nil {
		slog.Error("Server Crashed", "err", err)
	}
}
