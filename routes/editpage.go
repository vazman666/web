package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vazman666/web/models"
)

func EditPage(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	vars := mux.Vars(r)
	Id := vars["Id"]

	row := models.Database.QueryRow("select * from mybase.mybase where Id = ?", Id)
	prod := models.Product{}
	err := row.Scan(&prod.Id, &prod.Partnum, &prod.Qty, &prod.Price, &prod.Provider, &prod.Name, &prod.Remark, &prod.Date, &prod.Color)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("templates/edit.html")
		tmpl.Execute(w, prod)
	}
}
