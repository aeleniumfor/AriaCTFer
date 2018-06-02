package handler

import (
	"net/http"
	"github.com/labstack/echo"
	"AriaCTFer/tool"
)

func IndexPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{"name": "Dolly!"})
	}
}

//認証
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
			return c.Render(http.StatusOK, "register.html", map[string]interface{}{"name": is})
		} else {
			return c.Render(http.StatusOK, "register.html", map[string]interface{}{"name": is})
		}

	}
}
