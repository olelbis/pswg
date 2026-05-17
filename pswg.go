package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/olelbis/pswg/genutil"
)

const (
	exec = "pswg"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		return generate(genutil.DefaultPolicy(), stdout, stderr)
	}

	flags := flag.NewFlagSet(exec, flag.ContinueOnError)
	flags.SetOutput(stderr)
	length := flags.Int("l", genutil.MinPasswordLength, "password length")
	uppercase := flags.Int("u", genutil.MinUpChar, "number of uppercase characters")
	special := flags.Int("s", genutil.MinSpecChar, "number of special characters")
	numeric := flags.Int("n", genutil.MinNumChar, "number of numeric characters")
	showVersion := flags.Bool("version", false, "print version")
	flags.Usage = func() {
		fmt.Fprintf(stderr, `Usage:
	%s [-l length] [-u uppercase] [-s special] [-n numeric]
	%s -version

Options:
`, exec, exec)
		flags.PrintDefaults()
	}

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if *showVersion {
		fmt.Fprintf(stdout, "%s %s\n", exec, displayVersion())
		if commit != "unknown" || date != "unknown" {
			fmt.Fprintf(stdout, "commit %s\nbuilt %s\n", commit, date)
		}
		return 0
	}
	if flags.NArg() > 0 {
		fmt.Fprintln(stderr, "ERROR: unexpected argument:", flags.Arg(0))
		flags.Usage()
		return 2
	}

	return generate(genutil.Policy{
		Length:    *length,
		Uppercase: *uppercase,
		Special:   *special,
		Numeric:   *numeric,
	}, stdout, stderr)
}

func generate(policy genutil.Policy, stdout, stderr io.Writer) int {
	password, err := genutil.Generate(policy)
	if err != nil {
		fmt.Fprintln(stderr, "ERROR:", err)
		return 2
	}
	fmt.Fprintln(stdout, password)
	return 0
}

func displayVersion() string {
	if version != "dev" {
		return version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" || info.Main.Version == "(devel)" {
		return version
	}
	return info.Main.Version
}
