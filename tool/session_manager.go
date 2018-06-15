package tool

import (
	"net/http"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func main() {
	sessions.NewCookieStore([]byte("secret")) //シークレットキーの設定
}

type SessionManage struct {
	cookie *http.Cookie
}

func Session_init() *SessionManage {
	p := new(SessionManage)
	return p
}

func Err(err error) {
	if err != nil {
		panic(err)
	}
}

func (s SessionManage) Start(c echo.Context) *sessions.Session {
	sess, err := session.Get("session_id", c) //取得と同時に作成もやってくれると思う.....
	sess.Options = &sessions.Options{
		Path:     "/session",
		MaxAge:   10,
		HttpOnly: true,
	}
	Err(err)
	return sess
}
