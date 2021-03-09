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
	var (
		raw string
		nLM int = g.LM
		nUC int = g.UC

		nSC int = g.SC
		nNC int = g.NC
	)

	fmt.Println(len(os.Args))
	if len(os.Args) < 2 {
		// Print defaults message
		fmt.Println(defaults)
		//create raw password
		raw = g.Pick(nNC, g.NS) + g.Pick(nLM, g.LS) + g.Pick(nUC, strings.ToUpper(g.LS)) + g.Pick(nSC, g.SS)
		//print generated password
		fmt.Println("OUTPUT: ", g.Melee(raw))
		//exit from executable
		return
	} else if len(os.Args) <= 9 {

		for i := 1; i < len(os.Args); i++ {
			fmt.Println("ARG: ", os.Args[i])
			switch {
			case os.Args[i] == "-l":
				i++
				nLM, _ = strconv.Atoi(os.Args[i])
				fmt.Println("ELLE:", nLM)
			case os.Args[i] == "-u":
				i++
				nUC, _ = strconv.Atoi(os.Args[i])
				fmt.Println("UUUU", nUC)
			case os.Args[i] == "-s":
				i++
				nSC, _ = strconv.Atoi(os.Args[i])
				fmt.Println("SSSS:", nSC)
			case os.Args[i] == "-n":
				i++
				nNC, _ = strconv.Atoi(os.Args[i])
				fmt.Println("NNNN:", nNC)
			}
		}
		nLM = nLM - nUC - nSC - nNC
		raw = g.Pick(nNC, g.NS) + g.Pick(nLM, g.LS) + g.Pick(nUC, strings.ToUpper(g.LS)) + g.Pick(nSC, g.SS)
		fmt.Println("LENGHT: ", len(raw))
		fmt.Println("OUTPUT: ", g.Melee(raw))
	}
	/*
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
	*/
}
