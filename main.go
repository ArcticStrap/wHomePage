package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Running!")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/*", http.StripPrefix("/static/", fs))

	mainT, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatalf("%s\n", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		mainT.Execute(w, nil)
	})

	if os.Getenv("DEV") != "" {
		log.Fatal(http.ListenAndServeTLS("127.0.0.1:6969", "certs/cert.pem", "certs/key.pem", nil))
	} else {
		log.Fatal(http.ListenAndServe("127.0.0.1:6969", nil))
	}
}
