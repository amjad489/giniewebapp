package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		userAgentString := c.GetHeader("User-Agent")
		ua := ua.Parse(userAgentString)
		os := func() string { if ua.Desktop { return ua.OS + " " + ua.OSVersion } else if ua.Mobile{ return ua.OS + " " + ua.OSVersion} else if ua.Mobile{ return ua.OS + " " + ua.OSVersion} else { return "unknown" } }()
		fmt.Println(ua)
		c.JSON(200, gin.H{
			"message": "pong",
			"client": ua.Name,
			"os": os,
		})
	})
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}