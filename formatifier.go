// Copyright 2014 Brian J. Downs
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

//
// formatifier is a library to easily format strings in a user defined and predefined manner.
//

package formatifier

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Format the provided string as a Phone Number.  Only supports
// US numbers currently.
func ToPhone(theString string, delimiter string) (string, error) {
	f := New(theString)

	f.removeNonDigits()

	var buffer bytes.Buffer
	count := 0

	if len(f.theString) == 10 {
		buffer.WriteString("(")
		for _, i := range f.theString {
			count++

			buffer.WriteString(string(i))

			if count == 3 {
				buffer.WriteString(") ")
			} else if count == 6 {
				buffer.WriteString(fmt.Sprintf("%s", delimiter))
			}
		}
	} else if len(f.theString) == 11 {
		for _, i := range f.theString {
			count++

			buffer.WriteString(string(i))

			if count == 1 {
				buffer.WriteString(" (")
			} else if count == 4 {
				buffer.WriteString(") ")
			} else if count == 7 {
				buffer.WriteString(fmt.Sprintf("%s", delimiter))
			}
		}
	}
	return buffer.String(), nil
}

// Format the provided string as a URL.  HTTP and HTTPS
// are the only supported protocols at this time.
func ToURL(theString string, secure bool, subdomain string) (string, error) {
	f := New(theString)

	if len(f.theString) < 1 {
		return "", errors.New(lengthError)
	}

	f.makeLower()

	if secure {
		if len(subdomain) > 0 {
			return fmt.Sprintf("https://%s.%s", subdomain, f.theString), nil
		} else {
			return fmt.Sprintf("https://%s", f.theString), nil
		}
	} else {
		if len(subdomain) > 0 {
			return fmt.Sprintf("http://%s.%s", subdomain, f.theString), nil
		} else {
			return fmt.Sprintf("http://%s", f.theString), nil
		}
	}
}

// Format the provided string as a SSN.
func ToSSN(theString string, delimiter string) (string, error) {
	f := New(theString)

	f.removeNonDigits()

	if len(f.theString) != 9 {
		return "", errors.New("ERROR: String needs to be 9 digits for Social Security Numbers.")
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

// Format the provided string as a Lock Combo.
func ToLockCombo(theString string, delimiter string) (string, error) {
	f := New(theString)

	f.removeNonDigits()

	if len(f.theString) != 9 {
		return "", errors.New("ERROR: String needs to be 6 digits for Lock Combo format.")
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

// Format the provided string in International Standard Book Number
// (ISBN) format.
func ToISBN(theString string, delimiter string) (string, error) {
	f := New(theString)

	f.removeNonDigits()

	if len(f.theString) != 13 {
		return "", errors.New("ERROR: string must be 13 characters.")
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

// Format the provided string in Morse Code.
func ToMorseCode(theString string) (string, error) {
	f := New(theString)

	if len(f.theString) < 1 {
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

// Format the provided string in Pirate Speak.
func ToPirateSpeak(theString string) (string, error) {
	f := New(theString)

	if len(f.theString) < 1 {
		return "", errors.New(lengthError)
	}

	f.urlEncodeSpaces()

	response, err := http.Get(fmt.Sprintf(pirateLink, f.theString))
	if err != nil {
		return "", errors.New("ERROR: Unable to convert.")
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return string(contents), nil
	}
	return "", nil
}

// Format the provided string in IRSA.
// International Radio-Telephony Spelling Alphabet
func ToIRSA(theString string) (string, error) {
	f := New(theString)

	if len(f.theString) < 1 {
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

// Format the provided string in Leet Speak.
// TODO(bdowns328) Make select post match random.
func ToLeet(theString string) (string, error) {
	f := New(theString)

	if len(f.theString) < 1 {
		return "", errors.New(lengthError)
	}

	f.makeLower()
	rand.Seed(time.Now().UTC().UnixNano())

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
