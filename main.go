package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 0 {
		fmt.Println("(s) Sending, (r) Receiving or (m) Manage Contacts?")
	} else {
		contacts := ReadContacts()
		SaveContacts(contacts)
	}

}

// func sendFile(path string) err {
// 	conn, err := net.Dial("")
// }
