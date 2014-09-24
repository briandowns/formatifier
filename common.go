package formatifier

import (
	"regexp"
	"strings"
)

type Formatifier struct {
	theString string
}

// Create a new instance of the String object.
func New() *Formatifier {
	return &Formatifier{}
}

// Makes the user entered string lower case.
func (f *Formatifier) makeLower() {
	f.theString = strings.ToLower(f.theString)
}

// Remove any non digit characters from the string.
func (f *Formatifier) removeNonDigits() {
	rp := regexp.MustCompile(`\D`)
	f.theString = rp.ReplaceAllString(f.theString, "")
}

// Remove all non word characters.
func (f *Formatifier) removeNonWordChars() {
	if len(f.theString) > 0 {
		rp := regexp.MustCompile(`\W|\s|_`)
		f.theString = rp.ReplaceAllString(f.theString, "")
	}
}

// Leet speak map of string slices
var leet = map[string]string{
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
