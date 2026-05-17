package genutil

import (
	cr "crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
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
	pswg [-l <Password Length (Default: 12, upper limit 128)>] [-u <N. of Alphanumeric Uppercase>] [-s <N. of Special Char>] [-n <N. of Numeric Char>]
	pswg -version
	pswg --silent`
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

// Melee : Get a string in input e do some shuffle
func Melee(pwdin string) (string, error) {
	if utf8.RuneCountInString(pwdin) < MinPwdLenght {
		return "", errors.New("password is shorter than minimum length")
	}
	// Transform string to rune
	r := []rune(pwdin)

	for i := len(r) - 1; i > 0; i-- {
		j, err := cryptoInt(i + 1)
		if err != nil {
			return "", err
		}
		r[i], r[j] = r[j], r[i]
	}

	if unicode.IsNumber(r[0]) {
		for i := 1; i < len(r); i++ {
			if !unicode.IsNumber(r[i]) {
				r[0], r[i] = r[i], r[0]
				return string(r), nil
			}
		}
	}

	return string(r), nil
}

func Ispwdtoolong(passwordlenght int) bool {
	return passwordlenght > Maxpwdlenght
}

// PickCrypto : return random string of lenght L extract it form  K (crypto/random)
func PickCrypto(lenght int, keyrandom string) (string, error) {
	if lenght < 0 {
		return "", errors.New("length cannot be negative")
	}

	keyrandomLength := utf8.RuneCountInString(keyrandom)
	if keyrandomLength == 0 {
		return "", errors.New("keyrandom cannot be empty")
	}

	var ret string
	keyrandomRunes := []rune(keyrandom)
	// yet another "i" loop
	for i := 1; i <= lenght; i++ {
		result, err := cryptoInt(keyrandomLength)
		if err != nil {
			return "", err
		}
		//Use utf8.RuneCountInString to prevent index out of range (panic)
		//in case of multibyte character
		ret += string(keyrandomRunes[result])
	}
	return ret, nil
}

// DefPick : return random password based on predefined rule
func DefaultPasswordGenerator() error {
	// Print usage message
	fmt.Printf(Infocolor, UsageMessage)
	// Print defaults message
	fmt.Printf(Infocolor, DefautMessage)
	//create raw password
	raw, err := GeneratePassword(MinPwdLenght, MinUpChar, MinSpecChar, MinNumChar)
	if err != nil {
		return err
	}
	//print generated password
	fmt.Println("OUTPUT: " + raw)
	return nil
}

func GeneratePassword(length, uppercase, special, numeric int) (string, error) {
	lowercase := length - uppercase - special - numeric
	if lowercase < 0 {
		return "", errors.New("character requirements exceed password length")
	}

	numericPart, err := PickCrypto(numeric, NumericPool)
	if err != nil {
		return "", err
	}
	lowercasePart, err := PickCrypto(lowercase, AlphanumericPool)
	if err != nil {
		return "", err
	}
	uppercasePart, err := PickCrypto(uppercase, strings.ToUpper(AlphanumericPool))
	if err != nil {
		return "", err
	}
	specialPart, err := PickCrypto(special, SpecialCharPool)
	if err != nil {
		return "", err
	}

	return Melee(numericPart + lowercasePart + uppercasePart + specialPart)
}

func cryptoInt(max int) (int, error) {
	result, err := cr.Int(cr.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(result.Int64()), nil
}
