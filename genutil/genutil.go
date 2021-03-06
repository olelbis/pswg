package genutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	//NS numeric string
	NS string = "1234567890"
	//LS alphanumeric string
	LS string = "abcdefghijklmnopqrstuvwxyz"
	//SS special character string
	SS string = "!&%$£=?^+*][{}-_.:,;()><"
	//LM Minimum Lenght
	LM int = 12
	//UC Uppercase n of char
	UC int = 1
	//AC Aplhanumeric n of char
	AC int = 9
	//SC Special n of char
	SC int = 1
	//NC Numeric n of char
	NC int = 1
	//DEFMSG default message
	DEFMSG string = `No option specified or wrong number of arguments, use the defaults for generate random password:

1 Numeric
1 Special Char
1 Alphanumenrical Uppercase
9 Alphanumentical Lovercase
`
)

//Melee : Get a string in input e do some shuffle
func Melee(pwdin string) string {
	rand.Seed(time.Now().Unix())
	// Transform string to rune
	r := []rune(pwdin)
	//Use rand.Shuffle function to make a pseudo randomic char swap
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
	return string(r)
}

// Pick : return random string of lenght L extract it form  K
func Pick(L int, K string) (ret string) {
	rand.Seed(time.Now().Unix())

	// yet another "i" loop
	for i := 1; i <= L; i++ {

		//Use utf8.RuneCountInString to prevent index out of range (panic)
		//in case of multibyte character
		ret += string([]rune(K)[rand.Intn(utf8.RuneCountInString(K))])
	}
	return ret
}

// DefPick : return random password based on predefined rule
func DefPick() {
	var raw string
	// Print defaults message
	fmt.Println(DEFMSG)
	//create raw password
	raw = Pick(NC, NS) + Pick(AC, LS) + Pick(UC, strings.ToUpper(LS)) + Pick(SC, SS)
	//print generated password
	fmt.Println("OUTPUT: ", Melee(raw))
}
