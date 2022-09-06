//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func main() {
	type runeCountFunc func(rune) bool
	countSymbolsFunc := func(str string, f runeCountFunc) int {
		c := 0
		for _, r := range []rune(str) {
			if f(r) {
				c++
			}
		}
		return c
	}

	const (
		Letters = iota
		Numbers
		Spaces
		Punct
	)

	runeCountFunctions := make(map[int]func(rune) bool)
	runeCountFunctions[Letters] = func(r rune) bool { return unicode.IsLetter(r) }
	runeCountFunctions[Numbers] = func(r rune) bool { return unicode.IsNumber(r) }
	runeCountFunctions[Spaces] = func(r rune) bool { return unicode.IsSpace(r) }
	runeCountFunctions[Punct] = func(r rune) bool { return unicode.IsPunct(r) }

	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	lineC, numC, spaceC, punctC := 0, 0, 0, 0

	for _, l := range lines {
		lineC += countSymbolsFunc(l, runeCountFunctions[Letters])
		numC += countSymbolsFunc(l, runeCountFunctions[Numbers])
		spaceC += countSymbolsFunc(l, runeCountFunctions[Spaces])
		punctC += countSymbolsFunc(l, runeCountFunctions[Punct])
	}
	fmt.Printf("total counts \n")
	fmt.Printf(" _________________________________________\n")
	fmt.Printf("|Letters | Numbers | Spaces | Punctuation | \n")
	fmt.Printf("|   %v   |    %v   |  %v    |     %v      | \n", lineC, numC, spaceC, punctC)
	fmt.Printf("-------------------------------------------\n")

}
