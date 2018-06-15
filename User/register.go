package User

import "net/http"

type SessionManage struct {
	cookie *http.Cookie
}

func Session_init() *SessionManage {
	p := new(SessionManage)
	return p
}
