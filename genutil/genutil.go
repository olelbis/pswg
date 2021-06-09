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
	//Nice try to made a new world full of color
	// Ex.
	//	 \033[ = Escape non printable character
	//   1;34m = This mean 1 for bold of intensity(SGF) and 34 is for blue (3-4bit) last nbiut not least m is for use SGF
	//   %s = String Value from printf input
	//   \033[0m = back to default
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
	//NS numeric string
	NumericPool string = "1234567890"
	//LS alphanumeric string
	AlphanumericPool string = "abcdefghijklmnopqrstuvwxyz"
	//SS special character string
	SpecialCharPool string = "!&%$£=?^+*][{}-_.:,;()><"
	//MinPwdLenght Minimum Password Lenght
	MinPwdLenght int = 12
	//Maxpwdlenght Maximum Password Lenght
	Maxpwdlenght int = 128
	//UC Uppercase n of char
	MinUpChar int = 1
	//AC Aplhanumeric n of char
	MinAlphaChar int = 9
	//SC Special n of char
	MinSpecChar int = 1
	//NC Numeric n of char
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
	// Print usage message
	fmt.Printf(InfoColor, UsageMessage)
	// Print defaults message
	fmt.Printf(NoticeColor, DefautMessage)
	//create raw password
	raw = PickCrypto(MinNumChar, NumericPool) + PickCrypto(MinAlphaChar, AlphanumericPool) + PickCrypto(MinUpChar, strings.ToUpper(AlphanumericPool)) + PickCrypto(MinSpecChar, SpecialCharPool)
	//print generated password
	fmt.Printf(NoticeColor, "OUTPUT: "+Melee(raw))

}
