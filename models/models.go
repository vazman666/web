package models

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
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

var Database *sql.DB
var M sync.Mutex
