// Copyright 2015 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package formatifier

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ToPhone will format the provided string as a Phone Number.  Only supports
// US numbers currently.
func ToPhone(theString, delimiter string) (string, error) {
	f := New(theString)
	f.removeNonDigits()

	if f.length < 10 {
		return "", errors.New(lengthError)
	}

	var buffer bytes.Buffer
	count := 0
	switch f.length {
	case 10:
		buffer.WriteString("(")
		for _, i := range f.theString {
			count++
			buffer.WriteString(string(i))
			switch count {
			case 3:
				buffer.WriteString(") ")
			case 6:
				buffer.WriteString(fmt.Sprintf("%s", delimiter))
			}
		}
	case 11:
		for _, i := range f.theString {
			count++
			buffer.WriteString(string(i))
			switch count {
			case 1:
				buffer.WriteString(" (")
			case 4:
				buffer.WriteString(") ")
			case 7:
				buffer.WriteString(fmt.Sprintf("%s", delimiter))
			}
		}
	default:
		return "", errors.New("non US number given")
	}
	return buffer.String(), nil
}

// ToURL will format the provided string as a URL.  HTTP and HTTPS
// are the only supported protocols at this time.
func ToURL(theString, subdomain string, secure bool) (string, error) {
	f := New(theString)

	if f.length < 4 {
		return "", errors.New(lengthError)
	}

	f.makeLower()

	if secure {
		if len(subdomain) > 0 {
			return fmt.Sprintf("https://%s.%s", subdomain, f.theString), nil
		}
		return fmt.Sprintf("https://%s", f.theString), nil
	}

	if len(subdomain) > 0 {
		return fmt.Sprintf("http://%s.%s", subdomain, f.theString), nil
	}
	return fmt.Sprintf("http://%s", f.theString), nil
}

// ToSSN will format the provided string as a SSN.
func ToSSN(theString, delimiter string) (string, error) {
	f := New(theString)
	f.removeNonDigits()

	if f.length != 9 {
		return "", errors.New("string needs to be 9 digits for Social Security Numbers")
	}

	var buffer bytes.Buffer
	count := 0

	for _, i := range f.theString {
		count++

		buffer.WriteString(string(i))

		if count == 3 || count == 5 {
			buffer.WriteString(delimiter)
		}
	}
	return buffer.String(), nil
}

// ToLockCombo will format the provided string as a Lock Combo.
func ToLockCombo(theString, delimiter string) (string, error) {
	f := New(theString)
	f.removeNonDigits()

	if f.length != 6 {
		return "", errors.New("string needs to be 6 digits for Lock Combo format")
	}

	var buffer bytes.Buffer
	count := 0

	for _, i := range f.theString {
		count++

		buffer.WriteString(string(i))

		if count == 2 || count == 4 {
			buffer.WriteString(delimiter)
		}
	}
	return buffer.String(), nil
}

// ToISBN will format the provided string in International Standard Book Number
// (ISBN) format.
func ToISBN(theString, delimiter string) (string, error) {
	f := New(theString)
	f.removeNonDigits()

	if f.length != 13 {
		return "", errors.New("string must be 13 characters")
	}

	var buffer bytes.Buffer
	count := 0

	for _, i := range f.theString {
		count++

		buffer.WriteString(string(i))

		if count == 3 || count == 4 || count == 6 || count == 12 {
			buffer.WriteString(delimiter)
		}
	}
	return buffer.String(), nil
}

// ToMorseCode will format the provided string in Morse Code.
func ToMorseCode(theString string) (string, error) {
	f := New(theString)

	if f.length < 1 {
		return "", errors.New(lengthError)
	}

	f.makeLower()

	var buffer bytes.Buffer

	for _, i := range f.theString {
		key := string(i)
		if _, ok := morse[key]; ok {
			buffer.WriteString(morse[key])
		} else if key == " " {
			buffer.WriteString(" ")
		}
	}
	return buffer.String(), nil
}

// ToPirateSpeak will format the provided string in Pirate Speak.
func ToPirateSpeak(theString string) (string, error) {
	f := New(theString)

	if f.length < 1 {
		return "", errors.New(lengthError)
	}

	f.urlEncodeSpaces()

	response, err := http.Get(fmt.Sprintf(pirateLink, f.theString))
	if err != nil {
		return "", errors.New("unable to convert")
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

// ToIRSA will format the provided string in IRSA.
// International Radio-Telephony Spelling Alphabet
func ToIRSA(theString string) (string, error) {
	f := New(theString)

	if f.length < 1 {
		return "", errors.New(lengthError)
	}

	f.makeLower()

	var buffer bytes.Buffer

	for _, i := range f.theString {
		key := strings.ToLower(string(i))
		if _, ok := irsa[key]; ok {
			buffer.WriteString(irsa[key] + " ")
		}
	}
	return buffer.String(), nil
}

// ToLeet will format the provided string in Leet Speak.
// TODO(briandowns) Make select post match random.
func ToLeet(theString string) (string, error) {
	f := New(theString)

	if f.length < 1 {
		return "", errors.New(lengthError)
	}

	f.makeLower()

	words := strings.Fields(f.theString)
	var buffer bytes.Buffer

	for _, word := range words {
		key := string(word)
		if _, ok := leet[word]; ok {
			buffer.WriteString(leet[key][0] + " ")
		} else {
			for _, i := range word {
				letter := string(i)
				if _, ok := leet[letter]; ok {
					buffer.WriteString(leet[letter][0])
				}
			}
			buffer.WriteString(" ")
		}
	}
	return buffer.String(), nil
}
