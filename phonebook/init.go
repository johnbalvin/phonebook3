package phonebook

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

var clienteFS *firestore.Client

//Contact contains info from a contact
type Contact struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Telephone        string `json:"phone_number"`
	CreationDate     int64  `json:"-"`
	LastModification int64  `json:"-"`
}

var projectID = "johnbalvin" //this my project, you should assing you own
func init() {
	var err error
	ctx := context.Background()
	clienteFS, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal("clients -> Firestore -> err:", err)
	}
}
