package funct

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/vazman666/web/models"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		Partnum := r.FormValue("Partnum")
		Qty, _ := strconv.Atoi(r.FormValue("Qty"))
		Price, _ := strconv.ParseFloat(r.FormValue("Price"), 32)
		fmt.Println(reflect.TypeOf(Price))
		Provider := r.FormValue("Provider")
		Name := r.FormValue("Name")
		Remark := r.FormValue("Remark")
		Date := time.Now().Format("02 January 2006")
		Color := "#ffffff"
		_, err = models.Database.Exec("insert into mybase.mybase (Partnum, Qty, Price,Provider,Name,Remark,Date,Color) values (?, ?, ?, ?, ?, ?, ?, ?)",
			Partnum, Qty, Price, Provider, Name, Remark, Date, Color)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "templates/create.html")
	}
}
