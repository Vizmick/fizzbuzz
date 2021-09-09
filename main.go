package main

import (
	"example.com/fizzbuzz/fb"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/fizzbuzz", fb.GetFizzBuzz)
	router.GET("/frequent", fb.GetMostFrequent)
	router.Run("localhost:8080")
}
