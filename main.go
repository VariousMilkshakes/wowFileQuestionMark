package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	if len(os.Args) == 1 {
		fmt.Println("(s) Sending, (r) Receiving or (m) Manage Contacts?")
		choice, _ := reader.ReadString('\n')
		choice = CleanInput(choice)

		handleUserChoices(choice)

	} else {
		fmt.Println("File : ", os.Args[1])
	}
}

func handleUserChoices(choice string) {
	switch choice {
	case "s":
		fmt.Println("Sending File")
		break
	case "r":
		fmt.Println("Waiting to Receive File")
		break
	case "m":
		fmt.Println("~~~~~~CONTACTS~~~~~~")
		ManageContacts()
		break
	}
}

// CleanInput gets rid of all the junk with inputs
func CleanInput(raw string) (output string) {
	_ = "breakpoint"
	output = strings.TrimSpace(raw)
	return output
}
