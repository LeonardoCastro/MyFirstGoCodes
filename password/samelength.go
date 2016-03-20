package main

import (
	"fmt"
	"math/rand"
	"regexp"
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
	case n < m:
		password = Shorter(p, m, n)
	case n == m:
		password = SameLengths(p, n)
	case n > m:
		password = Shortening(p, m, n)
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

// FindPassword method to found safe passwords
func FindPassword(p Passphrase) string {

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	// frase que se quiere cifrar, informaciones personales
	// s1: palabra, s2: número
	password := p.Phrase
	s1 := p.PersonalInfo1
	s2 := p.PersonalInfo2

	// Cambiamos palabras por símbolos
	password = WordsForSym(password)

	// ----- Nos ocupamos de la fecha ------

	copiaS2 := strings.Split(s2, "")
	mapS2 := make(map[string]int)
	for i, s := range copiaS2 {
		mapS2[s] = i
	}

	// Quitamos 1968 de donde esté
	if strings.Contains(password, s2) {
		password = strings.Replace(password[:len(password)-len(s2)], s2, "", -1)
	}

	// Colocamos 1968 en el texto
	copiaPassword := strings.Split(password, "")
	var array []int

	for i := range copiaS2 {
		array = append(array, -(i+1)/(i+1))
	}

	for i, p := range copiaPassword {
		for j, s := range copiaS2 {
			if _, ok := m[p]; ok && s == m[p] {
				copiaPassword[i] = s
				array[mapS2[s]] = i
				copiaS2 = append(copiaS2[:j], copiaS2[j+1:]...)
			}
		}
	}

	// Colocamos los números restantes de 1968 en el orden
	// correcto
	Idx1 := []int{}
	Idx2 := []int{}

	//array = []int{-1, -1, 3, -1}
	// Obtenemos los índices en los cuales están los
	// números que sí se colocaron
	for i, a := range array {
		if i < len(array)-1 {
			if a != -1 && array[i+1] == -1 {
				Idx1 = append(Idx1, a+1)
			}
			if array[i+1] != -1 && a == -1 {
				Idx2 = append(Idx2, array[i+1])
			}
		}
	}

	// Colocamos los números que faltan

	// La primera letra de John es cambiada por un número
	idx := strings.Index(password, s1)
	copiaPassword[idx] = copiaS2[0]
	array[mapS2[copiaS2[0]]] = idx
	copiaS2 = append(copiaS2[:0], copiaS2[1:]...)

	// El número faltante se coloca en lo restante de John aleatoriamente
	set := false
	for set == false {
		i := idx + 1 + rand.Intn(Idx2[0]-idx-1)
		if ok, err := regexp.MatchString("[A-Za-z]", copiaPassword[i]); ok {
			if err != nil {
				fmt.Println(err)
			}
			copiaPassword[i] = copiaS2[0]
			array[mapS2[copiaS2[0]]] = i
			set = true
		}
	}

	// Se cambian las demás letras por números fuera del rango de 1968
	for i, str := range copiaPassword[:array[0]] {
		if s, ok := m[str]; ok {
			copiaPassword[i] = s
		}
	}

	for i, str := range copiaPassword[array[len(array)-1]+1:] {
		if s, ok := m[str]; ok {
			copiaPassword[i] = s
		}
	}

	// Se buscan tres letras seguidas para forzar cambios
	Idx3strings := []int{}
	for i := range copiaPassword[:len(copiaPassword)-3] {
		if ok, _ := regexp.MatchString("[A-Za-z]", copiaPassword[i]); ok {
			if ok, _ = regexp.MatchString("[A-Za-z]", copiaPassword[i+1]); ok {
				if ok, _ = regexp.MatchString("[A-Za-z]", copiaPassword[i+2]); ok {
					Idx3strings = append(Idx3strings, i)
				}
			}
		}
	}

	for _, i := range Idx3strings {
		if i <= array[0] || i >= array[3] {
			copiaPassword[i+1] = strconv.Itoa(rand.Intn(10))
		}
		if array[0] < i && i < array[3] {
			copiaPassword[i+1] = "%"
		}
	}

	return strings.Join(copiaPassword, "")
}

// CountingZeros function to found numbers not inserted into the password
func CountingZeros(array []int, idx1, idx2 int) int {
	count := 0
	for _, a := range array[idx1:idx2] {
		if a == 0 {
			count++
		}
	}
	return count
}

// WordsForSym changes words "And" or "Or" for symbols
func WordsForSym(password string) string {
	password = strings.Replace(password, "And", "&", 1)
	password = strings.Replace(password, "and", "&", 1)

	password = strings.Replace(password, "Or", "?", 1)
	password = strings.Replace(password, "or", "?", 1)

	return password
}
