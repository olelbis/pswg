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
	ver  string = "0.1.3.1"
)

func main() {
	_, _ = exec, ver
	// TO DO argument managment inside or outside main package!
	// First Stet if no args passed use defaults
	var (
		rawpassw         string
		nOfTotalChar     int = g.MinPwdLenght
		nOfUpperCaseChar int = g.MinUpChar
		maxpwdlenght     int = g.Maxpwdlenght
		nOfSpechialChar  int = g.MinSpecChar
		nOfNumericChar   int = g.MinNumChar

		err error
		x   int = 1 //Help me to remove a magic number and prevent problem with os.Args array
	)
	// Check if aguments are less than two, in this case i use the default function
	if len(os.Args) < 2 {
		g.DefaultPasswordGenerator()
		return
	} else if len(os.Args) <= 9 {

		for i := 1; i < len(os.Args); i++ {
			switch {
			case os.Args[i] == "-l":
				if len(os.Args)-x > i {
					i++
					nOfTotalChar, err = strconv.Atoi(os.Args[i])
					if err != nil {
						fmt.Println("ERROR:", err)
						fmt.Println(g.UsageMessage)
						return
					} else if g.Ispwdtoolong(nOfTotalChar) {
						nOfTotalChar = maxpwdlenght
					}
				}

			case os.Args[i] == "-u":
				if len(os.Args)-x > i {
					i++
					nOfUpperCaseChar, err = strconv.Atoi(os.Args[i])
					if err != nil {
						fmt.Println("ERROR:", err)
						//return
					}
				}

			case os.Args[i] == "-s":
				if len(os.Args)-x > i {
					i++
					nOfSpechialChar, err = strconv.Atoi(os.Args[i])
					if err != nil {
						fmt.Println("ERROR:", err)
						//return
					}
				}

			case os.Args[i] == "-n":
				if len(os.Args)-x > i {
					i++
					nOfNumericChar, err = strconv.Atoi(os.Args[i])
					if err != nil {
						fmt.Println("ERROR:", err)
						//return
					}
				}
			case os.Args[i] == "--silent":
				return
			default:
				fmt.Println(g.DefautMessage)
				//return
			}
		}

		if nOfTotalChar < nOfUpperCaseChar+nOfSpechialChar+nOfNumericChar {
			g.DefaultPasswordGenerator()
			return
		}

		//To obtain number of Alphanumeric charater
		nOfTotalChar = nOfTotalChar - nOfUpperCaseChar - nOfSpechialChar - nOfNumericChar

		//fmt.Println("LENGHT: ", len([]rune(raw)))
		//for i := 0; i <= 9; i++ {
		rawpassw = g.PickCrypto(nOfNumericChar, g.NumericPool) +
			g.PickCrypto(nOfTotalChar, g.AlphanumericPool) +
			g.PickCrypto(nOfUpperCaseChar, strings.ToUpper(g.AlphanumericPool)) +
			g.PickCrypto(nOfSpechialChar, g.SpecialCharPool)
		melee, _ := g.Melee(rawpassw)
		fmt.Println("OUTPUT: ", melee)
		//}
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
