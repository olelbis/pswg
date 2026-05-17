package genutil

import (
	cr "crypto/rand"
	"errors"
	"math/big"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	// MinPasswordLength is the minimum generated password length.
	MinPasswordLength = 12
	// MaxPasswordLength is the maximum generated password length.
	MaxPasswordLength = 128

	// Deprecated: use MinPasswordLength.
	MinPwdLenght = MinPasswordLength
	// Deprecated: use MaxPasswordLength.
	Maxpwdlenght = MaxPasswordLength

	NumericPool      = "1234567890"
	AlphanumericPool = "abcdefghijklmnopqrstuvwxyz"
	SpecialCharPool  = "!&%$£=?^+*][{}-_.:,;()><"

	MinUpChar    = 1
	MinAlphaChar = 9
	MinSpecChar  = 1
	MinNumChar   = 1
)

var UsageMessage string = `Usage:
	pswg [-l <Password Length (Default: 12, upper limit 128)>] [-u <N. of Alphanumeric Uppercase>] [-s <N. of Special Char>] [-n <N. of Numeric Char>]
	pswg -version
	pswg --silent`

// Policy describes the composition rules for generated passwords.
type Policy struct {
	Length    int
	Uppercase int
	Special   int
	Numeric   int
}

// DefaultPolicy returns the default password policy.
func DefaultPolicy() Policy {
	return Policy{
		Length:    MinPasswordLength,
		Uppercase: MinUpChar,
		Special:   MinSpecChar,
		Numeric:   MinNumChar,
	}
}

func (p Policy) lowercase() int {
	return p.Length - p.Uppercase - p.Special - p.Numeric
}

// Validate checks that a policy can be used to generate a password.
func (p Policy) Validate() error {
	if p.Length < MinPasswordLength {
		return errors.New("password length is shorter than minimum length")
	}
	if p.Length > MaxPasswordLength {
		return errors.New("password length is longer than maximum length")
	}
	if p.Uppercase < 0 || p.Special < 0 || p.Numeric < 0 {
		return errors.New("character counts cannot be negative")
	}
	if p.lowercase() < 0 {
		return errors.New("character requirements exceed password length")
	}
	return nil
}

// Melee returns pwdin with its runes shuffled using crypto/rand.
func Melee(pwdin string) (string, error) {
	if utf8.RuneCountInString(pwdin) < MinPasswordLength {
		return "", errors.New("password is shorter than minimum length")
	}
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
	return passwordlenght > MaxPasswordLength
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

	var ret strings.Builder
	ret.Grow(lenght)
	keyrandomRunes := []rune(keyrandom)
	for i := 1; i <= lenght; i++ {
		result, err := cryptoInt(keyrandomLength)
		if err != nil {
			return "", err
		}
		ret.WriteRune(keyrandomRunes[result])
	}
	return ret.String(), nil
}

// DefaultPasswordGenerator returns a random password using the default policy.
func DefaultPasswordGenerator() (string, error) {
	return Generate(DefaultPolicy())
}

// GeneratePassword returns a random password using the requested composition.
func GeneratePassword(length, uppercase, special, numeric int) (string, error) {
	return Generate(Policy{
		Length:    length,
		Uppercase: uppercase,
		Special:   special,
		Numeric:   numeric,
	})
}

// Generate returns a random password using policy.
func Generate(policy Policy) (string, error) {
	if err := policy.Validate(); err != nil {
		return "", err
	}

	numericPart, err := PickCrypto(policy.Numeric, NumericPool)
	if err != nil {
		return "", err
	}
	lowercasePart, err := PickCrypto(policy.lowercase(), AlphanumericPool)
	if err != nil {
		return "", err
	}
	uppercasePart, err := PickCrypto(policy.Uppercase, strings.ToUpper(AlphanumericPool))
	if err != nil {
		return "", err
	}
	specialPart, err := PickCrypto(policy.Special, SpecialCharPool)
	if err != nil {
		return "", err
	}

	return Melee(numericPart + lowercasePart + uppercasePart + specialPart)
}

func cryptoInt(max int) (int, error) {
	if max <= 0 {
		return 0, errors.New("max must be positive")
	}
	result, err := cr.Int(cr.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(result.Int64()), nil
}
