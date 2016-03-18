package main

import (
  "time"
  "math/rand"
  "strconv"
  "strings"
  "regexp"
)

// Shorter returns a passphrase when the desired length is greater than the length of the phrase
func Shorter(p Passphrase, m, n int) string {
	password := SameLengths(p, n)

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	for i := 0; i < (m-n); i ++ {
		idx := rand.Intn(n)
		password = password[:idx]+strconv.Itoa(idx)+password[idx:]
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

	array_password := strings.Split(password, "")

	array_s1 := strings.Split(s1, "")
	array_s2 := strings.Split(s2, "")

	paso := ns2/ns1 + 1

	switch {
	case ns1 >= ns2:
		password = ProcessSameLength(array_s1, array_s2, array_password, ns2, paso, n)
	case ns2 < ns1:
		password = ProcessSameLength(array_s2, array_s1, array_password, ns1, paso, n)
	}
	return password
}

// ProcessSameLength distributes s1 and s2 randomly into the password
func ProcessSameLength(array_s1, array_s2, array_password []string, ns2, paso, n int) string {

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	for i := 0; i < ns2; i ++ {

		var j int
		if paso-1 != 1 {
			j = rand.Intn(paso+1)+1
		} else if paso -1 == 1 {
			j = 1
		}

		array_password[i] = array_s1[i]
		array_password[i+j] = array_s2[i]
	}

	array_password = No3strings(array_password, n)

	return strings.Join(array_password, "")
}

// No3strings verifies that there is no three consecutive strings.
func No3strings(array_password []string, n int) []string {

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	for i, s := range(array_password[:n-3]) {
		if ok, _ := regexp.MatchString("[A-Za-z]", s); ok {
			if ok, _ := regexp.MatchString("[A-Za-z]", array_password[i+1]); ok {
				if ok, _ := regexp.MatchString("[A-Za-z]", array_password[i+2]); ok {
					array_password[i+1] = strconv.Itoa( rand.Intn(10) )
				}
			}
		}
	}
	return array_password
}
