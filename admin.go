package weserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"golang.org/x/crypto/bcrypt"
)

var dsClient *Client

func adminRouter() {
	dsClient, _ = NewClient(Host, "f05fb8eaabe0e163fe5a609ef08c5dd9d9784d629c1f8dcf47f0e3cfcc7810c3", "frankie")

	admin.GET("/login", ssoProvider)
	admin.GET("/pass", ssoLogin)
	admin.POST("/signup", signup)
	admin.GET("/session", getSession)
	admin.GET("/logout", logout)

}

type loginHandle struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Nonce    string `json:"nonce"`
}

func cryptoPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
func comparePassword(password string, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		// TODO: Properly handle error

		return false
	}

	return true
}

type SessionInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       string `json:"id"`
}

func logout(c echo.Context) (err error) {

	ret := new(Status)
	sess, _ := session.Get("session", c)
	var id string
	var ok bool
	if id, ok = sess.Values["external_id"].(string); !ok {
		ret.Error = -1
		ret.Msg = "external_id no exist"
		return c.JSON(http.StatusOK, ret)
	}
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	endpoint := fmt.Sprintf("/admin/users/%s/log_out.json", id)
	body, _, _ := dsClient.Post(endpoint, []byte(""))
	type Response struct {
		Success string `json:"success"`
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		ret.Error = -1
		ret.Data = "json format error"
		return c.JSON(http.StatusOK, ret)
	}
	ret.Msg = response.Success
	if response.Success == "OK" {
		ret.Error = 0

	} else {
		ret.Error = -1
	}
	return c.JSON(http.StatusOK, ret)
}

func getSession(c echo.Context) (err error) {
	s := new(disUser)
	ret := new(Status)
	sess, _ := session.Get("session", c)
	var ok bool
	ret.Error = -1
	if s.Admin, ok = sess.Values["admin"].(string); !ok {
		ret.Msg = "admin no exist"
		return c.JSON(http.StatusOK, ret)
	}
	if s.AvatarURL, ok = sess.Values["avatar_url"].(string); !ok {
		ret.Msg = "avatar_url no exist"
		return c.JSON(http.StatusOK, ret)
	}
	if s.Email, ok = sess.Values["email"].(string); !ok {
		ret.Msg = "email no exist"
		return c.JSON(http.StatusOK, ret)
	}
	if s.Groups, ok = sess.Values["groups"].(string); !ok {
		ret.Msg = "groups no exist"
		return c.JSON(http.StatusOK, ret)
	}
	if s.ID, ok = sess.Values["external_id"].(string); !ok {
		ret.Msg = "external_id no exist"
		return c.JSON(http.StatusOK, ret)
	}

	if s.Moderator, ok = sess.Values["moderator"].(string); !ok {
		ret.Msg = "moderator no exist"
		return c.JSON(http.StatusOK, ret)
	}

	if s.Nonce, ok = sess.Values["nonce"].(string); !ok {
		ret.Msg = "nonce no exist"
		return c.JSON(http.StatusOK, ret)
	}

	if s.ReturnSsoURL, ok = sess.Values["return_sso_url"].(string); !ok {
		ret.Msg = "return_sso_url no exist"
		return c.JSON(http.StatusOK, ret)
	}

	if s.Username, ok = sess.Values["username"].(string); !ok {
		ret.Msg = "username no exist"
		return c.JSON(http.StatusOK, ret)
	}

	ret.Error = 0
	ret.Data = s
	return c.JSON(http.StatusOK, ret)
}

func login(c echo.Context) (err error) {
	l := new(loginHandle)
	ret := new(Status)
	if err = c.Bind(l); err != nil {
		return
	}
	var result []Weuser
	db.Where("email = ?", l.Email).Find(&result)
	if len(result) == 0 {
		ret.Error = -1
		ret.Msg = "The user does not exist"
		return c.JSON(http.StatusOK, ret)
	}
	if comparePassword(l.Password, result[0].Password) {
		ret.Error = 0
		ret.Msg = "successful"
		ret.Data = result[0]
		sess, _ := session.Get("session", c)
		sess.Values["username"] = result[0].Username
		sess.Values["email"] = l.Email
		sess.Values["id"] = strconv.Itoa(int(result[0].ID))

		sess.Save(c.Request(), c.Response())
		if l.Nonce != "" {
			user := User{
				Email:      l.Email,
				ExternalId: strconv.Itoa(int(result[0].ID)),
				Username:   result[0].Username,
			}
			param, _ := ssoRedirect(l.Nonce, user)

			url := fmt.Sprintf("%s/session/sso_login?%s", Host, param.Encode())
			fmt.Println(url)
			return c.Redirect(http.StatusMovedPermanently, url)
		}
	} else {
		ret.Error = -2
		ret.Msg = "Password not correct"
	}

	return c.JSON(http.StatusOK, ret)
}

type sigupHandle struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Nonce    string `json:"nonce"`
}

func signup(c echo.Context) (err error) {
	s := new(sigupHandle)
	ret := new(Status)
	if err = c.Bind(s); err != nil {
		return
	}
	var result []Weuser
	db.Where("email=?", s.Email).Find(&result)
	if len(result) > 0 {
		ret.Msg = "The email already exists."
		ret.Error = -1
		return c.JSON(http.StatusOK, ret)
	}
	db.Where("username=?", s.Username).Find(&result)
	if len(result) > 0 {
		ret.Msg = "The username already exists."
		ret.Error = -1
		return c.JSON(http.StatusOK, ret)
	}
	//calculate the password hash

	hash, _ := cryptoPassword(s.Password)

	newUser := Weuser{
		Username: s.Username,
		Password: hash,
		Email:    s.Email,
		Level:    5, //normal user
	}

	if dbErr := db.Create(&newUser).Error; dbErr != nil {
		ret.Error = -1
		ret.Msg = dbErr.Error()
		return c.JSON(http.StatusOK, ret)
	}
	ret.Error = 0
	ret.Msg = "successful"
	ret.Data = newUser

	sess, _ := session.Get("session", c)
	sess.Values["username"] = s.Username
	sess.Values["email"] = s.Email
	sess.Values["id"] = strconv.Itoa(int(newUser.ID))
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, ret)
}
