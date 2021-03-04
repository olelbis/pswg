package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	g "passgen/genutil"
)

const (
	exec     string = "pswg"
	ver      string = "0.2"
	defaults string = `No option specified or wrong number of arguments, use the defaults for generate random password:

1 Numeric
1 Special Char
1 Alphanumenrical Uppercase
9 Alphanumentical Lovercase
`
)

func main() {

	// TO DO argument managment inside or outside main package!
	// First Stet if no args passed use defaults
	var raw string
	if len(os.Args) < 2 || len(os.Args) < 9 {
		// Print defaults message
		fmt.Println(defaults)
		//create raw password
		raw = g.Pick(1, g.NS) + g.Pick(9, g.LS) + g.Pick(1, strings.ToUpper(g.LS)) + g.Pick(1, g.SS)
		//print generated password
		fmt.Println("OUTPUT: ", g.Melee(raw))
		//exit from executable
		return
	}
	fmt.Println("ARG: ", len(os.Args))
	l, _ := strconv.Atoi(os.Args[1])

	if l > g.LM {
		fmt.Printf("The vaule of l is: %d\n", l)
	} else {
		l = g.LM
		fmt.Printf("The vaule of l set to default size: %d\n", l)
	}

	//Generate raw password probably i need to refactor Pick function?
	raw = g.Pick(6, g.NS) + g.Pick(4, g.LS) + g.Pick(2, strings.ToUpper(g.LS)) + g.Pick(1, g.SS)
	//For debug purpose in this alpha state
	//fmt.Println(rawpassw)

	//Melee... Nice name no?
	fmt.Println(g.Melee(raw))
	//fmt.Printf("%08b\n", 5)
}
