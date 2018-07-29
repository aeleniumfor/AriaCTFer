package msql

import (
	"AriaCTFer/stract"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Create_at time.Time
	Update_at time.Time
}

func errCheck(err error) {
	if err != nil {
		panic(err.Error())
	}

}

func DB_connect() *sql.DB {
	//dbへのコネクション関数
	var db_user string = os.Getenv("db_user")
	var db_password string = os.Getenv("db_password")
	var db_host string = os.Getenv("db_host")
	var db_name string = os.Getenv("db_name")
	var db_sslmode string = os.Getenv("db_sslmode")

	db_user = "root"
	db_password = "root"
	db_host = "docker.nenesan.org"
	db_name = "aria"
	db_sslmode = "disable"

	conn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s ", db_user, db_password, db_host, db_name, db_sslmode)
	db, err := sql.Open("postgres", conn)
	errCheck(err)
	return db
}

func DB_insert(name string, email string, password string) string {
	// データベースへの挿入
	db := DB_connect()
	defer db.Close()
	var last_id string
	query := "INSERT INTO users(name,email,password) VALUES($1,$2,$3) RETURNING id"
	err := db.QueryRow(query, name, email, password).Scan(&last_id)
	errCheck(err)
	return last_id
}

func DB_serch_user(name, email string) (n, e string) { //nameとemailを返す
	db := DB_connect()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM users WHERE name ='%s' or email = '%s'", name, email)
	fmt.Println(query)
	rows, err := db.Query(query)
	for rows.Next() {
		rows.Scan(&name, &email)
		errCheck(err)
	}

	return name, email
}

func DB_select_user(name string) string {
	db := DB_connect()
	defer db.Close()
	query := fmt.Sprintf("SELECT password  FROM users WHERE name ='%s'", name)
	var password string
	rows, err := db.Query(query)
	for rows.Next() {
		rows.Scan(&password)
		errCheck(err)
	}
	return password
}

func DB_select_all_user() []stract.Acount {
	db := DB_connect()
	defer db.Close()
	var account_array []stract.Acount
	account := stract.Acount{}
	rows, _ := db.Query("SELECT id,name,email FROM users")
	for rows.Next() {
		var id int
		var name string
		var email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			panic(err.Error())
		}
		account.Id = id
		account.Name = name
		account.Email = email
		account_array = append(account_array, account)
	}
	return account_array
}

func main() {
	DB_select_all_user()
}
