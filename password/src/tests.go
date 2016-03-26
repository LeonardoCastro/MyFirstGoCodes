package main

// Coord is the coordinate type used.
type Coord struct {
	position  int
	character string
}

// TestLength tests the length of the passphrase
func TestLength(phrase, s1, s2 string, m int) map[int]int {
	p := Passphrase{phrase, s1, s2}
	testArray := []string{}

	for i := 0; i < 1000; i++ {
		testArray = append(testArray, CompareLengths(p, m))
	}

	length := make(map[int]int)

	for _, a := range testArray {
		l := len(a)
		if _, ok := length[l]; ok {
			length[l]++
		} else {
			length[l] = 1
		}
	}

	return length

}

//TestFrecuency tests the frecuency of each character on the passphrases
func TestFrecuency(phrase, s1, s2 string, m int) map[Coord]int {
	p := Passphrase{phrase, s1, s2}
	testArray := []string{}

	for i := 0; i < 1000; i++ {
		testArray = append(testArray, CompareLengths(p, m))
	}
	presentChars := make(map[Coord]int)

	for _, a := range testArray {
		for j, s := range a {
			c := Coord{j, string(s)}
			if _, ok := presentChars[c]; ok {
				presentChars[c]++
			} else {
				presentChars[c] = 1
			}
		}
	}
	return presentChars
}
