package main

import (
	"fmt"
	"os"
	g "passgen/genutil"
	"strconv"
	"strings"
)

const (
	exec string = "pswg"
	ver  string = "0.3a"
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
		err error
	)

	if len(os.Args) < 2 {
		g.DefPick()
		return
	} else if len(os.Args) <= 9 {

		for i := 1; i < len(os.Args); i++ {
			switch {
			case os.Args[i] == "-l":
				if len(os.Args)-1 > i {
					i++
					nLM, err = strconv.Atoi(os.Args[i])
					if err != nil {
						fmt.Println("ERRORE:", err)
					}
				}

			case os.Args[i] == "-u":
				if len(os.Args)-1 > i {
					i++
					nUC, _ = strconv.Atoi(os.Args[i])
				}

			case os.Args[i] == "-s":
				if len(os.Args)-1 > i {
					i++
					nSC, _ = strconv.Atoi(os.Args[i])
				}

			case os.Args[i] == "-n":
				if len(os.Args)-1 > i {
					i++
					nNC, _ = strconv.Atoi(os.Args[i])
				}
			default:
				fmt.Println(g.DEFMSG)
			}
		}

		if nLM < (nUC + nSC + nNC) {
			g.DefPick()
			return
		}
		nLM = nLM - nUC - nSC - nNC
		//fmt.Println("VAL: ", nLM)
		//raw = g.Pick(nNC, g.NS) + g.Pick(nLM, g.LS) + g.Pick(nUC, strings.ToUpper(g.LS)) + g.Pick(nSC, g.SS)
		raw = g.PickCrypto(nNC, g.NS) + g.PickCrypto(nLM, g.LS) + g.PickCrypto(nUC, strings.ToUpper(g.LS)) + g.PickCrypto(nSC, g.SS)
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
