// Package genutil generates random passwords for the pswg CLI.
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

	NumericPool              = "1234567890"
	LowercasePool            = "abcdefghijklmnopqrstuvwxyz"
	SpecialCharPool          = "!&%$=?^+*][{}-_.:,;()><"
	ShellSafeSpecialCharPool = "@_:,."

	// Deprecated: use LowercasePool. The pool contains lowercase letters
	// only, so the historical name was misleading.
	AlphanumericPool = LowercasePool

	MinUpChar    = 1
	MinAlphaChar = 9
	MinSpecChar  = 1
	MinNumChar   = 1
)

// Policy describes the composition rules for generated passwords.
type Policy struct {
	Length    int
	Uppercase int
	Special   int
	Numeric   int
	ShellSafe bool
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

func (p Policy) specialPool() string {
	if p.ShellSafe {
		return ShellSafeSpecialCharPool
	}
	return SpecialCharPool
}

// Validate checks that a policy can be used to generate a password.
func (p Policy) Validate() error {
	if p.Length < MinPasswordLength {
		return errors.New("password length must be at least 12")
	}
	if p.Length > MaxPasswordLength {
		return errors.New("password length must be no more than 128")
	}
	if p.Uppercase < 0 || p.Special < 0 || p.Numeric < 0 {
		return errors.New("character counts cannot be negative")
	}
	if p.lowercase() < 0 {
		return errors.New("character requirements exceed password length")
	}
	return nil
}

// Shuffle returns password with its runes shuffled using crypto/rand.
//
// If password contains at least one non-numeric rune, the result never
// starts with a digit. This is enforced by rejection sampling (re-shuffling
// until the constraint holds), so the output is uniformly distributed over
// all permutations that satisfy the constraint. If every rune is numeric,
// the constraint is skipped and a plain uniform shuffle is returned.
func Shuffle(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}
	r := []rune(password)

	hasNonNumeric := false
	for _, c := range r {
		if !unicode.IsNumber(c) {
			hasNonNumeric = true
			break
		}
	}

	for {
		for i := len(r) - 1; i > 0; i-- {
			j, err := cryptoInt(i + 1)
			if err != nil {
				return "", err
			}
			r[i], r[j] = r[j], r[i]
		}
		if !hasNonNumeric || !unicode.IsNumber(r[0]) {
			return string(r), nil
		}
	}
}

// Deprecated: use Shuffle.
func Melee(pwdin string) (string, error) {
	return Shuffle(pwdin)
}

// IsPasswordTooLong reports whether length is above MaxPasswordLength.
func IsPasswordTooLong(length int) bool {
	return length > MaxPasswordLength
}

// Deprecated: use IsPasswordTooLong.
func Ispwdtoolong(passwordlenght int) bool {
	return IsPasswordTooLong(passwordlenght)
}

// PickRandom returns a random string of length characters sampled from pool.
func PickRandom(length int, pool string) (string, error) {
	if length < 0 {
		return "", errors.New("length cannot be negative")
	}

	poolLength := utf8.RuneCountInString(pool)
	if poolLength == 0 {
		return "", errors.New("pool cannot be empty")
	}

	var ret strings.Builder
	ret.Grow(length)
	poolRunes := []rune(pool)
	for i := 1; i <= length; i++ {
		result, err := cryptoInt(poolLength)
		if err != nil {
			return "", err
		}
		ret.WriteRune(poolRunes[result])
	}
	return ret.String(), nil
}

// Deprecated: use PickRandom.
func PickCrypto(lenght int, keyrandom string) (string, error) {
	return PickRandom(lenght, keyrandom)
}

// DefaultPassword returns a random password using the default policy.
func DefaultPassword() (string, error) {
	return Generate(DefaultPolicy())
}

// Deprecated: use DefaultPassword.
func DefaultPasswordGenerator() (string, error) {
	return DefaultPassword()
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

	numericPart, err := PickRandom(policy.Numeric, NumericPool)
	if err != nil {
		return "", err
	}
	lowercasePart, err := PickRandom(policy.lowercase(), LowercasePool)
	if err != nil {
		return "", err
	}
	uppercasePart, err := PickRandom(policy.Uppercase, strings.ToUpper(LowercasePool))
	if err != nil {
		return "", err
	}
	specialPart, err := PickRandom(policy.Special, policy.specialPool())
	if err != nil {
		return "", err
	}

	return Shuffle(numericPart + lowercasePart + uppercasePart + specialPart)
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
