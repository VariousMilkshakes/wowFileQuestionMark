package exchange

import (
	"bufio"
	"errors"
	"fmt"
	"net"

	"github.com/variousmilkshakes/wowFileQuestionMark/contacts"
	"github.com/variousmilkshakes/wowFileQuestionMark/easyInput"
	"github.com/variousmilkshakes/wowFileQuestionMark/goT"
)

// PrepareSending gathers contact and file
func PrepareSending(settings map[string]string) {

	// Catch send finish
	defer goT.Rp()

	file, err := easyInput.OpenFile(settings["filePath"])
	goT.Cep(err)

	// Ask for target contact
	target := contacts.ContactQuery("Send File To")

	// Get correct contact from all contacts
	cl := contacts.ReadContacts()
	peer := cl.FindContact(target)

	// Are you paired with peer
	if peer.ExPhrase == "EXCHANGE" {

		fmt.Println("Validating Connection with Peer")

		cp := ConnectToPeer(peer)
		for {

			// Start pairing process
			p, err := ValidatePeer(peer, cp)

			// If pair completes
			if err == nil {

				// Save exchange phrase to cl
				cl.UpdateContact(p)
				cl.SaveContacts()

				// Start sending file to peer
				SendToPeer(cp, file)

			}

		}

	} else {
		fmt.Println("Skipping Validation")

		// Already paired with peer
		cp := ConnectToPeer(peer)

		for {
			SendToPeer(cp, file)
		}
	}

}

// ConnectToPeer establishs connection with peer
func ConnectToPeer(peer *contacts.Contact) net.Conn {
	fullAddr := fmt.Sprintf("%s:%s", peer.Address, peer.Port)
	fmt.Println(fullAddr)
	conn, err := net.Dial("tcp", fullAddr)
	if err != nil {
		fmt.Println(err)
	}

	return conn
}

// ValidatePeer checks with peer if connection is accepted
func ValidatePeer(peer *contacts.Contact, conn net.Conn) (*contacts.Contact, error) {
	com := bufio.NewReader(conn)

	fmt.Fprintln(conn, "pair")
	message, _ := com.ReadString('\n')
	message = cleanInput(message)
	fmt.Println(message)

	if message == "ok" {
		phrase, _ := com.ReadString('\n')
		phrase = cleanInput(phrase)
		fmt.Println(phrase)
		fmt.Println("PAIRED")
		peer.ExPhrase = phrase
		return peer, nil
	}

	return peer, errors.New("Connection Refused")
}

// SendToPeer send data to peer
func SendToPeer(conn net.Conn, file easyInput.File) {
	fmt.Fprintln(conn, "test")

	com := bufio.NewReader(conn)
	check, err := com.ReadString('\n')
	if easyInput.CleanInput(check) != "ok" || err != nil {
		panic("Connection Refused")
	}

	fmt.Println("Sending Data ...")
	fmt.Fprintln(conn, file.FileName)
	fmt.Println(file.PlainText)
	fmt.Fprintln(conn, file.PlainText)
	panic("Finished Sending")
}
