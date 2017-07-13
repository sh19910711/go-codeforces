package main_test

import (
	"bytes"
	io "github.com/sh19910711/go-codeforces/io"
	"strings"
	"testing"
)

func input(c, p string) string {
	buf := &bytes.Buffer{}
	io.Examples(buf, c, p, ".input")
	return buf.String()
}

func Test810B(t *testing.T) {
	s := input("810", "B")

	if !strings.Contains(s, "4 2") {
		t.Errorf("input should contain 4 2")
	}

	if strings.Contains(s, "<br/>") {
		t.Errorf("input should not contain <br>")
	}
}

func Test828C(t *testing.T) {
	s := input("828", "C")

	if strings.Contains(s, "<br/>") {
		t.Errorf("input should not contain <br>")
	}
}
