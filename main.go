package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/variousmilkshakes/wowFileQuestionMark/contacts"
	"github.com/variousmilkshakes/wowFileQuestionMark/easyInput"
	"github.com/variousmilkshakes/wowFileQuestionMark/exchange"
)

func main() {

	// Has a path to a file been provided
	if len(os.Args) == 1 {
		// Control tree
		Start()
	} else {
		fmt.Println("File : ", os.Args[1])
		// Go straight to ask which contact
	}

	fmt.Println("Got ot")

}

// Start is the begin of the control tree
func Start() {

	for {
		reader := bufio.NewReader(os.Stdin)

		// Read settings from file
		settings := exchange.ReadPhrase()
		fmt.Println(settings["Exchange_String"])

		fmt.Println("(s) Sending, (r) Receiving or (m) Manage Contacts?")
		choice, _ := reader.ReadString('\n')
		choice = easyInput.CleanInput(choice)

		handleUserChoices(choice, settings)
	}

}

func handleUserChoices(choice string, settings map[string]string) {

	switch choice {
	case "s":
		fmt.Println("Sending File")
		fmt.Println("What file?")
		filePath, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		settings["filePath"] = easyInput.CleanInput(filePath)
		exchange.PrepareSending(settings)
		break

	case "r":
		fmt.Println("Waiting to Receive File")
		exchange.StartListening()
		break

	case "m":
		fmt.Println("~~~~~~ CONTACTS ~~~~~~")
		contacts.ManageContacts()
		break

	}

}
