package main

import (
	"L-F/configs/database"
	"L-F/configs/router"

	"L-F/configs/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	t := gin.Default()
	database.InitDB()
	router.Router(t)
	t.Run()
}
