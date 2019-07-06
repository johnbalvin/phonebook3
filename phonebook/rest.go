package phonebook

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

//this is a defaults token(obviusly in real world scenario it may be different for each user and for each permission)

var tokenDelete = "kdsfpojui43"
var tokenUpdate = "kfgjyi79"
var tokenCreate = "ldfg08dsoas"

//Read is the function to handle read API
func Read(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		contacts, err := getContacts(100)
		if err != nil {
			log.Println("read:1 -> err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(contacts); err != nil {
			log.Println("read:2 -> err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	contact, exists, err := getContactByID(id)
	if err != nil {
		log.Println("read:3 -> err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !exists {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err := json.NewEncoder(w).Encode(contact); err != nil {
		log.Println("read:4 -> err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//Create is the function to handle create API
func Create(w http.ResponseWriter, r *http.Request) {
	autorization := r.Header.Get("Authorization")
	splitToken := strings.Split(autorization, "Bearer ")
	if len(splitToken) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	reqToken := splitToken[1]
	if reqToken != tokenCreate {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	telephone := r.FormValue("telephone")
	contact, err := addContact(name, address, telephone)
	if err != nil {
		log.Println("create:1 -> err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(contact); err != nil {
		log.Println("create:2 -> err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//Delete is the function to handle Delete API
func Delete(w http.ResponseWriter, r *http.Request) {
	autorization := r.Header.Get("Authorization")
	splitToken := strings.Split(autorization, "Bearer ")
	if len(splitToken) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	reqToken := splitToken[1]
	if reqToken != tokenCreate {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id := r.FormValue("id")
	if err := deleteContact(id); err != nil {
		log.Println("delete -> err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//Update is the function to handle update API
func Update(w http.ResponseWriter, r *http.Request) {
	autorization := r.Header.Get("Authorization")
	splitToken := strings.Split(autorization, "Bearer ")
	if len(splitToken) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	reqToken := splitToken[1]
	if reqToken != tokenDelete {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	what := r.FormValue("field")
	id := r.FormValue("id")
	value := r.FormValue("value")
	switch what {
	case "telephone":
		if err := updateTelephone(id, value); err != nil {
			log.Println("update:1 -> err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "name":
		if err := updateName(id, value); err != nil {
			log.Println("update:2 -> err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "address":
		if err := updateAddress(id, value); err != nil {
			log.Println("update:3 -> err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
	}
}
