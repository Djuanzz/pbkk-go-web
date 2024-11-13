package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Djuanzz/pbkk-go-web/webapp"
)

func main() {
	fmt.Println("Adnan Abdullah Juan | 5025221155")
	fmt.Println("Golang Web Application")

	p1 := &webapp.Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.Save()
	p2, _ := webapp.LoadPage("TestPage")
	fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", webapp.MakeHandler(webapp.ViewHandler))
	http.HandleFunc("/edit/", webapp.MakeHandler(webapp.EditHandler))
	http.HandleFunc("/save/", webapp.MakeHandler(webapp.SaveHandler))

	log.Fatal(http.ListenAndServe(":5000", nil))
}
