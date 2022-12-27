package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Estructuras
type Usuario struct {
	UserName string
	Edad     int
}

// var templates = template.Must(template.New("T").ParseGlob("templates/*.html"))
var templates = template.Must(template.New("T").ParseGlob("../templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("../templates/error/error.html"))

// Funcion Handler Error

func HandlerError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

// Funcion para renderizar templates
func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		HandlerError(rw, http.StatusInternalServerError)
	}
}

// Handlers
func Index(rw http.ResponseWriter, r *http.Request) {
	usuario := Usuario{"Alex", 26}
	renderTemplate(rw, "index.html", usuario)
}

func Registro(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "registro.html", nil)
}

// Función principal
func main() {
	// Archivos Estaticos
	staticFile := http.FileServer(http.Dir("static"))

	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	// mux static File
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))

	//Server
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor esta corriendo en puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	log.Fatal(server.ListenAndServe())
}
