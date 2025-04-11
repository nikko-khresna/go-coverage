package util_test

import (
	"regexp"
	"testing"

	"github.com/nikko-khresna/go-coverage/util"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := util.Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
