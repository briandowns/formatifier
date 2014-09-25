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

/*
 formatifier is a library to easily format strings in a user defined and predefined manner.
*/

package formatifier

import (
	"regexp"
	"strings"
)

const (
	lengthError = "ERROR: String not long enough to convert."
	pirateLink  = "http://www.isithackday.com/arrpi.php?text=%s"
)

// Type to hold user input string.
type formatifier struct {
	theString string
}

// Create a new instance of the String object.
func New(s string) *formatifier {
	return &formatifier{theString: s}
}

// Makes the user entered string lower case.
func (f *formatifier) makeLower() {
	f.theString = strings.ToLower(f.theString)
}

// Remove any non digit characters from the string.
func (f *formatifier) removeNonDigits() {
	rp := regexp.MustCompile(`\D`)
	f.theString = rp.ReplaceAllString(f.theString, "")
}

// Remove all non word characters.
func (f *formatifier) removeNonWordChars() {
	if len(f.theString) > 0 {
		rp := regexp.MustCompile(`\W|\s|_`)
		f.theString = rp.ReplaceAllString(f.theString, "")
	}
}

func (f *formatifier) urlEncodeSpaces() {
	rp := regexp.MustCompile(`\s`)
	f.theString = rp.ReplaceAllString(f.theString, "%20")
}

// Leet speak map of string slices
var leet = map[string][]string{
	"leet":     []string{"1337"},
	"the":      []string{"teh"},
	"cool":     []string{"kewl"},
	"dude":     []string{"d00d"},
	"you":      []string{"u"},
	"noob":     []string{"n00b"},
	"noobs":    []string{"n00bs"},
	"own":      []string{"pwn"},
	"owned":    []string{"pwned"},
	"rocks":    []string{"roxx0rs"},
	"exploits": []string{"sploitz"},
	"woot":     []string{"w00t"},
	"hacker":   []string{"hax0r"},
	"hackers":  []string{"hax0rz"},
	"a":        []string{"4", "@"},
	"b":        []string{"8", "]3", "]8", "|3", "|8", "13"},
	"c":        []string{"(", "{"},
	"d":        []string{")", "[}", "|)", "|}", "|>"},
	"e":        []string{"3"},
	"f":        []string{"|=", "ph"},
	"g":        []string{"6", "9", "&"},
	"h":        []string{"#", "|-|"},
	"i":        []string{"1", "!", "|"},
	"j":        []string{"_|", "u|"},
	"k":        []string{"|<", "|{"},
	"l":        []string{"|", "1", "|_"},
	"m":        []string{"/\\/\\", "|\\/|"},
	"n":        []string{"/\\/", "|\\|"},
	"o":        []string{"0", "()"},
	"p":        []string{"|D", "|*"},
	"q":        []string{"(,)", "O\\", "[]\\"},
	"r":        []string{"|2", "|?", "][2"},
	"s":        []string{"5", "$"},
	"t":        []string{"7", "+"},
	"u":        []string{"(_)", "|_|"},
	"v":        []string{"\\/", "\\\\//"},
	"w":        []string{"\\/\\/", "|/\\|", "VV"},
	"x":        []string{"><", "}{"},
	"y":        []string{"'/", "%"},
	"z":        []string{"2", "7_"},
}

var irsa = map[string]string{
	" ": " | ",
	"a": "alfa",
	"b": "bravo",
	"c": "charlie",
	"d": "delta",
	"e": "echo",
	"f": "foxtrot",
	"g": "golf",
	"h": "hotel",
	"i": "india",
	"j": "juliet",
	"k": "kilo",
	"l": "lima",
	"m": "mike",
	"n": "november",
	"o": "oscar",
	"p": "papa",
	"q": "quebec",
	"r": "romeo",
	"s": "sierra",
	"t": "tango",
	"u": "uniform",
	"v": "victor",
	"w": "whiskey",
	"x": "x-ray",
	"y": "yankee",
	"z": "zulu",
}

var morse = map[string]string{
	"a":  ". _",
	"b":  "_ . . .",
	"c":  "_ . _ .",
	"d":  "_ . .",
	"e":  ".",
	"f":  ". . _ .",
	"g":  "_ _ .",
	"h":  ". . . .",
	"i":  ". .",
	"j":  ". _ _ _",
	"k":  "_ . _",
	"l":  ". _ . .",
	"m":  "_ _",
	"n":  "_ .",
	"o":  "_ _ _",
	"p":  ". _ _ .",
	"q":  "_ _ . _",
	"r":  ". _ .",
	"s":  ". . .",
	"t":  "_",
	"u":  ". . _",
	"v":  ". . . _",
	"w":  ". _ _",
	"x":  "_ . . _",
	"y":  "_ . _ _",
	"z":  "_ _ . .",
	"0":  "_ _ _ _ _",
	"1":  ". _ _ _ _",
	"2":  ". . _ _ _",
	"3":  ". . . _ _",
	"4":  ". . . . _",
	"5":  ". . . . .",
	"6":  "_ . . . .",
	"7":  "_ _ . . .",
	"8":  "_ _ _ . .",
	"9":  "_ _ _ _ .",
	".":  "·–·–·– ",
	",":  "––··–– ",
	"?":  "··––·· ",
	"'":  "·––––· ",
	"!":  "–·–·–– ",
	"/":  "–··–· ",
	"(":  "–·––· ",
	")":  "–·––·– ",
	"&":  "·–··· ",
	":":  "–––··· ",
	";":  "–·–·–· ",
	"=":  "–···– ",
	"+":  "·–·–· ",
	"–":  "–····– ",
	"_":  "··––·– ",
	"\"": "·–··–· ",
	"$":  "···–··– ",
	"@":  "·––·–· ",
}
