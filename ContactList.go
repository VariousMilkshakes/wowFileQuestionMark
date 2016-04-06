package main

// ContactList holds contacts
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ContactList struct {
	directory []Contact
}

// Create builds contact list
func (cl ContactList) Create(contacts []Contact) ContactList {
	cl.directory = contacts
	return cl
}

// DisplayContacts formats contacts all pretty like
func (cl ContactList) DisplayContacts() {
	for _, contact := range cl.directory {
		fmt.Println("   Name : ", contact.Name)
		fmt.Printf("Address : %s:%s\n", contact.Address, contact.Port)
	}
}

// SaveContacts saves contacts to file
func (cl ContactList) SaveContacts() {
	output, _ := json.Marshal(cl.directory)
	err := ioutil.WriteFile("contacts.json", output, 0755)
	if err != nil {
		fmt.Println(err)
	}
}

// ContactChoices directs user choice to contact handler
func (cl ContactList) ContactChoices(choice string) {
	switch choice {
	case "a":
		fmt.Println("Add Contact")
		break
	case "r":
		fmt.Println("Remove Contact")
		break
	case "c":
		fmt.Println("Change Contact")
		break
	}
}
