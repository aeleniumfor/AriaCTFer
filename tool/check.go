package tool

func ValidationAll(name, password1, password2, email string) bool {
	if emailcheck(email) && passwordcheck(password1, password2) && namecheck(name, password1) {
		return true //すべてのチェックを通ったら
	} else {
		return false
	}
}

func emailcheck(email string) bool {
	if email != "" {
		//emailが空じゃなかったら
		return true
	} else {
		//emailが空の場合
		return false
	}
}

func passwordcheck(password1, password2 string) bool {
	if len(password1) >= 8 {
		//パスワードの長さが8以上
		if password1 == password2 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func namecheck(name, password1 string) bool {
	if name != "" {
		//空文字じゃなければ
		if name != password1 {
			//passwordとnameが一致してなければtrueを返す.
			return true
		} else {
			return false
		}
	} else {
		//空文字で渡ってきた場合
		return false
	}
}
