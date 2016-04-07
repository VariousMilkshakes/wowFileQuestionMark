package contacts

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// ContactList holds contacts
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
		fmt.Println("   Name :", contact.Name)
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
		newContact := BuildContact()
		cl.AddContact(newContact)
		break
	case "r":
		fmt.Println("Remove Contact")
		target := ContactQuery("Remove")
		cl.RemoveContact(target)
		break
	case "c":
		fmt.Println("Change Contact")
		target := ContactQuery("Modify")
		cl.ModifyContact(target)
		break
	}
}

// AddContact appends new contact onto current directory
func (cl ContactList) AddContact(newContact Contact) {
	defer func() {
		if r := recover(); r != nil {

			dir := cl.directory
			length := len(dir)

			if length == cap(dir) {
				// Directory is full, needs to expand
				newDir := make([]Contact, len(dir), 2*len(dir)+1)
				copy(newDir, dir)
				dir = newDir
			}

			dir = dir[0 : length+1]
			dir[length] = newContact

			cl.directory = dir

			fmt.Println("Contact Added")
		} else {
			fmt.Println("Contact Already Exists")
		}
	}()

	cl.FindContact(newContact.Name)
}

// RemoveContact removes contact with provided name
func (cl ContactList) RemoveContact(targetName string) error {
	for indx, contact := range cl.directory {
		if contact.Name == targetName {
			lower := cl.directory[:indx-1]
			upper := cl.directory[indx+1:]

			cl.directory = append(lower, upper...)
			return nil
		}
	}

	return errors.New("Cannot Find Contact!")
}

// FindContact returns contact with the same name
func (cl ContactList) FindContact(targetName string) *Contact {
	for _, contact := range cl.directory {
		if contact.Name == targetName {
			return &contact
		}
	}

	panic("No Contact Found!")
}

// ModifyContact changes details of a contact
func (cl ContactList) ModifyContact(targetName string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Cannot Find Contact!")
		}
	}()

	target := cl.FindContact(targetName)

	newName, err := modQuestion("Name")
	if err == nil {
		target.Name = newName
	}

	newIP, err := modQuestion("IP")
	if err == nil {
		target.Address = newIP
	}

	newPort, err := modQuestion("Port")
	if err == nil {
		target.Port = newPort
	}

	cl.DisplayContacts()
	cl.SaveContacts()
}

func modQuestion(catagory string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Change %s (leave blank to not change): \n", catagory)

	input, _ := reader.ReadString('\n')
	if input = cleanInput(input); input != "" {
		return input, nil
	}

	return "", errors.New("No Change")
}
