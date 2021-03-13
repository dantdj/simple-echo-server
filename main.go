package main

import (
	"math"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		workDone := doWork()
		c.JSON(200, gin.H{
			"message": "pong",
			"work":    workDone,
		})
	})
	r.GET("/pong", func(c *gin.Context) {
		workDone := doWork()
		c.JSON(200, gin.H{
			"message": "ping",
			"work":    workDone,
		})
	})
	r.Run("0.0.0.0:3000")
}

func doWork() string {
	for i := 0; i < 1000000000; i++ {
		a := 2 * i
		value := float64(a * i)
		math.Sqrt(value)
	}
	return "done"
}
