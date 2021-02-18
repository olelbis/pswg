package genutil

import (
	"math/rand"
	"time"
)

//NN numeric string
const NN string = "1234567890"

//LS alphanumberic string
const LS string = "abcdefghijklmnopqrstuvwxyz"

//SS special character string
const SS string = "!&%$£=?ù^+*][{}-_.:,;()><"

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
