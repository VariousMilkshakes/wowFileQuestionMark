package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/variousmilkshakes/wowFileQuestionMark/contacts"
	"github.com/variousmilkshakes/wowFileQuestionMark/exchange"
)

func main() {
	if len(os.Args) == 1 {
		Start()
	} else {
		fmt.Println("File : ", os.Args[1])
	}

	fmt.Println("Got ot")
}

// Start is the begin of the control tree
func Start() {
	for {
		reader := bufio.NewReader(os.Stdin)

		exchange.ReadPhrase()
		fmt.Println(exchange.Settings["Exchange_String"])

		fmt.Println("(s) Sending, (r) Receiving or (m) Manage Contacts?")
		choice, _ := reader.ReadString('\n')
		choice = cleanInput(choice)

		handleUserChoices(choice)
	}
}

func handleUserChoices(choice string) {
	switch choice {
	case "s":
		fmt.Println("Sending File")
		exchange.PrepareSending()
		break
	case "r":
		fmt.Println("Waiting to Receive File")
		exchange.StartListening()
		break
	case "m":
		fmt.Println("~~~~~~CONTACTS~~~~~~")
		contacts.ManageContacts()
		break
	}
}

func cleanInput(raw string) (output string) {
	_ = "breakpoint"
	output = strings.TrimSpace(raw)
	return output
}
