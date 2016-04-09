package goT

import "fmt"

// Cep checks and then panics on error
func Cep(e error) {
	if e != nil {
		panic(e)
	}
}

// Pbc checks and prints error, returning a bool if error has occured
func Pbc(e error) bool {
	if e != nil {
		fmt.Println("Error : ", e)
		return true
	}

	return false
}

// Rp recovers from panic (used with defer)
func Rp() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
