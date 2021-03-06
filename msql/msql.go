package msql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
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
	conn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s ", "root", "root", "docker.nenesan.org", "aria", "disable")
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

func DB_serch_user(name, email string) (n, e string) {//nameとemailを返す
	db := DB_connect()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM users WHERE name ='%s' or email = '%s'", name, email)
	fmt.Println(query)
	rows, err := db.Query(query)
	for rows.Next() {
		rows.Scan(&name, email)
		errCheck(err)
	}

	return name, email
}

func DB_select_user(name string) string {
	db := DB_connect()
	defer db.Close()
	query := fmt.Sprintf("SELECT password  FROM users WHERE name ='%s'", name)
	var password string;
	rows, err := db.Query(query)
	for rows.Next() {
		rows.Scan(&password)
		errCheck(err)
	}
	return password
}
