package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/olelbis/pswg/genutil"
)

const (
	exec string = "pswg"
	ver  string = "0.4-alpha"
)

func main() {
	_, _ = exec, ver

	if len(os.Args) < 2 {
		if err := genutil.DefaultPasswordGenerator(); err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
		return
	}

	flags := flag.NewFlagSet(exec, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	length := flags.Int("l", genutil.MinPwdLenght, "password length")
	uppercase := flags.Int("u", genutil.MinUpChar, "number of uppercase characters")
	special := flags.Int("s", genutil.MinSpecChar, "number of special characters")
	numeric := flags.Int("n", genutil.MinNumChar, "number of numeric characters")
	silent := flags.Bool("silent", false, "do not generate output")
	showVersion := flags.Bool("version", false, "print version")
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, genutil.UsageMessage)
	}

	if err := flags.Parse(os.Args[1:]); err != nil {
		os.Exit(2)
	}
	if *showVersion {
		fmt.Printf("%s %s\n", exec, ver)
		return
	}
	if *silent {
		return
	}
	if flags.NArg() > 0 {
		fmt.Fprintln(os.Stderr, "ERROR: unexpected argument:", flags.Arg(0))
		flags.Usage()
		os.Exit(2)
	}
	if *length < genutil.MinPwdLenght {
		fmt.Fprintf(os.Stderr, "ERROR: password length must be at least %d\n", genutil.MinPwdLenght)
		os.Exit(2)
	}
	if genutil.Ispwdtoolong(*length) {
		*length = genutil.Maxpwdlenght
	}
	if *uppercase < 0 || *special < 0 || *numeric < 0 {
		fmt.Fprintln(os.Stderr, "ERROR: character counts cannot be negative")
		os.Exit(2)
	}

	password, err := genutil.GeneratePassword(*length, *uppercase, *special, *numeric)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
	fmt.Println("OUTPUT:", password)
}
