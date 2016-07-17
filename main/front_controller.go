package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes() *gin.Engine {
	//default router with basic middlewares, like login, security, compression
	//alternatively, gin.New() comes without middlewares
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html") //not responding to template updates
	r.GET("/", func(c *gin.Context) {
		// c.String(http.StatusOK, "Hello from %v", "Gin")
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	//optional parameter example: /employees/*id
	r.GET("/employees/:id/vacation", func(c *gin.Context) {
		id := c.Param("id")
		timesOff, ok := TimesOff[id]

		if !ok {
			c.String(http.StatusNotFound, "404 - Page Not Found")
			return
		}

		c.HTML(http.StatusOK, "vacation-overview.html", map[string]interface{}{
			"TimesOff": timesOff,
		})
	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	r.StaticFile("favicon.ico", "./public/img/favicon.ico")
	// r.Static("/assets", "./public")
	r.StaticFS("/assets", http.Dir("./public"))

	return r
}
