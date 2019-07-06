package main

import (
	"log"
	"net/http"
	"phonebook3/phonebook"
)

func main() {
	handleFuncs()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func handleFuncs() {
	http.HandleFunc("/1/phonebook/create", phonebook.Create)
	http.HandleFunc("/1/phonebook/update", phonebook.Update)
	http.HandleFunc("/1/phonebook/delete", phonebook.Delete)
	http.HandleFunc("/1/phonebook/read", phonebook.Read)
}
