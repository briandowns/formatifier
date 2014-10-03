package formatifier

import (
	"testing"
	"strings"
)

func TestToURL(t *testing.T) {
	result1, err1 := ToURL("github.com", false, "www")
	if err1 != nil {
		t.Error(err1)
	}

	if !strings.ContainsAny(result1, "http | www") {
		t.Error("Didn't find expected protocol")
	}

	result2, err2 := ToURL("github.com", true, "www")
	if err2 != nil {
		t.Error(err2)
	}

	if !strings.ContainsAny(result2, "https | gist") {
		t.Error("Didn't find expected protocol")
	}
}

func TestToSSN(t *testing.T) {
	result, err := ToSSN("986732987", "-")
	if err != nil {
		t.Error(err)
	}

	resultLength := len(result)

	if resultLength != 11 {
		t.Error("Expected an 8 char string and got: ", resultLength)
	}
}

func TestToLockCombo(t *testing.T) {
	result, err := ToLockCombo("651096", "-")
	if err != nil {
		t.Error(err)
	}

	resultLength := len(result)

	if resultLength != 8 {
		t.Error("Expected an 8 char string and got: ", resultLength)
	}
}

