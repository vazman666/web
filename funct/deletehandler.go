package funct

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vazman666/web/models"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	vars := mux.Vars(r)
	Id := vars["Id"]

	_, err := models.Database.Exec("delete from mybase.mybase where Id = ?", Id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}
