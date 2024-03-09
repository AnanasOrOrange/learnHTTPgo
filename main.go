package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handlerCyber(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("index.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func mainPage(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("index.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/helloCyber", handlerCyber)
	http.HandleFunc("/guestBook", mainPage)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}
