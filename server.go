package weserver

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

var db *gorm.DB
var hwdb *echo.Group
var admin *echo.Group

var staticPath string

func home(c echo.Context) error {

	sess, _ := session.Get("session", c)
	fmt.Println(sess.Values["username"])
	fmt.Println(sess.Values["email"])

	return c.File(staticPath + "/index.html")
	// return c.String(http.StatusOK, "Hello, World!")
}
func test(c echo.Context) error {

	sess, _ := session.Get("session", c)

	sess.Values["username"] = "frankie"
	sess.Values["email"] = "gmail"
	sess.Save(c.Request(), c.Response())

	// return c.File(staticPath + "/index.html")
	return c.String(http.StatusOK, "Hello, World!")
}

func dbConn(dbconn string) {
	db = DbConnect(dbconn)
}

// ServerInit init echo
func ServerInit(dbconn string) {
	path, _ := filepath.Abs(".")
	staticPath = filepath.Join(path, "dist")

	e := echo.New()
	e.Static("/", staticPath)

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(Secert))))
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://192.168.31.94:8080", "http://localhost:8080", "http://192.168.31.129:8080", "http://192.168.31.94:1323"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/test", test)

	e.GET("/", home)
	e.GET("/sso", ssoLogin)
	hwdb = e.Group("/hwdb")
	hwdbRouter()

	admin = e.Group("/admin")
	adminRouter()
	dbConn(dbconn)

	db.LogMode(true)
	e.Logger.Fatal(e.Start(":2018"))

}
