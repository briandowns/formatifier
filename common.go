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

// Package formatifier is a library to easily format strings in a user defined
// and predefined manner.
package formatifier

import (
	"crypto/rand"
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
	length    int
}

// New will create an instance of the String object.
func New(s string) *formatifier { return &formatifier{theString: s, length: len(s)} }

// makeLower will turn the user entered string to lower case
func (f *formatifier) makeLower() { f.theString = strings.ToLower(f.theString) }

// removeNonDigits removes any non digit characters from the string.
func (f *formatifier) removeNonDigits() {
	rp := regexp.MustCompile(`\D`)
	f.theString = rp.ReplaceAllString(f.theString, "")
}

// removeNonWordChars removes all non word characters.
func (f *formatifier) removeNonWordChars() {
	if len(f.theString) > 0 {
		rp := regexp.MustCompile(`\W|\s|_`)
		f.theString = rp.ReplaceAllString(f.theString, "")
	}
}

// urlEncodeSpaces will replace spaces with "%20"'s
func (f *formatifier) urlEncodeSpaces() {
	rp := regexp.MustCompile(`\s`)
	f.theString = rp.ReplaceAllString(f.theString, "%20")
}

// random select will return a random selection from an int slice
func randomSelect(a []int) int {
	var tmpIndex int
	length := len(a)
	randBytes := make([]byte, length)
	if _, err := rand.Read(randBytes); err == nil {
		tmpIndex = int(randBytes[0]) % length
	}
	return a[tmpIndex]
}

// leet speak map of string slices
var leet = map[string][]string{
	"leet":     {"1337"},
	"the":      {"teh"},
	"cool":     {"kewl"},
	"dude":     {"d00d"},
	"you":      {"u"},
	"noob":     {"n00b"},
	"noobs":    {"n00bs"},
	"own":      {"pwn"},
	"owned":    {"pwned"},
	"rocks":    {"roxx0rs"},
	"exploits": {"sploitz"},
	"woot":     {"w00t"},
	"hacker":   {"hax0r"},
	"hackers":  {"hax0rz"},
	"a":        {"4", "@"},
	"b":        {"8", "]3", "]8", "|3", "|8", "13"},
	"c":        {"(", "{"},
	"d":        {")", "[}", "|)", "|}", "|>"},
	"e":        {"3"},
	"f":        {"|=", "ph"},
	"g":        {"6", "9", "&"},
	"h":        {"#", "|-|"},
	"i":        {"1", "!", "|"},
	"j":        {"_|", "u|"},
	"k":        {"|<", "|{"},
	"l":        {"|", "1", "|_"},
	"m":        {"/\\/\\", "|\\/|"},
	"n":        {"/\\/", "|\\|"},
	"o":        {"0", "()"},
	"p":        {"|D", "|*"},
	"q":        {"(,)", "O\\", "[]\\"},
	"r":        {"|2", "|?", "][2"},
	"s":        {"5", "$"},
	"t":        {"7", "+"},
	"u":        {"(_)", "|_|"},
	"v":        {"\\/", "\\\\//"},
	"w":        {"\\/\\/", "|/\\|", "VV"},
	"x":        {"><", "}{"},
	"y":        {"'/", "%"},
	"z":        {"2", "7_"},
}

// irsa conversion map
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

// morese holds conversion chars for Morse Code
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
