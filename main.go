package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Site Web en Go</h1><p>Page d'accueil</p>")
	})

	http.HandleFunc("/url", URLHandler)
	http.HandleFunc("/color", ColorHandler)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func URLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>URL visitée</h1><p>Vous avez visité : %s</p>", r.URL.Path)
}

func ColorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Colors</h1><style>*{background-color: #006400;}</style>")
}
