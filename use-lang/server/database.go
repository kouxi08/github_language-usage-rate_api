package server

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func setup() (*sql.DB, error) {
	var err error
	db, err = sql.Open("mysql", "gouser:gopass@tcp(127.0.0.1:3306)/go_mysql8_development?charset=utf8mb4")
	if err != nil {
		return nil, fmt.Errorf("error: Could not connect to the database: %v", err)
	}
	return db, nil
}

func Dbconnection() {

	db, err := setup()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 接続が終了したらクローズする
	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Println("データベース接続失敗")
		return
	} else {
		fmt.Println("データベース接続成功")

	}
}

func Insert(user string, m map[string]float64) {

	count := m["count"]
	delete(m, "count")

	db, err := setup()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	// github_dataへのSQLの準備
	ins, err := db.Prepare("INSERT INTO github_data VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	ins1, err := db.Prepare("INSERT INTO repository VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins1.Close()

	// SQLの実行

	for lang, rate := range m {
		_, err := ins.Exec(user, lang, rate)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = ins1.Exec(user, count)
	if err != nil {
		log.Fatal(err)
	}

}
