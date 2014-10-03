package formatifier

import (
	"testing"
	"strings"
)

func TestToPhone(t *testing.T) {
	//
}

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

func TestToISBN(t *testing.T) {
	//
}

func TestToMorseCode(t *testing.T) {
	//
}

func TestToPirateSpeak(t *testing.T) {
	//
}

func TestToIRSA(t *testing.T) {
	//
}

func TestToLeet(t *testing.T) {
	result, err := ToLeet("that is cool dude")
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(result, "7#47 15 kewl d00d") {
		t.Error("Didn't find the expected converted string")
	}
}
