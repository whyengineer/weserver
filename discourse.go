package weserver

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
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
	fmt.Println(string(qs))
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
	nonce := v.Get("nonce")
	endpoint := fmt.Sprintf("/#/login/%s", nonce)
	return c.Redirect(http.StatusMovedPermanently, endpoint)

	// sess, _ := session.Get("session", c)
	// // if sess.Values["username"] == nil || sess.Values["email"] == nil || sess.Values["id"] == nil {
	// // 	ret.Error = -1
	// // 	ret.Msg = "The user does not exist"
	// // 	return c.JSON(http.StatusOK, ret)
	// // }

	// // sess, _ := session.Get("session", c)
	// // if sess.Values["username"] == nil || sess.Values["email"] == nil {
	// // 	return c.Redirect(http.StatusMovedPermanently, "/#/login")
	// // }
	// fmt.Println(sess.Values["username"])
	// fmt.Println(sess.Values["email"])
	// fmt.Println(sess.Values["id"])

	// // username := sess.Values["username"].(string)
	// // email := sess.Values["email"].(string)
	// // id := sess.Values["id"].(string)

	// user := User{
	// 	Email:      "frankie@whyengineer.com",
	// 	ExternalId: "2",
	// 	Username:   "frankie",
	// }
	// fmt.Println(user)
	// param, _ := ssoRedirect(nonce, user)
	// url := fmt.Sprintf("%s/session/sso_login?%s", Host, param.Encode())
	// fmt.Println(url)
	// return c.Redirect(http.StatusMovedPermanently, url)
}
