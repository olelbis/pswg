package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	g "passgen/genutil"
)

func main() {
	//Inizializzo il generatore di numeri pseudo randomici
	//rand.Seed(time.Now().Unix())

	// TO DO gestione degli argomeni da command line
	l, _ := strconv.Atoi(os.Args[1])
	if l > g.LM {
		fmt.Printf("The vaule of l is: %d\n", l)
	} else {
		l = g.LM
		fmt.Printf("The vaule of l set to default size: %d\n", l)
	}

	//Genero la password raw
	rawpassw := g.Pick(6, g.NN) + g.Pick(4, g.LS) + g.Pick(2, strings.ToUpper(g.LS)) + g.Pick(1, g.SS)
	//Stampo la password così come è
	//fmt.Println(rawpassw)

	//Mischio i caratteri e li restituisco a terminale
	fmt.Println(g.Melee(rawpassw))
	//fmt.Printf("%08b\n", 5)
}
