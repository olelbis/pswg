package genutil

import (
	"math/rand"
	"time"
)

//Melee : Trasformo la stringa in array unicode(rune)
func Melee(pwdin string) string {
	rand.Seed(time.Now().Unix())
	// Trasformo la stringa in array unicode(rune)
	nRune := []rune(pwdin)
	//Uso la funzione shuffle per effetture lo swap pseudo randomico dei caratteri
	rand.Shuffle(len(nRune), func(i, j int) {
		nRune[i], nRune[j] = nRune[j], nRune[i]
	})
	return string(nRune)
}
