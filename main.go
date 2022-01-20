package main

import (
	"newsfeeder/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatbase()

	r.Run()
}
