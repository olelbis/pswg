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
	//Nice try to made a new world full of color... Some sort uf inutility
	// Ex.
	//	 \033[ = Escape non printable character
	//   1;34m = This mean 1 for bold of intensity(SGF) and 34 is for blue (3-4bit) last nbiut not least m is for use SGF
	//   %s = String Value from printf input
	//   \033[0m = back to default
	// More infos at https://en.wikipedia.org/wiki/ANSI_escape_code
	Infocolor    = "\033[0;32m%s\033[0m"
	Noticecolor  = "\033[0;36m%s\033[0m"
	Warningcolor = "\033[1;33m%s\033[0m"
	Outputcolor  = "\033[44;30m%s\033[0m"
	Errorcolor   = "\033[1;31m%s\033[0m"
	Debugcolor   = "\033[0;36m%s\033[0m"
	//MinPwdLenght Minimum Password Lenght
	MinPwdLenght int = 12
	//Maxpwdlenght Maximum Password Lenght
	Maxpwdlenght int = 128
)

var (
	//NumericPool numeric string
	NumericPool string = "1234567890"
	//AlphanumericPool alphanumeric string
	AlphanumericPool string = "abcdefghijklmnopqrstuvwxyz"
	//SpecialCharPool special character string
	SpecialCharPool string = "!&%$£=?^+*][{}-_.:,;()><"
	//MinUpChar Uppercase n of char
	MinUpChar int = 1
	//MinAlphaChar Aplhanumeric n of char
	MinAlphaChar int = 9
	//MinSpecChar Special n of char
	MinSpecChar int = 1
	//MinNumChar Numeric n of char
	MinNumChar   int    = 1
	UsageMessage string = `Usage:
	pswg -l <Password Length (Default: 12, upper limit 128)> -u <N. of Alphanumeric Uppercase> -s <N. of Special Char> -n <N. of Numeric Char>`
	//DefautMessage default message
	DefautMessage string = `
Default password value used, possible causes: 
-	no option specified
-	wrong number of arguments

Password seed with:
1 Numeric
1 Special Char
1 Alphanumeric Uppercase
9 Alphanumeric Lovercase
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

func Ispwdtoolong(passwordlenght int) bool {
	if passwordlenght > Maxpwdlenght {
		return true
	} else {
		return false
	}
}

// PickCrypto : return random string of lenght L extract it form  K (crypto/random)
func PickCrypto(lenght int, keyrandom string) (ret string) {

	// yet another "i" loop
	for i := 1; i <= lenght; i++ {
		result, _ := cr.Int(cr.Reader, big.NewInt(int64(utf8.RuneCountInString(keyrandom))))
		//Use utf8.RuneCountInString to prevent index out of range (panic)
		//in case of multibyte character
		ret += string([]rune(keyrandom)[int(result.Int64())])
	}
	return ret
}

// DefPick : return random password based on predefined rule
func DefaultPasswordGenerator() {
	var raw string
	// Print usage message
	fmt.Printf(Infocolor, UsageMessage)
	// Print defaults message
	fmt.Printf(Infocolor, DefautMessage)
	//create raw password
	raw = PickCrypto(MinNumChar, NumericPool) + PickCrypto(MinAlphaChar, AlphanumericPool) + PickCrypto(MinUpChar, strings.ToUpper(AlphanumericPool)) + PickCrypto(MinSpecChar, SpecialCharPool)
	//print generated password
	fmt.Println("OUTPUT: " + Melee(raw))

}
