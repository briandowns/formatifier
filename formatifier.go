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
// Examples:
// Phone:
// URL: http://play.golang.org/p/W_hL358mUa
// SSN: http://play.golang.org/p/iMmK9SL_jS
// Lock: http://play.golang.org/p/_5wf5YyEUw
// ISBN: http://play.golang.org/p/k27ZRvCeo7
// Morse: http://play.golang.org/p/M_xtujXwwC
// Pirate: http://play.golang.org/p/QVF0PTx3cm
// Leet: http://play.golang.org/p/I6g21KKdjk
// IRSA: http://play.golang.org/p/_nRQwWywKW
//

package formatifier

import (
	"errors"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

// Format the provided string as a Phone Number.  Sticking to the ISO
// standard for country abbreviations.
/*
 China 1234 5678
 France  01–23–45–67–89
 Poland  (12) 345.67.89
 Singapore 123 4567
 Thailand  (01) 234–5678 or (012) 34–5678
 United Kingdom  0123 456 7890 or 01234 567890
 United States 1 (123) 456 7890
*/
func (f *Formatifier) ToPhone(geo string, international string, delimiter string) (string, error) {
	if len(geo) != 2 {
		return nil, errors.New("ERROR: Incorrect GEO code format")
	}

	f.removeNonDigits()

	return "", nil
}

// Format the provided string as a URL.  HTTP and HTTPS
// are the only supported protocols at this time.
func (f *Formatifier) ToURL(secure bool, subdomain string) (string, error) {
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
func (f *Formatifier) ToSSN(delimiter string) (string, error) {
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
func (f *Formatifier) ToLockCombo(delimiter string) (string, error) {
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
func (f *Formatifier) ToISBN(delimiter string) (string, error) {
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
func (f *Formatifier) ToMorseCode() (string, error) {
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
func (f *Formatifier) ToPirateSpeak() (string, error) {
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
func (f *Formatifier) ToIRSA() (string, error) {
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
func (f *Formatifier) ToLeet() (string, error) {
	if len(f.theString) < 1 {
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
