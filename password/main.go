package main

import (
	"fmt"
	"strconv"
)

var m = map[string]string{
	"I": "1",
	"i": "1",
	"E": "3",
	"e": "3",
	"A": "4",
	"a": "4",
	"S": "5",
	"s": "5",
	"t": "7",
	"T": "7",
	"b": "8",
	"B": "8",
	"o": "0",
	"O": "0",
}

//Passphrase type to calculate the pass phrase
type Passphrase struct {
	Phrase        string
	PersonalInfo1 string
	PersonalInfo2 string
}

func main() {
	var phrse, length, s1, s2 string
	fmt.Println("Welcome to the passphrase generator.")

	fmt.Println("Please insert your personal phrase without spaces (e.g. MyNameIsJohnAndIWasBornOn1968).")
	fmt.Scanln(&phrse)

	fmt.Println("What is the desired length for your password (number or 'same' if you want the same length as your phrase).")
	fmt.Scanln(&length)

	fmt.Println("Thank you. Now please insert your personal info present on the past phrase (e.g. John)")
	fmt.Scanln(&s1)

	fmt.Println("Great! Now insert your second personal info (e.g.1968)")
	fmt.Scanln(&s2)

	var m int
	if length == "same" {
		m = len(phrse)
	} else {
		m, _ = strconv.Atoi(length)
	}
	p := Passphrase{phrse, s1, s2}
	pssphrase := CompareLengths(p, m)
	fmt.Println("Your passphrase could be:")
	fmt.Println(pssphrase)
}

// CompareLengths chooses the right function to treat the Phrase and Personal Infos.
func CompareLengths(p Passphrase, m int) string {
	n := len(p.Phrase)
	var password string
	switch {
	case n == m:
		fmt.Println("Same lengths.")
		password = SameLengths(p, n)

	case n > m:
		fmt.Println("Phrase longer than desired.")
		password = Longer(p, m)
	case n < m:
		fmt.Println("Phrase shorter than desired.")
		password = Shorter(p, m, n)
	}
	return password
}
