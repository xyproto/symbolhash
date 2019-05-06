package symbolhash

// Package for representing a string as a unicode hash of a given size.
// Less entropy is represented as more repetition.
// Useful for listing sensitive data like passwords, without revealing the actual data,
// for instance in a control panel for administrators.

import (
	"strings"
	"unicode/utf8"
)

type SymbolHash struct {
	data string
	size int
}

// Create a new SymbolHash, takes a string and a length
func New(data string, size int) *SymbolHash {
	return &SymbolHash{data, size}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Returns a unicode representation hash
func (hash *SymbolHash) String() string {
	var (
		symbols       = []string{"➳", "⚔", "♛", "⚓", "✉", "✪", "✏", "טּ", "٩", "ღ", "•", "▰", "◉", "◑", "◔", "◕", "☃", "☺", "♡", "♥", "♾", "✌", "❀", "❂", "⨂", "ツ", "Ꙩ", "ꙩ", "Ꙫ", "ꙫ", "ꙮ", "מּ", "תּ", "❤", "☀", "☆", "☂", "☻", "♞", "☯", "☭", "☢", "→", "☎", "❄", "♫", "✂", "▷", "✇", "♎", "⇧", "☮", "♻", "⌘", "☘", "⚡", "★", "⚘", "♩", "☕", "☁", "☄", "⍨", "✈"}
		num, sparenum uint8
		sparecounter  int
		repr          string
	)
	for _, b := range hash.data {
		num = uint8(b)
		num = num << 2 // drop the two first bits
		num = num >> 2
		twolastbits := uint8(b) - num // get the two first bits
		if (num > 0) && (int(num) < len(symbols)) {
			// if the symbol was already added, don't add it
			if !strings.Contains(repr, symbols[num]) {
				repr += symbols[num]
			}
		}
		if sparecounter == 0 {
			sparenum = twolastbits
			sparecounter++
		} else {
			sparenum = sparenum << 2 // shift the bits 2 places up
			sparenum += twolastbits  // add the leftover bits
		}
		if sparecounter == 4 {
			// have shifted the bits four times now, time to use it
			if (sparenum > 0) && (int(sparenum) < len(symbols)) {
				// if the symbol was already added, don't add it
				if !strings.Contains(repr, symbols[sparenum]) {
					repr += symbols[sparenum]
				}
			}
			// reset the counter and placeholder
			sparecounter = 0
			sparenum = 0
		}
	}

	// If it's too short, repeat it
	retval := ""
	for utf8.RuneCountInString(retval) < hash.size {
		retval += repr
	}

	// If it's too long, cut it
	if utf8.RuneCountInString(retval) > hash.size {
		retval = retval[:min(len(retval), hash.size)]
	}

	return retval
}
