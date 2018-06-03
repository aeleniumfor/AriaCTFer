package handler

import (
	"net/http"
	"github.com/labstack/echo"
	"AriaCTFer/tool"
	"AriaCTFer/msql"
	"fmt"
)

func IndexPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{"name": "Dolly!"})
	}
}

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
// 認証
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

//登録
func Register_GET_Page() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "register.html", map[string]interface{}{"name": "Dolly!"})
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

		return c.Render(http.StatusOK, "login.html", map[string]interface{}{"name": "Dolly!"})
	}
}

func Login_POST_Page() echo.HandlerFunc {
	return func(c echo.Context) error {
		var name string = c.FormValue("name")
		//var email string = c.FormValue("email")
		var password1 string = c.FormValue("password1")
		if tool.CheckPasswordHash(password1, msql.DB_select_user(name)) {
			fmt.Println("test")
			return c.Render(http.StatusOK, "index.html", map[string]interface{}{"name": "Dolly!"})
		} else {
			return c.Render(http.StatusOK, "login.html", map[string]interface{}{"name": "Dolly!"})
		}
	}
}
