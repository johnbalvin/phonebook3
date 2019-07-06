package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"phonebook3/phonebook"
	"testing"
)

func TestCreateOK(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/1/phonebook/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := url.Values{}
	var name, telephone, address = "John", "56456", "afdfdf"
	q.Add("name", name)
	q.Add("address", address)
	q.Add("telephone", telephone)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", "Bearer ldfg08dsoas")

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(phonebook.Create)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var contact1 phonebook.Contact
	data, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("Not returning info from user: got %v",
			data)
	}
	if err := json.Unmarshal(data, &contact1); err != nil {
		t.Fatalf("Error in unmarshalling err: %s data: got %v",
			err, data)
	}

	req, err = http.NewRequest("GET", "http://localhost:8080/1/phonebook/read", nil)
	if err != nil {
		t.Fatal(err)
	}
	q = url.Values{}
	q.Add("id", contact1.ID)
	req.URL.RawQuery = q.Encode()

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(phonebook.Read)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	contact1 = phonebook.Contact{}
	data, err = ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("Not returning info from user: got %s",
			data)
	}
	if err := json.Unmarshal(data, &contact1); err != nil {
		t.Fatalf("Error in unmarshalling data: got %v",
			err)
	}
	if name != contact1.Name || telephone != contact1.Telephone || address != contact1.Address {
		t.Fatalf("Not returning correct ingo from user: got %v",
			contact1)
	}
}
func TestCreateFail(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/1/phonebook/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := url.Values{}
	var name, telephone, address = "John", "56456", "afdfdf"
	q.Add("name", name)
	q.Add("address", address)
	q.Add("telephone", telephone)
	req.URL.RawQuery = q.Encode()
	//	req.Header.Set("Authorization", "Bearer ldfg08dsoas")

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(phonebook.Create)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
	req.Header.Set("Authorization", "Bearer kfgjyi79")
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}
