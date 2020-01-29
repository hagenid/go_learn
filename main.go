package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// create new MyObject

type Rec struct {
	Username string ` json:"username"`
	Password string `json:"password"`
}

// test
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", myController)
	log.Fatal(http.ListenAndServe(":3001", router))
	//	http.HandleFunc("/", myController)
	//log.Fatal(http.ListenAndServe(":3000", nil))

}

func myController(w http.ResponseWriter, r *http.Request) {
	a := "http://localhost:3001/getUser?username=Ivan"
	mux.Vars(r)
	w.Write([]byte("<h1>Hello World</h1>"))
}
