package funct

import (
	"log"
	"net/http"

	"github.com/vazman666/web/models"
)

func Del(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	_, err := models.Database.Query("DELETE FROM mybase.mybase WHERE color='#fa8e47';")

	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)

}
