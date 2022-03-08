package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Default struct {
	IP      string `json:"ip"`
	Message string `json:"message"`
}

type DateTime struct {
	DateTime time.Time `json:"datetime"`
}

type Msg struct {
	Message string `json:"message"`
}

// Group
// Path: /
func addDefault(rg *gin.RouterGroup) {
	rp := rg.Group("/")
	res := &Default{}

	rp.GET("/", func(c *gin.Context) {
		res.IP = c.ClientIP()
		res.Message = "💡" // UTF-32LE, Unicode U+1F4A1
		c.JSON(http.StatusOK, res)
	})
}

// Group
// Path: /ping
func addPingRoutes(rg *gin.RouterGroup) {
	rp := rg.Group("/ping")
	res := &Msg{}

	rp.GET("/", func(c *gin.Context) {
		res.Message = "pong"
		c.JSON(http.StatusOK, res)
	})
}

// Group
// Path: /date
func addGetDate(rg *gin.RouterGroup) {
	rp := rg.Group("/date")
	res := &DateTime{}

	rp.GET("/", func(c *gin.Context) {
		res.DateTime = time.Now().UTC() // "2022-03-04T22:05:25.23544Z"
		c.JSON(http.StatusOK, res)
	})
}
