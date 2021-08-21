package funct

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/vazman666/web/models"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
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
		if p.Color == "#fa8e47" {
			products = append(products, p)
			fmt.Println(products)
		}
	}
	sum := float32(0.0)
	for _, j := range products {
		fmt.Println(j.Price, j.Qty)
		sum = sum + j.Price*(float32)(j.Qty)
	}
	tmpl, _ := template.ParseFiles("templates/show.html")
	tmpl.Execute(w, products)
	fmt.Fprintf(w, "\nОбщая сумма: %6.2f\n", sum)
	http.Redirect(w, r, "/", 301)
}
