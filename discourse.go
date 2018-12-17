package weserver

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

const (
	Domain = "http://localhost:8080"
	Host   = "https://forum.whyengineer.com"
	Secert = "frankielovelll"
)

type User struct {
	Nonce      string `json:"nonce"`
	Email      string `json:"email"`
	ExternalId string `json:"external_id`
	Username   string `json:"username"`
	Name       string `json:"name"`
}

type Sso struct {
	Sso string `json:"sso"`
	Sig string `json:"sig"`
}

func rndInit() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func rndStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type urlPath struct {
	Url string `json:"url"`
}

func ssoProvider(c echo.Context) (err error) {
	u := new(urlPath)
	if err = c.Bind(u); err != nil {
		return
	}
	rndInit()
	nonce := rndStr(32)
	// url := c.Request().(*standard.URL).URL

	v := make(url.Values)
	v.Set("nonce", nonce)
	v.Set("return_sso_url", u.Url)
	fmt.Println(nonce, u.Url)
	p := base64.StdEncoding.EncodeToString([]byte(v.Encode()))
	h := hmac.New(sha256.New, []byte(Secert))
	h.Write([]byte(p))

	sig := hex.EncodeToString(h.Sum(nil))
	v = make(url.Values)
	v.Set("sso", p)
	v.Set("sig", sig)
	fmt.Println(string(v.Encode()))

	path := fmt.Sprintf("%s/session/sso_provider?%s", Host, v.Encode())
	fmt.Println(path)
	return c.Redirect(http.StatusMovedPermanently, path)

	// return c.NoContent(http.StatusOK)
}

func ssoRedirect(nonce string, user User) (url.Values, error) {
	v := make(url.Values)
	v.Set("nonce", nonce)
	v.Set("email", user.Email)
	v.Set("external_id", user.ExternalId)
	v.Set("username", user.Username)
	v.Set("name", user.Name)

	p := base64.StdEncoding.EncodeToString([]byte(v.Encode()))

	h := hmac.New(sha256.New, []byte(Secert))
	_, err := h.Write([]byte(p))
	if err != nil {
		return nil, err
	}
	sig := hex.EncodeToString(h.Sum(nil))

	v = make(url.Values)
	v.Set("sso", p)
	v.Set("sig", sig)
	return v, nil
}

type disUser struct {
	ID           string `json:"external_id"`
	Nonce        string `json:"nonce"`
	ReturnSsoURL string `json:"return_sso_url"`
	Username     string `json:"username"`
	Admin        string `json:"admin"`
	Email        string `json:"email"`
	Groups       string `json:"groups"`
	Moderator    string `json:"moderator"`
	AvatarURL    string `json:"avatr_url"`
}

func ssoLogin(c echo.Context) (err error) {
	s := new(Sso)
	if err = c.Bind(s); err != nil {
		return
	}
	qs, err := base64.StdEncoding.DecodeString(s.Sso)
	if err != nil {
		fmt.Println(err)
		return
	}
	mac := hmac.New(sha256.New, []byte(Secert))
	mac.Write([]byte(s.Sso))
	expecteMac := mac.Sum(nil)
	sig := hex.EncodeToString(expecteMac)
	if sig != s.Sig {
		err = errors.New("sig unmatch")
		return
	}
	v, err := url.ParseQuery(string(qs))
	if err != nil {
		err = errors.New("query string error")
		return err
	}
	// userInfo := new(disUser)

	// userInfo.Admin = v.Get("admin")
	// userInfo.AvatarURL = v.Get("avatar_url")
	// userInfo.Email = v.Get("email")
	// userInfo.Groups = v.Get("groups")
	// userInfo.ID = v.Get("external_id")
	// userInfo.Moderator = v.Get("moderator")
	// userInfo.Nonce = v.Get("nonce")
	// userInfo.ReturnSsoURL = v.Get("return_sso_url")
	// userInfo.Username = v.Get("username")

	sess, _ := session.Get("session", c)
	sess.Values["admin"] = v.Get("admin")
	sess.Values["avatar_url"] = v.Get("avatar_url")
	sess.Values["email"] = v.Get("email")
	sess.Values["groups"] = v.Get("groups")
	sess.Values["external_id"] = v.Get("external_id")
	sess.Values["moderator"] = v.Get("moderator")
	sess.Values["nonce"] = v.Get("nonce")
	sess.Values["return_sso_url"] = v.Get("return_sso_url")
	sess.Values["username"] = v.Get("username")

	sess.Save(c.Request(), c.Response())

	// fmt.Println(userInfo)
	return c.Redirect(http.StatusMovedPermanently, "http://localhost:8080")

}
