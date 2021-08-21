package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/vazman666/web/funct"
	"github.com/vazman666/web/models"
)

//var m sync.Mutex
func main() {

	db, err := sql.Open("mysql", "vazman:rbhgbxb1@/mybase")

	if err != nil {
		log.Println(err)
	}
	models.Database = db
	defer db.Close()

	s1 := mux.NewRouter()

	s1.HandleFunc("/", funct.IndexHandler)
	s1.HandleFunc("/create", funct.CreateHandler)
	s1.HandleFunc("/Checkout", funct.Checkout)
	s1.HandleFunc("/add/{Id:[0-9]+}", funct.Count)
	s1.HandleFunc("/del", funct.Del)
	s1.HandleFunc("/edit/{Id:[0-9]+}", funct.EditPage).Methods("GET")
	s1.HandleFunc("/edit/{Id:[0-9]+}", funct.EditHandler).Methods("POST")
	s1.HandleFunc("/delete/{Id:[0-9]+}", funct.DeleteHandler)

	http.Handle("/", s1)

	fmt.Println("Server is listening...8181")
	http.ListenAndServe(":8181", nil)
}
