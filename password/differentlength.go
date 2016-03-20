package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Shorter returns a passphrase when the desired length is greater than the length of the phrase
func Shorter(p Passphrase, m, n int) string {
	password := SameLengths(p, n)

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	for i := 0; i < (m - n); i++ {
		idx := rand.Intn(n)
		password = password[:idx] + strconv.Itoa(idx) + password[idx:]
	}
	return password
}

// SameLengths returns the passphrase when the length of the phrase is the same as the desired length for the passphrase
func SameLengths(p Passphrase, n int) string {
	password := p.Phrase
	s1 := p.PersonalInfo1
	s2 := p.PersonalInfo2

	ns1 := len(s1)
	ns2 := len(s2)

	// Quitamos s2 de donde esté
	if strings.Contains(password, s2) {
		password = strings.Replace(password, s2, "", -1)
	}

	// Quitamos s2 de donde esté
	if strings.Contains(password, s1) {
		password = strings.Replace(password, s1, "", -1)
	}

	// arrayPassword := strings.Split(password, "")
	fmt.Println(password)
	arrayS1 := strings.Split(s1, "")
	arrayS2 := strings.Split(s2, "")

	paso := ns2/ns1 + 1

	switch {
	case ns1 >= ns2:
		password = ProcessSameLength(arrayS1, arrayS2, password, ns2, ns1, paso)
	case ns2 > ns1:
		password = ProcessSameLength(arrayS2, arrayS1, password, ns2, ns1, paso)
	}
	return password
}

// ProcessSameLength distributes s1 and s2 randomly into the password
func ProcessSameLength(arrayS1, arrayS2 []string, password string, ns2, ns1, paso int) string {

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	var cuenta1, cuenta2, L int

	for L < ns1+ns2 {

		s := rand.Intn(2)

		if s == 0 {
			password, cuenta1 = Fill(password, arrayS1, cuenta1)
		} else {
			password, cuenta2 = Fill(password, arrayS2, cuenta2)
		}
		L = cuenta1 + cuenta2
	}

	arrayPassword := No3strings(strings.Split(password, ""))

	return strings.Join(arrayPassword, "")
}

// Fill fills password with personal info
func Fill(password string, array []string, cuenta int) (string, int) {

	if cuenta < len(array) {
		j := rand.Intn(len(password) + 1)
		password = password[:j] + array[cuenta] + password[j:]
		cuenta++
	}
	return password, cuenta
}

// No3strings verifies that there is no three consecutive strings.
func No3strings(arrayPassword []string) []string {

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	for i, s := range arrayPassword[:len(arrayPassword)-3] {
		if ok, _ := regexp.MatchString("[A-Za-z]", s); ok {
			if ok, _ := regexp.MatchString("[A-Za-z]", arrayPassword[i+1]); ok {
				if ok, _ := regexp.MatchString("[A-Za-z]", arrayPassword[i+2]); ok {
					arrayPassword[i+1] = strconv.Itoa(rand.Intn(10))
				}
			}
		}
	}
	return arrayPassword
}
