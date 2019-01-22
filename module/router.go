package module

import (
    "net/http"
	"html/template"
)

func authHandler(c http.ResponseWriter, r *http.Request) {
	var templ = template.Must(template.ParseFiles("templates/index.html"))
	data := struct {
		Host       string
	}{r.Host}
	templ.Execute(c, data)
}

func gameHandler(c http.ResponseWriter, r *http.Request) {
	var templ = template.Must(template.ParseFiles("templates/game.html"))
	data := struct {
		Host       string
	}{r.Host}
	templ.Execute(c, data)
}

func Router() {
	//log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
	fs := http.FileServer(http.Dir("./templates/"))
	http.Handle("/", fs)
	http.Handle("/start/", fs)
	http.HandleFunc("/cor/", authHandler)
	http.HandleFunc("/game/", gameHandler)
}