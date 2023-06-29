package main

import (
	"log"
	"neversitup/logictest"
	"neversitup/authen"
	"neversitup/order"
	"neversitup/product"
	"neversitup/user"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tylerb/graceful"
)

func main() {
	log.Println("API start running")
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/logictest", logictest.ValidatePinCode)

	e.POST("/login", authen.Login)

	r := e.Group("/api")
	r.POST("/createUser", user.CreateUser)
	r.PUT("/getUser/:id", user.GetUser)

	r.POST("/createProduct", product.CreateProduct)
	r.GET("/getProduct", product.GetProduct)

	r.POST("/createOrder", order.CreateOrder)
	r.GET("/getOrder", order.GetOrder)
	r.DELETE("/deleteOrder/:id", order.DeleteOrder)

	e.Server.Addr = ":8080"
	graceful.ListenAndServe(e.Server, 5*time.Second)
}

