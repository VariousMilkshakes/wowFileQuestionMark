package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Contact to send files to
type Contact struct {
	Name     string `json:"name"`
	ExPhrase string `json:"exchangePhrase"`
	Address  string `json:"ipAddress"`
	Port     string `json:"port"`
}

// ReadContacts reads contact files
func ReadContacts() ContactList {
	content, err := ioutil.ReadFile("contacts.json")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	var contacts []Contact
	json.Unmarshal(content, &contacts)

	var newList ContactList
	newList = newList.Create(contacts)

	return newList
}

// ManageContacts handles contacts
func ManageContacts() {
	contacts := ReadContacts()
	contacts.DisplayContacts()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("(a) Add, (r) Remove or (c) Change Contacts?")
	choice, _ := reader.ReadString('\n')
	choice = CleanInput(choice)

	contacts.ContactChoices(choice)
}
