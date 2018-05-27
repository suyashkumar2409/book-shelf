package main

import (
	"net/http"
	"fmt"
	"html/template"
	"github.com/suyashkumar2409/book-shelf/templates"
	"github.com/suyashkumar2409/book-shelf/templates/templates_struct"
	"github.com/suyashkumar2409/book-shelf/db"
	"encoding/json"
)

var(
	dbHandler * db.DatabaseHandler
)

func homeHandler(w http.ResponseWriter, r * http.Request){
	t := template.Must(template.ParseFiles(templates.GetTemplatePath("index.html")))
	index := templates_struct.Index{Name:"Suyash"}
	if name := r.FormValue("name") ; name != ""{
		index.Name = name
	}
	index.DBstatus = dbHandler.IsAlive()
	if err := t.Execute(w, index) ; err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleSearch(w http.ResponseWriter, r * http.Request){
	{
		results := []templates_struct.SearchResult{
			{"Alice in Wonderland", "Lewis-Carrol", "1800", "12344"},
		}
		jsonResults, err := json.Marshal(results)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(jsonResults)
	}
}

func registerHandlers(){
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/search/", handleSearch)
}

func setupHTTPListener(){
	if err := http.ListenAndServe(":8080", nil) ; err != nil{
		fmt.Println("Error in Listen and Serve : %s" + err.Error())
	}
}

func main(){
	dbHandler = db.NewDatabaseHandler()
	if err := 	dbHandler.LoadDB(); err != nil{
		fmt.Printf("Error in Load DB : %s",err.Error())
	}
	registerHandlers()
	setupHTTPListener()
}
