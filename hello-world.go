package main

import (
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

// Data for HTML templates
// type MyData struct {
// 	Title string

// 	Homelinktxt string
// 	Homelink    string

// 	Hellolinktxt string
// 	Hellolink    string

// 	Headerlinktxt string
// 	Headerlink    string
// }

func hello(w http.ResponseWriter, req *http.Request) {

	// Create data - should centralize
	// helloData := MyData{
	// 	Title:         "Sysdig Hello World - Hello",
	// 	Homelinktxt:   "Home",
	// 	Homelink:      "/",
	// 	Hellolinktxt:  "Hello",
	// 	Hellolink:     "/hello",
	// 	Headerlinktxt: "Header",
	// 	Headerlink:    "/headers",
	// }

	// Title := "Sysdig Hello World - Home"
	// // Output html header
	// parsedTemplate, _ := template.ParseFiles("template-header.html")
	// err := parsedTemplate.Execute(w, Title)
	// if err != nil {
	// 	log.Println("Error executing template :", err)
	// 	return
	// }

	// data in the page
	fmt.Fprintf(w, "Hi there!\n")

	// Output html footer
	// parsedTemplate2, _ := template.ParseFiles("template-footer.html")
	// err2 := parsedTemplate2.Execute(w, nil)
	// if err != nil {
	// 	log.Println("Error executing template :", err2)
	// 	return
	// }
}

func defaultPage(w http.ResponseWriter, req *http.Request) {

	unescaped := `<h1>Hi there! H1</h1>`
	escaped := html.EscapeString(unescaped)
	fmt.Fprintf(w, "%s\n", escaped)

	// Create data - should centralize
	// homeData := MyData{
	// 	Title:         "Sysdig Hello World - Home",
	// 	Homelinktxt:   "Home",
	// 	Homelink:      "/",
	// 	Hellolinktxt:  "Hello",
	// 	Hellolink:     "/hello",
	// 	Headerlinktxt: "Header",
	// 	Headerlink:    "/headers",
	// }

	// Output html for home

	// Title := "Sysdig Hello World - Home"
	// parsedTemplate, _ := template.ParseFiles("home-template.html")
	// err := parsedTemplate.Execute(w, Title)
	// if err != nil {
	// 	log.Println("Error executing template :", err)
	// 	return
	// }
}

func headers(w http.ResponseWriter, req *http.Request) {

	// Create data - should centralize
	// headerData := MyData{
	// 	Title:         "Sysdig Hello World - Headers",
	// 	Homelinktxt:   "Home",
	// 	Homelink:      "/",
	// 	Hellolinktxt:  "Hello",
	// 	Hellolink:     "/hello",
	// 	Headerlinktxt: "Header",
	// 	Headerlink:    "/headers",
	// }

	// Output html header
	// parsedTemplate, _ := template.ParseFiles("template-header.html")
	// err := parsedTemplate.Execute(w, headerData)
	// if err != nil {
	// 	log.Println("Error executing template :", err)
	// 	return
	// }

	// data in the page - output html headers
	for name, headers := range req.Header {
		for _, h := range headers {

			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

	// Output html footer
	// parsedTemplate2, _ := template.ParseFiles("template-footer.html")
	// err2 := parsedTemplate2.Execute(w, nil)
	// if err != nil {
	// 	log.Println("Error executing template :", err2)
	// 	return
	// }
}

func main() {

	// Setup Logging
	var log = logrus.New()
	log.Out = os.Stdout

	// Needed to serve static assets
	fs := http.FileServer(http.Dir("./img"))
	http.Handle("/img/", http.StripPrefix("/img/", fs))

	// Menu Items
	http.HandleFunc("/", defaultPage)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Serve it on a platter
	http.ListenAndServe(":8090", nil)
}
