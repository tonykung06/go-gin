package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
		// c.SetAccepted("application/json", "application/x-www-form-urlencoded", "text/html")
		// fmt.Println("updated client's Accept header", c.Accepted, c.Request.Header.Get("Accept"))
		//func (c *Context) Negotiate(code int, config Negotiate)

		id := c.Param("id") //returns empty string when the item doesn't exist

		//to pull in type query arg from /employees/42/vacation?query=Holiday
		//c.Query("type") gives "Holiday"
		//with default value: c.DefaultQuery("type", "PTO")

		//handle form post
		//POST /employees/42/vacation/add
		//type=PTO&amount=8&startDate=12082016
		//c.PostForm("amount") gives "8"
		//with default value: c.DefaultPostForm("amount", "0")

		timesOff, ok := TimesOff[id]

		if !ok {
			c.String(http.StatusNotFound, "404 - Page Not Found")
			return
		}

		c.HTML(http.StatusOK, "vacation-overview.html", gin.H{
			"TimesOff": timesOff,
		})
	})

	r.POST("/employees/:id/vacation/new", func(c *gin.Context) {
		var timeOff TimeOff
		err := c.BindJSON(&timeOff)
		fmt.Println(timeOff)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		id := c.Param("id")
		timesOff, ok := TimesOff[id]
		if !ok {
			TimesOff[id] = []TimeOff{}
		}
		TimesOff[id] = append(timesOff, timeOff)
		// default response: c.Status(http.StatusOK)

		// c.JSON(http.StatusCreated, gin.H{
		// 	id: 123,
		// })
		c.JSON(http.StatusCreated, &timeOff)
	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", employees)
	})

	//cannot have admin.Get("/employees/add"), which will confuse gin with the following route
	//the following needs to be the only route to serve /admin/employees/<whatever>
	admin.GET("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {
			c.HTML(http.StatusOK, "admin-employee-add.html", nil)
			return
		}
		employee, ok := employees[id]

		if !ok {
			c.String(http.StatusNotFound, "404 - Not Found")
			return
		}

		c.HTML(http.StatusOK, "admin-employee-edit.html", map[string]interface{}{
			"Employee": employee,
		})
	})
	admin.POST("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {
			// pto, ptoErr := strconv.ParseFloat(c.PostForm("pto"), 32)
			// if ptoErr != nil {
			// 	c.String(http.StatusBadRequest, ptoErr.Error())
			// 	return
			// }

			startDate, startDateErr := time.Parse("2006-01-02", c.PostForm("startDate"))
			if startDateErr != nil {
				c.String(http.StatusBadRequest, startDateErr.Error())
				return
			}

			var emp Employee
			// emp.FirstName = c.PostForm("firstName")
			// emp.LastName = c.PostForm("lastName")
			// emp.Position = c.PostForm("position")
			// emp.TotalPTO = float32(pto)
			bindErr := c.Bind(&emp)
			if bindErr != nil {
				c.String(http.StatusBadRequest, startDateErr.Error())
				return
			}
			emp.ID = 42
			emp.Status = "Active"
			emp.StartDate = startDate
			employees["42"] = emp

			c.Redirect(http.StatusMovedPermanently, "/admin/employees/42")
		}
	})

	r.StaticFile("favicon.ico", "./public/img/favicon.ico")
	// r.Static("/assets", "./public")
	r.StaticFS("/assets", http.Dir("./public"))

	return r
}
