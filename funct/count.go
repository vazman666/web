package funct

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vazman666/web/models"
)

func Count(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	fmt.Println("conunt work")

	vars := mux.Vars(r)
	Id := vars["Id"]
	fmt.Println("Id=", Id)
	row := models.Database.QueryRow("select * from mybase.mybase where Id = ?", Id)
	prod := models.Product{}
	err := row.Scan(&prod.Id, &prod.Partnum, &prod.Qty, &prod.Price, &prod.Provider, &prod.Name, &prod.Remark, &prod.Date, &prod.Color)
	if err != nil {
		fmt.Println(err)

	}
	if prod.Color == "#ffffff" {

		prod.Color = "#fa8e47"
	} else {

		prod.Color = "#ffffff"
	}

	_, err = models.Database.Exec("upDate mybase.mybase set Color=? where Id=?", prod.Color, Id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}
