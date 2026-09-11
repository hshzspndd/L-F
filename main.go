package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delve/delve/pkg/config"
)

func main() {
	config.LoadConfig()
	t := gin.Default()

	t.Run()
}
