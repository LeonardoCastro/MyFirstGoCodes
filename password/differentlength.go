package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Longer returns a passphrase when the desired length is smaller than the phrase's length.
func Longer(p Passphrase, m int) string {
	p.Phrase = WordsForSym(p.Phrase)
	n := len(p.Phrase)
	password := p.Phrase
	switch {
	case n > m:
		password = Shortening(p, m, n)
	case n < m:
		password = Shorter(p, m, n)
	case n == m:
		password = SameLengths(p, n)
	}
	return password
}

// Shortening erases chars from the user's phrase until the desired length is obtained.
func Shortening(p Passphrase, m, n int) string {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	arrayPassword := strings.Split(p.Phrase, "")
	s1 := p.PersonalInfo1
	s2 := p.PersonalInfo2

	idx1 := strings.Index(p.Phrase, s1)
	idx2 := strings.Index(p.Phrase, s2)

	changes := 0
	for changes < n-m+1 {
		if i := rand.Intn(n); (i < idx1 || i >= idx1+len(s1)) && (i < idx2 || i >= idx2+len(s2)) {
			arrayPassword[i] = ""
			changes++
		}
	}
	fmt.Println(arrayPassword)
	p.Phrase = strings.Join(arrayPassword, "")
	password := SameLengths(p, n)

	return password
}

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

// WordsForSym changes words "And" or "Or" for symbols
func WordsForSym(password string) string {
	password = strings.Replace(password, "And", "&", 1)
	password = strings.Replace(password, "and", "&", 1)

	password = strings.Replace(password, "Or", "?", 1)
	password = strings.Replace(password, "or", "?", 1)

	return password
}
