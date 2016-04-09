package exchange

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/variousmilkshakes/wowFileQuestionMark/easyInput"
)

const (
	defaultPort     = 3827
	defaultExchange = "default"
)

// StartListening begins waiting for files
func StartListening(settings map[string]string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println(settings["Port"])
	ln, err := net.Listen("tcp", ":7654")

	if err != nil {
		fmt.Println("Cannot start server")
		fmt.Println(err)
	}

	for {
		connection, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go ValidateConnection(connection)
	}
}

// ValidateConnection between clients
func ValidateConnection(connection net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println("Phrase")
	exchange, _ := bufio.NewReader(connection).ReadString('\n')
	exchangeString := cleanInput(string(exchange))
	fmt.Println(exchangeString)

	if "test" == exchangeString {
		fmt.Println("Paired")

		fmt.Fprintln(connection, "ok")
		receiveData(connection)
	} else {
		newClient(connection)
	}
}

func receiveData(connection net.Conn) {
	fmt.Println("Waiting for Data...")

	com := bufio.NewReader(connection)

	fileName, _ := com.ReadString('\n')
	fileName = cleanInput(fileName)
	fmt.Println(string(fileName))

	fileData, _ := com.ReadString('\n')
	// fmt.Println(fileData)

	newFile := easyInput.File{
		[]byte(fileData),
		fileName,
		"/",
		fileData,
	}

	newFile.WriteFile()

	fmt.Fprintln(connection, "done")
	panic("Finished Connection")
}

func newClient(connection net.Conn) {
	fmt.Println("Do you want to pair")
	r := bufio.NewReader(os.Stdin)
	response, _ := r.ReadString('\n')
	response = cleanInput(response)
	if response == "y" {
		fmt.Println("Pairing")
		test := fmt.Sprintln(Settings["Exchange_String"])
		fmt.Println(test)
		// connection.Write([]byte("test"))
		fmt.Fprintln(connection, "ok")
		fmt.Fprintln(connection, "test")
		fmt.Println("Paired")
		ValidateConnection(connection)
	} else {
		fmt.Println("Connection Rejected")
	}
}

func cleanInput(raw string) (output string) {
	_ = "breakpoint"
	output = strings.TrimSpace(raw)
	return output
}
