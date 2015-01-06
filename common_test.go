package formatifier

import (
	"strings"
	"testing"
)

var (
	testString  = "aBcDeFgHiJkLmNoPqRsTuVwXyZ"
	testString2 = "aBcD8FgHiJ9LmNoPqR7TuVwXyZ"
	testString3 = "aBcD8FgHiJ9Lm*oPqR|TuVwXyZ"
)

func TestNew(t *testing.T) {
	t.Parallel()
	n := New(testString)
	if len(n.theString) != 26 {
		t.Error("Unusable instance of New()")
	}
}

func TestmakeLower(t *testing.T) {
	t.Parallel()
	n := New(testString)
	n.makeLower()
	if strings.ToLower(n.theString) != n.theString {
		t.Error("Unable to make string lower case")
	}
}

func TestremoveNonDigits(t *testing.T) {
	t.Parallel()
	n := New(testString2)
	n.removeNonWordChars()
	if n.theString != "897" {
		t.Error("Unable to remove digits from string")
	}
}

func TestremoveNonWordChars(t *testing.T) {
	t.Parallel()
	n := New(testString3)
	n.removeNonWordChars()
	if n.theString != "acDFgHiJLmoPqRTuVwXyZ" {
		t.Error("Unable to remove")
	}
}

func TesturlEncodeSpaces(t *testing.T) {
	t.Parallel()
}

func TestrandomSelect(t *testing.T) {
	t.Parallel()
}
