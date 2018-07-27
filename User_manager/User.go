package main

import (
	"AriaCTFer/msql"
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
func User_init(name, email, password1, password2 string) *User_manage {
	user := new(User_manage)
	user.name = name
	user.email = email
	user.password1 = password1
	user.password2 = password2
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

// セッターはレシーバをポインタにしないと値が変更されない
func (user *User_manage) SetName(name string) string {
	user.name = name
	return user.name
}

func (user User_manage) GetPpassword() string {
	//パスワード取得
	return user.password1
}

func (user User_manage) _is() (bool, tool.Check_err) {
	is, err := tool.ValidationAll(user.name, user.password1, user.password2, user.email)
	return is, err
}

//func (user User_manage) Register() string {
//	if is { //Validationチェック
//		n, e := msql.DB_serch_user(user.name, user.email)
//		fmt.Println(n)
//		fmt.Println(e)
//		fmt.Println(n == user.name)
//		fmt.Println(e == user.email)
//		if n != user.name || e != user.email {
//			msql.DB_insert(user.name, user.email, user.password1)
//			return ""
//		} else {
//			return "エラー"
//		}
//	} else {
//		return "エラー"
//	}
//}

func main() {
	var user *User_manage = User_init("", "test", "passwrd1", "passowrd")
	is, account_err := user._is()
	//fmt.Println(user.SetName("admin"))
	fmt.Println(is)
	fmt.Println(account_err)
}
