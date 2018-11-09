package weserver

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var db *gorm.DB
var hwdb *echo.Group

// ServerInit init echo
func ServerInit() {
	e := echo.New()
	hwdb = e.Group("/hwdb")
	hwdbRouter()

	db = DbConnect()
	db.LogMode(true)
	e.Logger.Fatal(e.Start(":1323"))
}
