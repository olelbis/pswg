package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"unicode/utf8"
)

const nn, ls, us, ss string = "1234567890", "abcdefghijklmnopqrstuvwxyz", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "!&%$£=?ù^+*][{}-_.:,;()><"

const lm int = 12

func main() {
	//Inizializzo il generatore di numeri pseudo randomici
	rand.Seed(time.Now().Unix())

	// TO DO gestione degli argomeni da command line
	l, _ := strconv.Atoi(os.Args[1])
	if l > lm {
		fmt.Printf("The vaule of l is: %d\n", l)
	} else {
		l = lm
		fmt.Printf("The vaule of l set to default size: %d\n", l)
	}

	//Genero la password raw
	rawpassw := pick(6, nn) + pick(4, ls) + pick(2, us) + pick(1, ss)
	//Stampo la password così come è
	//fmt.Println(rawpassw)

	//Mischio i caratteri e li restituisco a terminale
	fmt.Println(melee(rawpassw))
}

// pick: ritorna una stringa random di lunghezza L estraendola da K
func pick(L int, K string) (ret string) {
	//var c int
	for i := 1; i <= L; i++ {

		//Use utf8.RuneCountInString per prevvenire indexofbound (panic)

		ret += string([]rune(K)[rand.Intn(utf8.RuneCountInString(K))])
	}
	return ret
}

func melee(pwdin string) string {
	// Trasformo la stringa in array unicode(rune)
	nRune := []rune(pwdin)
	//Uso la funzione shuffle per effetture lo swap pseudo randomico dei caratteri
	rand.Shuffle(len(nRune), func(i, j int) {
		nRune[i], nRune[j] = nRune[j], nRune[i]
	})
	return string(nRune)
}
