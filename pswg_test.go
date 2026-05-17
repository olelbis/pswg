package main

import (
	"bytes"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/olelbis/pswg/genutil"
)

func TestRunDefaultPrintsOnlyPassword(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run(nil, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("run(nil) exit code = %d; want 0", code)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q; want empty", stderr.String())
	}

	password := strings.TrimSpace(stdout.String())
	if strings.Contains(password, "OUTPUT") || strings.Contains(password, "Usage") {
		t.Fatalf("stdout = %q; want password only", stdout.String())
	}
	if utf8.RuneCountInString(password) != genutil.MinPasswordLength {
		t.Fatalf("password length = %d; want %d", utf8.RuneCountInString(password), genutil.MinPasswordLength)
	}
}

func TestRunHelpReturnsSuccess(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"-h"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("run(-h) exit code = %d; want 0", code)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q; want empty", stdout.String())
	}
	if !strings.Contains(stderr.String(), "Usage:") {
		t.Fatalf("stderr = %q; want usage", stderr.String())
	}
}

func TestRunRejectsTooLongPassword(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"-l", "9999"}, &stdout, &stderr)
	if code != 2 {
		t.Fatalf("run(-l 9999) exit code = %d; want 2", code)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q; want empty", stdout.String())
	}
	if !strings.Contains(stderr.String(), "no more than 128") {
		t.Fatalf("stderr = %q; want maximum length error", stderr.String())
	}
}

func TestRunRejectsSilentFlag(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"--silent"}, &stdout, &stderr)
	if code != 2 {
		t.Fatalf("run(--silent) exit code = %d; want 2", code)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q; want empty", stdout.String())
	}
	if !strings.Contains(stderr.String(), "flag provided but not defined") {
		t.Fatalf("stderr = %q; want unknown flag error", stderr.String())
	}
}

func TestRunVersionUsesInjectedVersion(t *testing.T) {
	oldVersion := version
	version = "test-version"
	t.Cleanup(func() {
		version = oldVersion
	})

	var stdout, stderr bytes.Buffer
	code := run([]string{"-version"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("run(-version) exit code = %d; want 0", code)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q; want empty", stderr.String())
	}
	if got := strings.TrimSpace(stdout.String()); got != "pswg test-version" {
		t.Fatalf("stdout = %q; want version", got)
	}
}
