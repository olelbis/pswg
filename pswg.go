package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	g "passgen/genutil"
)

const (
	execname string = "pswg"
	version  string = "0.1"
)

func main() {

	// TO DO argument managment inside or outside main package
	l, _ := strconv.Atoi(os.Args[1])
	if l > g.LM {
		fmt.Printf("The vaule of l is: %d\n", l)
	} else {
		l = g.LM
		fmt.Printf("The vaule of l set to default size: %d\n", l)
	}

	//Generate raw password probably i need to refactor Pick function?
	rawpassw := g.Pick(6, g.NN) + g.Pick(4, g.LS) + g.Pick(2, strings.ToUpper(g.LS)) + g.Pick(1, g.SS)
	//For debug purpose in this alpha state
	//fmt.Println(rawpassw)

	//Melee... Nice name no?
	fmt.Println(g.Melee(rawpassw))
	//fmt.Printf("%08b\n", 5)
}
