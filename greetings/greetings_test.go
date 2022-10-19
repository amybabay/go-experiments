package greetings

import (
    "testing"
    "regexp"
)

func TestHello(t *testing.T) {
    // Valid input; should succeed
    name := "Amy"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello(name)
    if !want.MatchString(msg) || err != nil {
        t.Errorf(`Hello("Amy") == %q, %v want match for %#q, nil`, msg, err, want)
    }
}

func TestHelloEmpty(t *testing.T) {
    // Invalid input; should return error
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Errorf(`Hello("") == %q, %v, want "", error`, msg, err)
    }
}
