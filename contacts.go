package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Contact to send files to
type Contact struct {
	Name     string `json:"name"`
	ExPhrase string `json:"exchangePhrase"`
	Address  string `json:"ipAddress"`
}

// ReadContacts reads contact files
func ReadContacts() []Contact {
	content, err := ioutil.ReadFile("contacts.json")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	var contacts []Contact
	json.Unmarshal(content, &contacts)
	fmt.Println(contacts)

	return contacts
}

// SaveContacts saves contacts to file
func SaveContacts(contactList []Contact) {
	contactList[0].Name = "simon"

	output, _ := json.Marshal(contactList)
	err := ioutil.WriteFile("contacts.json", output, 0755)
	if err != nil {
		fmt.Println(err)
	}

}
