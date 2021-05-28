package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"time"
)

var INFO = "%v(%v) -> %v:%v\n"

func FakeData(e echo.Context) error {

	fmt.Printf(INFO, e.Request().Method, e.Request().Proto, e.Request().Host, e.Request().URL)

	resp := map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success",
		"Data":    DataFake(),
	}

	fmt.Println("Response: ", resp)

	return e.JSON(http.StatusOK, resp)
}

func InitRouter(e *echo.Echo) {

	gv1 := e.Group("private/fake/data/api/v1.0/")

	gv1.GET("coffee", FakeData)

}

func AppRun() {
	e := echo.New()

	// Set timeout and disable keep alive
	e.Server.SetKeepAlivesEnabled(false)
	e.Server.ReadTimeout = time.Minute * 60
	e.Server.WriteTimeout = time.Minute * 60

	e.Use(middleware.CORS())

	InitRouter(e)

	e.Start(":9096")

}

func main() {
	AppRun()
}
