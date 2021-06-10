package handler


import (
	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	"net/http"
	//"sync/atomic"
	//"fmt"
)

func StartServer() {
	e := echo.New()
	e.GET("/", GetPage)
	e.Start(":3000")
}

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "homepage.html", map[string]interface{} {
		"name": "home",
	})
}
func GetPage(c echo.Context) error{
	return c.String(http.StatusOK, "Webserver says hello")
}

/*
func main() {
	StartServer()
} */