package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/vazman666/web/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Product struct {
	Id       int
	Partnum  string
	Qty      int
	Price    float32
	Provider string
	Name     string
	Remark   string
	Date     string
	Color    string
}

var database *sql.DB
var m sync.Mutex

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()
	vars := mux.Vars(r)
	Id := vars["Id"]

	_, err := database.Exec("delete from mybase.mybase where Id = ?", Id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

/*func EditPage(w http.ResponseWriter, r *http.Request) {
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
}*/

func EditHandler(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()
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

	_, err = database.Exec("upDate mybase.mybase set Partnum=?, Qty=?, Price=?,Provider=?,Name=?,Remark=?,Date=? , Color=? where Id=?",
		Partnum, Qty, Price, Provider, Name, Remark, Date, Color, Id)

	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", 301)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()
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
		_, err = database.Exec("insert into mybase.mybase (Partnum, Qty, Price,Provider,Name,Remark,Date,Color) values (?, ?, ?, ?, ?, ?, ?, ?)",
			Partnum, Qty, Price, Provider, Name, Remark, Date, Color)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "templates/create.html")
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	rows, err := database.Query("select * from mybase.mybase")

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	products := []Product{}

	for rows.Next() {

		p := Product{}
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
func Count(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()
	fmt.Println("conunt work")

	vars := mux.Vars(r)
	Id := vars["Id"]
	fmt.Println("Id=", Id)
	row := database.QueryRow("select * from mybase.mybase where Id = ?", Id)
	prod := Product{}
	err := row.Scan(&prod.Id, &prod.Partnum, &prod.Qty, &prod.Price, &prod.Provider, &prod.Name, &prod.Remark, &prod.Date, &prod.Color)
	if err != nil {
		fmt.Println(err)

	}
	if prod.Color == "#ffffff" {

		prod.Color = "#fa8e47"
	} else {

		prod.Color = "#ffffff"
	}

	_, err = database.Exec("upDate mybase.mybase set Color=? where Id=?", prod.Color, Id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()
	rows, err := database.Query("select * from mybase.mybase")

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	products := []Product{}

	for rows.Next() {

		p := Product{}
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

func Del(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()
	_, err := database.Query("DELETE FROM mybase.mybase WHERE color='#fa8e47';")

	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)

}

func main() {

	db, err := sql.Open("mysql", "vazman:rbhgbxb1@/mybase")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()

	s1 := mux.NewRouter()

	s1.HandleFunc("/", IndexHandler)
	s1.HandleFunc("/create", CreateHandler)
	s1.HandleFunc("/Checkout", Checkout)
	//router.HandleFunc("/add", Count)
	s1.HandleFunc("/add/{Id:[0-9]+}", Count)
	s1.HandleFunc("/del", Del)
	s1.HandleFunc("/edit/{Id:[0-9]+}", routes.EditPage).Methods("GET")
	s1.HandleFunc("/edit/{Id:[0-9]+}", EditHandler).Methods("POST")
	s1.HandleFunc("/delete/{Id:[0-9]+}", DeleteHandler)

	http.Handle("/", s1)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
