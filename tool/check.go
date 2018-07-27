package tool

type Check_err struct {
	emailcheck                 string
	passwordcheck              string
	namecheck                  string
	name_null_check            string
	name_password__match_check string
}

func ValidationAll(name, password1, password2, email string) (bool, Check_err) {
	var err Check_err = Check_err{}
	if emailcheck(email, &err) && password_length_check(password1, password2, &err) && namecheck(name, password1, &err) && name_null_check(name, &err) && name_password_match_check(name, password2, &err) {
		return true, err //すべてのチェックを通ったら
	} else {
		return false, err
	}
}

func emailcheck(email string, err *Check_err) bool {
	if email != "" {
		//emailが空じゃなかったら
		return true
	} else {
		//emailが空の場合
		err.emailcheck = "emailエラー"
		return false
	}
}

func password_length_check(password1, password2 string, err *Check_err) bool {
	if len(password1) >= 8 {
		//パスワードの長さが8以上
		if password1 == password2 {
			return true
		} else {
			err.passwordcheck = "パスワードエラー"
			return false
		}
	} else {
		err.passwordcheck = "パスワードエラー"
		return false
	}
}

func namecheck(name, password1 string, err *Check_err) bool {
	if name != "" {
		//空文字じゃなければ
		if name != password1 {
			//passwordとnameが一致してなければtrueを返す.
			return true
		} else {
			err.namecheck = "nameエラー"
			return false
		}
	} else {
		//空文字で渡ってきた場合
		return false
	}
}

func name_password_match_check(name, password string, err *Check_err) bool {
	if name == password {
		err.name_password__match_check = "name is match to password"
		return false
	} else {
		return true
	}
}

func name_null_check(name string, err *Check_err) bool {
	if name == "" {
		//nameが空文字だったら
		err.name_null_check = "name is empty"
		return false
	} else {
		return true
	}
}
