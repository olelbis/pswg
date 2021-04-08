package main

import (
	"fmt"
	"os"
	g "passgen/genutil"
	"strconv"
	"strings"
)

const (
	Exec string = "pswg"
	Ver  string = "0.4a"
)

func main() {

	// TO DO argument managment inside or outside main package!
	// First Stet if no args passed use defaults
	var (
		raw string
		nLM int = g.MinPwdLenght
		nUC int = g.MinUpChar

		nSC int = g.MinSpecChar
		nNC int = g.MinNumChar
		err error
		x   int = 1 //Help me to remove a magic number and prevent problem with os.Args array
	)

	if len(os.Args) < 2 {
		g.DefPick()
		return
	} else if len(os.Args) <= 9 {

		for i := 1; i < len(os.Args); i++ {
			switch {
			case os.Args[i] == "-l":
				if len(os.Args)-x > i {
					i++
					nLM, err = strconv.Atoi(os.Args[i])
					if err != nil {
						fmt.Println("ERROR:", err)
					}
				}

			case os.Args[i] == "-u":
				if len(os.Args)-x > i {
					i++
					nUC, _ = strconv.Atoi(os.Args[i])
					if err != nil {
						fmt.Println("ERROR:", err)
					}
				}

			case os.Args[i] == "-s":
				if len(os.Args)-x > i {
					i++
					nSC, _ = strconv.Atoi(os.Args[i])
					if err != nil {
						fmt.Println("ERROR:", err)
					}
				}

			case os.Args[i] == "-n":
				if len(os.Args)-x > i {
					i++
					nNC, _ = strconv.Atoi(os.Args[i])
					if err != nil {
						fmt.Println("ERROR:", err)
					}
				}
			default:
				fmt.Println(g.DEFMSG)
			}
		}

		if nLM < nUC+nSC+nNC {
			g.DefPick()
			return
		}

		//To obtain number of Alphanumeric charater
		nLM = nLM - nUC - nSC - nNC

		raw = g.PickCrypto(nNC, g.NumericPool) + g.PickCrypto(nLM, g.AlphanumericPool) + g.PickCrypto(nUC, strings.ToUpper(g.AlphanumericPool)) + g.PickCrypto(nSC, g.SpecialCharPool)
		//fmt.Println("LENGHT: ", len([]rune(raw)))
		fmt.Println("OUTPUT: ", g.Melee(raw))
	}
	//fmt.Println("TEST:", g.PickCrypto(nLM, g.LS))
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
