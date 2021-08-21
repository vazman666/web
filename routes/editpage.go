package routes

import (
	"hello/sources/pkg/mod/github.com/gorilla/mux@v1.8.0"
	"html/template"
	"log"
	"net/http"
)

func EditPage(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()
	vars := mux.Vars(r)
	Id := vars["Id"]

	row := database.QueryRow("select * from mybase.mybase where Id = ?", Id)
	prod := Product{}
	err := row.Scan(&prod.Id, &prod.Partnum, &prod.Qty, &prod.Price, &prod.Provider, &prod.Name, &prod.Remark, &prod.Date, &prod.Color)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("templates/edit.html")
		tmpl.Execute(w, prod)
	}
}
