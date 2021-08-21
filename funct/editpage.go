package funct

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

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
func EditHandler(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	Id := r.FormValue("Id")
	Partnum := r.FormValue("Partnum")
	Qty, _ := strconv.Atoi(r.FormValue("Qty"))
	Price, _ := strconv.ParseFloat(r.FormValue("Price"), 32)
	Provider := r.FormValue("Provider")
	Name := r.FormValue("Name")
	Remark := r.FormValue("Remark")
	Date := time.Now().Format("02 January 2006")
	Color := "#ffffff"

	_, err = models.Database.Exec("upDate mybase.mybase set Partnum=?, Qty=?, Price=?,Provider=?,Name=?,Remark=?,Date=? , Color=? where Id=?",
		Partnum, Qty, Price, Provider, Name, Remark, Date, Color, Id)

	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", 301)
}
