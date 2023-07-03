package main

import (
	"log"
	"neversitup/api/order"
	"neversitup/api/product"
	"neversitup/api/user"
	"neversitup/authen"
	"neversitup/logictest"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tylerb/graceful"

	_ "neversitup/docs"

	config "github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title NeversitupTest
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath
// @schemes http
func main() {
	log.Println("API start running")
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/logictest", logictest.ValidatePinCode)

	e.POST("/login", authen.Login)

	r := e.Group(config.GetString("service.endpoint"))
	r.POST("/register", user.CreateUser)
	r.GET("/getUser", user.GetUser)

	r.POST("/createProduct", product.CreateProduct)
	r.GET("/getProduct", product.GetProduct)

	r.POST("/createOrder", order.CreateOrder)
	r.GET("/getOrder", order.GetOrder)
	r.DELETE("/deleteOrder/:id", order.DeleteOrder)

	e.Server.Addr = ":" + config.GetString("service.port")
	graceful.ListenAndServe(e.Server, 5*time.Second)
}

func init () {
	config.SetConfigFile("config/dev.yml")
	if err := config.ReadInConfig(); err != nil {
		log.Fatal("Fatal error env config (file)")
	}
}

