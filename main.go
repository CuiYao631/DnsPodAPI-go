package main

import (
	"DDNS-c/controller"
	"DDNS-c/usecase"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var LISTEN = os.Getenv("LISTEN")
var HOST = os.Getenv("HOST")

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	client := &http.Client{}
	uc := usecase.MakeUsecase(client, HOST)
	crtl := controller.MakeController(uc)

	g := e.Group("/dnspod")
	g.POST("/updnns", crtl.SetUpRecord)
	g.GET("/test", crtl.CreateRecord)
	g.POST("/post", crtl.ListRecord)

	e.Logger.Fatal(e.Start(":3000"))
}
