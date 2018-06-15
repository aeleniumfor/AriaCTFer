package handler

import (
	"net/http"
	"github.com/labstack/echo"

	"AriaCTFer/tool"
	"AriaCTFer/msql"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"fmt"
)

func Err(err error) {
	if err != nil {
		panic(err)
	}
}

func IndexPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		fmt.Println(sess.Values)
		if sess.Values["username"] == nil {
			return c.Render(http.StatusOK, "index.html", nil)
		} else {
			return c.Render(http.StatusOK, "index_login.html", map[string]interface{}{"username": sess.Values["username"]})
		}
	}
}

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
// 認証
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

//登録
func Register_GET_Page() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "register.html", nil)
	}
}
func Register_POST_Page() echo.HandlerFunc {
	return func(c echo.Context) error {
		var name string = c.FormValue("name")
		var email string = c.FormValue("email")
		var password1 string = c.FormValue("password1")
		var password2 string = c.FormValue("password2")
		var is bool = tool.ValidationAll(name, email, password1, password2)
		if is == true {
			password1, _ := tool.HashPassword(password1)
			msql.DB_insert(name, email, password1)
			return c.Redirect(http.StatusMovedPermanently, "login_get")
		} else {
			return c.Render(http.StatusOK, "register.html", map[string]interface{}{"message": is})
		}

	}
}

//ログイン
func Login_GET_Page() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", nil)
	}
}

func Login_POST_Page() echo.HandlerFunc {
	return func(c echo.Context) error {
		var name string = c.FormValue("name")
		//var email string = c.FormValue("email")
		var password1 string = c.FormValue("password1")
		if tool.CheckPasswordHash(password1, msql.DB_select_user(name)) {
			sess, err := session.Get("session", c)
			sess.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   86400 * 7,
				HttpOnly: true,
			}
			Err(err);
			sess.Values["username"] = name
			sess.Save(c.Request(), c.Response())

			return c.Redirect(http.StatusFound, "/")
		} else {
			return c.Render(http.StatusOK, "login.html", nil)
		}
	}
}

//ログアウト
func Logout_GET_Page() echo.HandlerFunc {
	return func(c echo.Context) error {

		sess, err := session.Get("session", c)
		sess.Options.MaxAge = -1
		delete(sess.Values, "username")
		sess.Save(c.Request(), c.Response())
		Err(err)
		return c.Redirect(http.StatusFound, "/")
	}
}
