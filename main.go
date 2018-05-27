package main

import (
	"net/http"
	"fmt"
	"html/template"
	"github.com/suyashkumar2409/book-shelf/templates"
	"github.com/suyashkumar2409/book-shelf/templates/templates_struct"
)

func homeHandler(w http.ResponseWriter, r * http.Request){
	t := template.Must(template.ParseFiles(templates.GetTemplatePath("index.html")))
	index := templates_struct.Index{Name:"Suyash"}
	if name := r.FormValue("name") ; name != ""{
		index.Name = name
	}
	if err := t.Execute(w, index) ; err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func registerHandlers(){
	http.HandleFunc("/", homeHandler)
}

func main(){
	registerHandlers()
	if err := http.ListenAndServe(":8080", nil) ; err != nil{
		fmt.Println(err)
	}
}
