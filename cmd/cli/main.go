package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jwDevOps/atlas-backend/internal/config"
	"github.com/jwDevOps/atlas-backend/internal/controller"
	"github.com/jwDevOps/atlas-backend/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.Load()
	log.Printf("loading database %s from configuration.\n", cfg.Database.Name)

	dbManager := database.Init(cfg.Database.Name)
	dbManager.CreateSchema()

	router := gin.Default()
	//router.SetTrustedProxies([]string{"100.104.187.87"})

	router.GET("/publishers", func(c *gin.Context) {
		controller.GetAllPublishers(c, dbManager)
	})
	router.POST("/publishers", func(c *gin.Context) {
		controller.AddPublishers(c, dbManager)
	})

	router.DELETE("/publishers/:id", func(c *gin.Context) {
		controller.DeletePublisher(c, dbManager)
	})
	router.Run(":8001")
}
