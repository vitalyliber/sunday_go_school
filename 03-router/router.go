package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func HelloPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := r.FormValue("name")
	LengthValidation(name, w)
}

func HelloGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	LengthValidation(name, w)
}

func HelloGetSecond(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := r.FormValue("name")
	LengthValidation(name, w)
}

func LengthValidation(name string, w http.ResponseWriter) {
	if len(name) > 2 {
		fmt.Fprintf(w, "Hello, %s! \n", name)
	} else {
		fmt.Fprintf(w, "Broken request \n")
		w.WriteHeader(400)
	}
}

func main() {
	router := httprouter.New()
	router.POST("/hello", HelloPost)
	router.GET("/hello/:name", HelloGet)
	router.GET("/hello", HelloGetSecond)

	log.Fatal(http.ListenAndServe(":8000", router))
}
