package phonebook

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

func addContact(name, address, telephone string) (Contact, error) {
	now := time.Now().Unix()
	ctx := context.Background()
	doc := clienteFS.Collection("Contacts").NewDoc()
	contact := Contact{Name: name, Address: address, Telephone: telephone, CreationDate: now, LastModification: now}
	contact.ID = doc.ID
	if _, err := doc.Create(ctx, contact); err != nil {
		log.Println("phonebook -> Add:1 -> err:", err)
		return contact, err
	}
	return contact, nil
}

func getContacts(size int) ([]Contact, error) {
	ctx := context.Background()
	snapShops, err := clienteFS.Collection("Contacts").Limit(size).Documents(ctx).GetAll()
	if err != nil {
		log.Println("phonebook -> GetContacts:1 -> err:", err)
		return nil, err
	}
	var contacts []Contact
	for _, snap := range snapShops {
		var contact Contact
		if err = snap.DataTo(&contact); err != nil {
			log.Println("phonebook -> GetContacts:2 -> err:", err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func getContactByID(id string) (Contact, bool, error) {
	ctx := context.Background()
	snapShops, err := clienteFS.Collection("Contacts").Where("ID", "==", id).Documents(ctx).GetAll()
	if err != nil {
		log.Println("phonebook -> getContactByID:1 -> err:", err)
		return Contact{}, false, err
	}
	if len(snapShops) != 1 {
		return Contact{}, false, nil
	}
	var contact Contact
	if err = snapShops[0].DataTo(&contact); err != nil {
		log.Println("phonebook -> getContactByID:2 -> err:", err)
		return contact, false, err
	}
	return contact, true, err
}
func updateName(id, name string) error {
	doc := clienteFS.Collection("Contacts").Doc(id)
	ref := []firestore.Update{{Path: "Name", Value: name}}
	ctx := context.Background()
	if _, err := doc.Update(ctx, ref); err != nil {
		log.Println("phonebook -> updateName:1 -> err:", err)
		return err
	}
	return nil
}
func updateTelephone(id, telephone string) error {
	doc := clienteFS.Collection("Contacts").Doc(id)
	ref := []firestore.Update{{Path: "Telephone", Value: telephone}}
	ctx := context.Background()
	if _, err := doc.Update(ctx, ref); err != nil {
		log.Println("phonebook -> updateTelephone:1 -> err:", err)
		return err
	}
	return nil
}

func updateAddress(id, address string) error {
	doc := clienteFS.Collection("Contacts").Doc(id)
	ref := []firestore.Update{{Path: "Address", Value: address}}
	ctx := context.Background()
	if _, err := doc.Update(ctx, ref); err != nil {
		log.Println("phonebook -> updateAddress:1 -> err:", err)
		return err
	}
	return nil
}

func deleteContact(id string) error {
	doc := clienteFS.Collection("Contacts").Doc(id)
	ctx := context.Background()
	if _, err := doc.Delete(ctx); err != nil {
		log.Println("phonebook -> deleteContact:1 -> err:", err)
		return err
	}
	return nil
}
