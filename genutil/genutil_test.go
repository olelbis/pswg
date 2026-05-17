package genutil

import (
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestMeleeRejectsShortPassword(t *testing.T) {
	if melee, err := Melee(""); err == nil {
		t.Fatalf("Melee(\"\") = %q, nil; want error", melee)
	}
}

func TestMeleeFullKeepsLength(t *testing.T) {
	input := "S0n0B1sKu3D0!!"
	melee, err := Melee(input)
	if err != nil {
		t.Fatal(err)
	}
	if utf8.RuneCountInString(melee) != utf8.RuneCountInString(input) {
		t.Fatalf("Melee(%q) length = %d; want %d", input, utf8.RuneCountInString(melee), utf8.RuneCountInString(input))
	}
}

func TestMeleeAllNumericTerminates(t *testing.T) {
	input := "123456789012"
	melee, err := Melee(input)
	if err != nil {
		t.Fatal(err)
	}
	if utf8.RuneCountInString(melee) != utf8.RuneCountInString(input) {
		t.Fatalf("Melee(%q) length = %d; want %d", input, utf8.RuneCountInString(melee), utf8.RuneCountInString(input))
	}
}

func TestPickCryptoRejectsEmptyPool(t *testing.T) {
	if got, err := PickCrypto(1, ""); err == nil {
		t.Fatalf("PickCrypto(1, \"\") = %q, nil; want error", got)
	}
}

func TestSpecialCharPoolIsASCII(t *testing.T) {
	for _, r := range SpecialCharPool {
		if r > unicode.MaxASCII {
			t.Fatalf("SpecialCharPool contains non-ASCII rune %q", r)
		}
	}
}

func TestShellSafeSpecialCharPool(t *testing.T) {
	for _, r := range ShellSafeSpecialCharPool {
		if r > unicode.MaxASCII {
			t.Fatalf("ShellSafeSpecialCharPool contains non-ASCII rune %q", r)
		}
		if strings.ContainsRune("!#$&'()*;<>?[\\]`{|}~ \"\n\t-", r) {
			t.Fatalf("ShellSafeSpecialCharPool contains shell-risky rune %q", r)
		}
	}
}

func TestGeneratePasswordUsesRequestedLength(t *testing.T) {
	password, err := GeneratePassword(16, 2, 2, 2)
	if err != nil {
		t.Fatal(err)
	}
	if utf8.RuneCountInString(password) != 16 {
		t.Fatalf("GeneratePassword length = %d; want 16", utf8.RuneCountInString(password))
	}
	if !strings.ContainsAny(password, NumericPool) {
		t.Fatal("GeneratePassword did not include numeric characters")
	}
}

func TestGeneratePasswordUsesRequestedComposition(t *testing.T) {
	password, err := GeneratePassword(16, 2, 3, 4)
	if err != nil {
		t.Fatal(err)
	}

	var uppercase, special, numeric, lowercase int
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			uppercase++
		case unicode.IsNumber(r):
			numeric++
		case strings.ContainsRune(SpecialCharPool, r):
			special++
		case unicode.IsLower(r):
			lowercase++
		default:
			t.Fatalf("unexpected rune %q in password %q", r, password)
		}
	}

	if uppercase != 2 || special != 3 || numeric != 4 || lowercase != 7 {
		t.Fatalf("composition uppercase=%d special=%d numeric=%d lowercase=%d; want 2, 3, 4, 7", uppercase, special, numeric, lowercase)
	}
}

func TestGeneratePasswordCanUseShellSafeSpecials(t *testing.T) {
	password, err := Generate(Policy{Length: 16, Uppercase: 2, Special: 4, Numeric: 2, ShellSafe: true})
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range password {
		if strings.ContainsRune(SpecialCharPool, r) && !strings.ContainsRune(ShellSafeSpecialCharPool, r) {
			t.Fatalf("password %q contains non-shell-safe special rune %q", password, r)
		}
	}
}

func TestGeneratePasswordRejectsInvalidPolicy(t *testing.T) {
	tests := []struct {
		name   string
		policy Policy
	}{
		{
			name:   "too short",
			policy: Policy{Length: MinPasswordLength - 1, Uppercase: 1, Special: 1, Numeric: 1},
		},
		{
			name:   "too long",
			policy: Policy{Length: MaxPasswordLength + 1, Uppercase: 1, Special: 1, Numeric: 1},
		},
		{
			name:   "negative count",
			policy: Policy{Length: MinPasswordLength, Uppercase: -1, Special: 1, Numeric: 1},
		},
		{
			name:   "requirements exceed length",
			policy: Policy{Length: MinPasswordLength, Uppercase: MinPasswordLength + 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Generate(tt.policy); err == nil {
				t.Fatalf("Generate(%+v) = %q, nil; want error", tt.policy, got)
			}
		})
	}
}

func TestDefaultPasswordGenerator(t *testing.T) {
	password, err := DefaultPasswordGenerator()
	if err != nil {
		t.Fatal(err)
	}
	if utf8.RuneCountInString(password) != MinPasswordLength {
		t.Fatalf("DefaultPasswordGenerator length = %d; want %d", utf8.RuneCountInString(password), MinPasswordLength)
	}
}
