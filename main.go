package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	panic(r.Run(":80"))
}
