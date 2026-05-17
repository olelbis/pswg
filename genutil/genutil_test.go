package genutil

import (
	"strings"
	"testing"
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
