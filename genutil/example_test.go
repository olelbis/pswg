package genutil_test

import (
	"fmt"
	"unicode/utf8"

	"github.com/olelbis/pswg/genutil"
)

func ExampleDefaultPassword() {
	password, err := genutil.DefaultPassword()
	if err != nil {
		panic(err)
	}

	fmt.Println(utf8.RuneCountInString(password))
	// Output: 12
}

func ExampleGenerate() {
	password, err := genutil.Generate(genutil.Policy{
		Length:    24,
		Uppercase: 4,
		Special:   2,
		Numeric:   4,
		ShellSafe: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(utf8.RuneCountInString(password))
	// Output: 24
}

func ExamplePickRandom() {
	value, err := genutil.PickRandom(6, genutil.NumericPool)
	if err != nil {
		panic(err)
	}

	fmt.Println(utf8.RuneCountInString(value))
	// Output: 6
}
