package auth

import (
	"github.com/gorilla/sessions"
)

func Err(err error) {
	if err != nil {
		panic(err)
	}
}

func login_check(sess *sessions.Session) {
	if sess.Values["name"] == "" {
	}
}
