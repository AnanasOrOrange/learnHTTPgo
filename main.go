package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}
func handlerCyber(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("index.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func mainPage(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("mainPage.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}
func addSignature(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestBook", http.StatusFound)
}
func main() {
	// html + css
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/helloCyber", handlerCyber)

	// example from book
	http.HandleFunc("/guestBook", mainPage)
	http.HandleFunc("/guestBook/new", addSignature)
	http.HandleFunc("/guestbook/create", createHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}
