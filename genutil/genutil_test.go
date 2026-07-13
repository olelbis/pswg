package genutil

import (
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestShuffleRejectsEmptyPassword(t *testing.T) {
	if shuffled, err := Shuffle(""); err == nil {
		t.Fatalf("Shuffle(\"\") = %q, nil; want error", shuffled)
	}
}

func TestShuffleKeepsLength(t *testing.T) {
	input := "S0n0B1sKu3D0!!"
	shuffled, err := Shuffle(input)
	if err != nil {
		t.Fatal(err)
	}
	if utf8.RuneCountInString(shuffled) != utf8.RuneCountInString(input) {
		t.Fatalf("Shuffle(%q) length = %d; want %d", input, utf8.RuneCountInString(shuffled), utf8.RuneCountInString(input))
	}
}

func TestShuffleAllNumericTerminates(t *testing.T) {
	input := "123456789012"
	shuffled, err := Shuffle(input)
	if err != nil {
		t.Fatal(err)
	}
	if utf8.RuneCountInString(shuffled) != utf8.RuneCountInString(input) {
		t.Fatalf("Shuffle(%q) length = %d; want %d", input, utf8.RuneCountInString(shuffled), utf8.RuneCountInString(input))
	}
}

func TestPickRandomRejectsEmptyPool(t *testing.T) {
	if got, err := PickRandom(1, ""); err == nil {
		t.Fatalf("PickRandom(1, \"\") = %q, nil; want error", got)
	}
}

func TestPickRandomRejectsNegativeLength(t *testing.T) {
	if got, err := PickRandom(-1, NumericPool); err == nil {
		t.Fatalf("PickRandom(-1, NumericPool) = %q, nil; want error", got)
	}
}

func TestIsPasswordTooLong(t *testing.T) {
	if IsPasswordTooLong(MaxPasswordLength) {
		t.Fatalf("IsPasswordTooLong(%d) = true; want false", MaxPasswordLength)
	}
	if !IsPasswordTooLong(MaxPasswordLength + 1) {
		t.Fatalf("IsPasswordTooLong(%d) = false; want true", MaxPasswordLength+1)
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

func TestDefaultPassword(t *testing.T) {
	password, err := DefaultPassword()
	if err != nil {
		t.Fatal(err)
	}
	if utf8.RuneCountInString(password) != MinPasswordLength {
		t.Fatalf("DefaultPassword length = %d; want %d", utf8.RuneCountInString(password), MinPasswordLength)
	}
}

func TestDeprecatedCompatibilityWrappers(t *testing.T) {
	if _, err := Melee("S0n0B1sKu3D0!!"); err != nil {
		t.Fatalf("Melee compatibility wrapper returned error: %v", err)
	}
	if got, err := PickCrypto(4, NumericPool); err != nil || utf8.RuneCountInString(got) != 4 {
		t.Fatalf("PickCrypto compatibility wrapper = %q, %v; want 4 runes, nil", got, err)
	}
	if _, err := DefaultPasswordGenerator(); err != nil {
		t.Fatalf("DefaultPasswordGenerator compatibility wrapper returned error: %v", err)
	}
	if !Ispwdtoolong(MaxPasswordLength + 1) {
		t.Fatal("Ispwdtoolong compatibility wrapper = false; want true")
	}
}

func TestShuffleNeverStartsWithDigitWhenPossible(t *testing.T) {
	input := "1234567890ab"
	for i := 0; i < 200; i++ {
		shuffled, err := Shuffle(input)
		if err != nil {
			t.Fatal(err)
		}
		if unicode.IsNumber([]rune(shuffled)[0]) {
			t.Fatalf("Shuffle(%q) = %q; starts with a digit", input, shuffled)
		}
	}
}

func TestShuffleShortNonEmptyAllowed(t *testing.T) {
	shuffled, err := Shuffle("ab1")
	if err != nil {
		t.Fatalf("Shuffle(\"ab1\") returned error: %v", err)
	}
	if utf8.RuneCountInString(shuffled) != 3 {
		t.Fatalf("Shuffle(\"ab1\") length = %d; want 3", utf8.RuneCountInString(shuffled))
	}
}
