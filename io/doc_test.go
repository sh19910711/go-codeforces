package main_test

import (
	"bytes"
	io "github.com/sh19910711/go-codeforces/io"
	"strings"
	"testing"
)

func Test810B(t *testing.T) {
	buf := &bytes.Buffer{}
	io.Examples(buf, "810", "B", ".input")
	s := buf.String()

	if !strings.Contains(s, "4 2") {
		t.Errorf("input should contain 4 2")
	}

	if strings.Contains(s, "<br/>") {
		t.Errorf("input should not contain <br>")
	}
}
