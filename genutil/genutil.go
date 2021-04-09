package genutil

import (
	cr "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	//NS numeric string
	NumericPool string = "1234567890"
	//LS alphanumeric string
	AlphanumericPool string = "abcdefghijklmnopqrstuvwxyz"
	//SS special character string
	SpecialCharPool string = "!&%$£=?^+*][{}-_.:,;()><"
	//LM Minimum Lenght
	MinPwdLenght int = 12
	//UC Uppercase n of char
	MinUpChar int = 1
	//AC Aplhanumeric n of char
	MinAlphaChar int = 9
	//SC Special n of char
	MinSpecChar int = 1
	//NC Numeric n of char
	MinNumChar int = 1
	//DEFMSG default message
	DefautMessage string = `No option specified or wrong number of arguments, use the defaults for generate random password:

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
random:
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
	// Check if first char it's a number
	if unicode.IsNumber(r[0]) {
		//If true goto rand.Shuffle function using label
		goto random
	}
	return string(r)
}

// Osolete i'll remove it
// Pick : return random string of lenght L extract it form  K (math/random)
/* func Pick(L int, K string) (ret string) {
	rand.Seed(time.Now().Unix())

	// yet another "i" loop
	for i := 1; i <= L; i++ {

		//Use utf8.RuneCountInString to prevent index out of range (panic)
		//in case of multibyte character
		ret += string([]rune(K)[rand.Intn(utf8.RuneCountInString(K))])
	}
	return ret
} */

// PickCrypto : return random string of lenght L extract it form  K (crypto/random)
func PickCrypto(L int, K string) (ret string) {

	// yet another "i" loop
	for i := 1; i <= L; i++ {
		result, _ := cr.Int(cr.Reader, big.NewInt(int64(utf8.RuneCountInString(K))))
		//Use utf8.RuneCountInString to prevent index out of range (panic)
		//in case of multibyte character
		ret += string([]rune(K)[int(result.Int64())])
	}
	return ret
}

// DefPick : return random password based on predefined rule
func DefaultPasswordGenerator() {
	var raw string
	// Print defaults message
	fmt.Println(DefautMessage)
	//create raw password
	raw = PickCrypto(MinNumChar, NumericPool) + PickCrypto(MinAlphaChar, AlphanumericPool) + PickCrypto(MinUpChar, strings.ToUpper(AlphanumericPool)) + PickCrypto(MinSpecChar, SpecialCharPool)
	//print generated password
	fmt.Println("OUTPUT: ", Melee(raw))
}
