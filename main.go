package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	name string
	age  int
	job  string
}

// type contact struct {
// 	Name    string
// 	Email   string
// 	Address string
// }

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to  the go server")
}
func firstHandler(w http.ResponseWriter, r *http.Request) {

	p := user{
		name: "Rahul",
		age:  30,
		job:  "Analyst",
	}
	fmt.Fprintf(w, "the name of the person %s and their age is %d and their Job is %s \n", p.name, p.age, p.job)
}
func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseform error : %v", err)
		return
	}

	// if r.Method != http.MethodPost {
	// 	fmt.Fprint(w, "Something went wrong", http.StatusBadRequest)
	// 	return
	// }
	fmt.Fprintln(w, "Post request successfull")

	Name := r.FormValue("name")
	Email := r.FormValue("email")
	Address := r.FormValue("address")

	fmt.Fprintf(w, " Name of the person:  %s\n", Name)
	fmt.Fprintf(w, "Email of the person : %s\n", Email)
	fmt.Fprintf(w, "Address of the person: %s\n", Address)

}

func main() {
	r := mux.NewRouter()
	fileserver := http.FileServer(http.Dir("static"))

	r.Handle("/form.html", fileserver)
	// r.HandleFunc("/", helloHandler).Methods("GET")
	r.HandleFunc("/", firstHandler).Methods("GET")
	r.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}

}
