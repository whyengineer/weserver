package weserver

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var db *gorm.DB
var hwdb *echo.Group

// ServerInit init echo
func ServerInit() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://192.168.31.94:8080", "http://localhost:8080", "http://192.168.31.129:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	hwdb = e.Group("/hwdb")
	hwdbRouter()

	db = DbConnect()
	db.LogMode(true)
	e.Logger.Fatal(e.Start(":1323"))
}
