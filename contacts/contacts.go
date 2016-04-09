package contacts

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/variousmilkshakes/wowFileQuestionMark/easyInput"
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
	choice = easyInput.CleanInput(choice)

	contacts.ContactChoices(choice)
}

// ContactQuery asks the user to choose a contact's name
func ContactQuery(qString string) string {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("What contact would you like to %s?\n", qString)
	choice, _ := r.ReadString('\n')
	return easyInput.CleanInput(choice)
}

// BuildContact makes a contact from user input
func BuildContact() Contact {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Name :")
	name, _ := reader.ReadString('\n')
	name = easyInput.CleanInput(name)
	fmt.Println("IP Address :")
	IP, _ := reader.ReadString('\n')
	IP = easyInput.CleanInput(IP)
	fmt.Println("Port :")
	port, _ := reader.ReadString('\n')
	port = easyInput.CleanInput(port)

	return Contact{
		Name:     name,
		Address:  IP,
		Port:     port,
		ExPhrase: "EXCHANGE",
	}
}
