package exchange

import (
	"bufio"
	"errors"
	"fmt"
	"net"

	"github.com/variousmilkshakes/wowFileQuestionMark/contacts"
)

// PrepareSending gathers contact and file
func PrepareSending() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	target := contacts.ContactQuery("Send File To")
	cl := contacts.ReadContacts()
	peer := cl.FindContact(target)

	if peer.ExPhrase == "EXCHANGE" {
		fmt.Println("Validating Connection with Peer")
		cp := ConnectToPeer(peer)
		for {
			err := ValidatePeer(peer, cp)
			if err == nil {
				fmt.Println(cl)
				cl.SaveContacts()
				SendToPeer(cp)
			}
		}
	}

}

// ConnectToPeer establishs connection with peer
func ConnectToPeer(peer *contacts.Contact) net.Conn {
	// fullAddr := fmt.Sprintf("%s:%s", peer.Address, peer.Port)
	conn, err := net.Dial("tcp", "192.168.0.26:7654")
	if err != nil {
		fmt.Println(err)
	}

	return conn
}

// ValidatePeer checks with peer if connection is accepted
func ValidatePeer(peer *contacts.Contact, conn net.Conn) error {
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
		return nil
	}

	return errors.New("Connection Refused")
}

// SendToPeer send data to peer
func SendToPeer(conn net.Conn) {
	fmt.Fprintln(conn, "test")
	fmt.Println("Sending Data ...")
	fmt.Fprintln(conn, "cool.txt")
	fmt.Fprintln(conn, []byte("hello world!"))
	panic("Finished Sending")
}
