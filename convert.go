package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id       int
	Partnum  string
	Qty      int
	Price    float32
	Provider string
	Name     string
	Remark   string
	Date     string
	Color    string
}

var DATA []Product
var Db *sql.DB
var database *sql.DB

func main() {
	Db, err := sql.Open("sqlite3", "main.db") //открываем БД
	if err != nil {
		fmt.Println(nil)
		return

	}

	rows, err := Db.Query("SELECT * FROM data") //Загружаем данные из БД
	if err != nil {
		fmt.Println(nil)
		return
	}
	db, err := sql.Open("mysql", "vazman:rbhgbxb1@/mybase")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()

	for rows.Next() {
		var tmp Product
		err = rows.Scan(&tmp.Partnum, &tmp.Qty, &tmp.Price, &tmp.Provider, &tmp.Name, &tmp.Remark, &tmp.Date)
		fmt.Println(tmp)
		Color := "#ffffff"
		_, err = database.Exec("insert into mybase.mybase (Partnum, Qty, Price,Provider,Name,Remark,Date,Color) values (?, ?, ?, ?, ?, ?, ?, ?)",
			tmp.Partnum, tmp.Qty, tmp.Price, tmp.Provider, tmp.Name, tmp.Remark, tmp.Date, Color)
		if err != nil {
			log.Println(err)
		}

		DATA = append(DATA, tmp)
	}
	fmt.Println(DATA)

}
