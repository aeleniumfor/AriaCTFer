package main

import (
	"AriaCTFer/msql"
	"AriaCTFer/stract"
	"AriaCTFer/tool"
	"fmt"
)

type User_manage struct {
	name          string
	email         string
	password_hash string
	password1     string
	password2     string
	is            string
}

// コンストラクタ
// 構造体がpublicでも、そのフィールドが閉じたスコープの場合、
// パッケージ外から{}を使った初期化はできない。
func User_init() *User_manage {
	user := new(User_manage)
	return user
}

// カプセル化
// メソッド名を大文字からはじめることで、publicなスコープにする。
func (user User_manage) GetName() string {
	return user.name
}

func (user User_manage) all_info() User_manage {
	msql.DB_connect()

	return user
}

func (user *User_manage) SetData(name, email, password1, password2 string) {
	user.name = name
	user.email = email
	user.password1 = password1
	user.password2 = password2
}

// セッターはレシーバをポインタにしないと値が変更されない
func (user *User_manage) SetName(name string) {
	user.name = name
}

func (user User_manage) GetPpassword() string {
	//パスワード取得
	return user.password1
}

func (user User_manage) _is() (bool, tool.Check_err) {
	is, err := tool.ValidationAll(user.name, user.password1, user.password2, user.email)
	return is, err
}

func (user User_manage) User_all_reference() []stract.Acount {
	return msql.DB_select_all_user()
}

func (user User_manage) User_registration() {
	is, err := tool.ValidationAll(user.name, user.password1, user.password2, user.email)
	if is {
		fmt.Println("ok")
	} else {

		fmt.Println(err)
	}
}

func main() {
	user := User_init()
	user.SetData("", "test@gmail.com", "1password", "password")
	user.User_registration()
}
