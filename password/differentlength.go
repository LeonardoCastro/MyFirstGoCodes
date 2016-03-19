package main

import (
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

	// Quitamos 1968 de donde esté
	if strings.Contains(password, s2) {
		password = strings.Replace(password, s2, "", -1)
	}

	arrayPassword := strings.Split(password, "")

	arrayS1 := strings.Split(s1, "")
	arrayS2 := strings.Split(s2, "")

	paso := ns2/ns1 + 1

	switch {
	case ns1 >= ns2:
		password = ProcessSameLength(arrayS1, arrayS2, arrayPassword, ns2, paso, n)
	case ns2 < ns1:
		password = ProcessSameLength(arrayS2, arrayS1, arrayPassword, ns1, paso, n)
	}
	return password
}

// ProcessSameLength distributes s1 and s2 randomly into the password
func ProcessSameLength(arrayS1, arrayS2, arrayPassword []string, ns2, paso, n int) string {

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	for i := 0; i < ns2; i++ {
		var j int
		if paso-1 != 1 {
			j = rand.Intn(paso+1) + 1
		} else if paso-1 == 1 {
			j = 1
		}

		arrayPassword[i] = arrayS1[i]
		arrayPassword[i+j] = arrayS2[i]
	}

	arrayPassword = No3strings(arrayPassword)

	return strings.Join(arrayPassword, "")
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
