package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	g "passgen/genutil"
)

const nn, ls, ss string = "1234567890", "abcdefghijklmnopqrstuvwxyz", "!&%$£=?ù^+*][{}-_.:,;()><"

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
	rawpassw := pick(6, nn) + pick(4, ls) + pick(2, strings.ToUpper(ls)) + pick(1, ss)
	//Stampo la password così come è
	//fmt.Println(rawpassw)

	//Mischio i caratteri e li restituisco a terminale
	fmt.Println(g.Melee(rawpassw))
	fmt.Printf("%08b\n", 5)
}

// pick: ritorna una stringa random di lunghezza L estraendola da K
func pick(L int, K string) (ret string) {
	//var c int
	for i := 1; i <= L; i++ {

		//Use utf8.RuneCountInString per prevvenire index out of range (panic)

		ret += string([]rune(K)[rand.Intn(utf8.RuneCountInString(K))])
	}
	return ret
}
