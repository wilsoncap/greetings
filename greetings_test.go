package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Wilson"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Wilson")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Wilson) = %q, %v, requiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, requiere "", error`, msg, err)
	}
}
