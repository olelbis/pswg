package genutil

import (
	"math/rand"
	"time"
	"unicode/utf8"
)

const (
	//NN numeric string
	NN string = "1234567890"

	//LS alphanumeric string
	LS string = "abcdefghijklmnopqrstuvwxyz"

	//SS special character string
	SS string = "!&%$£=?^+*][{}-_.:,;()><"

	//LM Minimum Lenght
	LM int = 12
)

//Melee : Get a string in input e do some shuffle
func Melee(pwdin string) string {
	rand.Seed(time.Now().Unix())
	// Transform string to rune
	nRune := []rune(pwdin)
	//Use rand.Shuffle function to make a pseudo randomic char swap
	rand.Shuffle(len(nRune), func(i, j int) {
		nRune[i], nRune[j] = nRune[j], nRune[i]
	})
	return string(nRune)
}

// Pick : return random string of lenght L extract it form  K
func Pick(L int, K string) (ret string) {
	rand.Seed(time.Now().Unix())

	// yet another "i" loop
	for i := 1; i <= L; i++ {

		//Use utf8.RuneCountInString to prevent index out of range (panic)
		//in case of multibyte character
		ret += string([]rune(K)[rand.Intn(utf8.RuneCountInString(K))])
	}
	return ret
}
