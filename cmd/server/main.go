package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	SERV_ADDR       = ":8080"
	DEFAULT_WEB_DIR = "./assets/"
)

func main() {
	var web_path string = DEFAULT_WEB_DIR
	if len(os.Args) == 2 {
		web_path = os.Args[1]
	}
	fileServer := http.FileServer(http.Dir(web_path))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello there!\n")
	})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
