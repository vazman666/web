package funct

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/vazman666/web/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()

	rows, err := models.Database.Query("select * from mybase.mybase")

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	products := []models.Product{}

	for rows.Next() {

		p := models.Product{}
		err := rows.Scan(&p.Id, &p.Partnum, &p.Qty, &p.Price, &p.Provider, &p.Name, &p.Remark, &p.Date, &p.Color)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)

	}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, products)

}
