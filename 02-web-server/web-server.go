package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if len(name) > 4 {
		fmt.Fprintf(w, "Hello, %s! \n", name)
	} else {
		fmt.Fprintf(w, "Broken request \n")
		w.WriteHeader(400)
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
