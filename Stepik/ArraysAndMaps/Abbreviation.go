package main

import (
	"fmt"
	"strings"
	"unicode"
	"bufio"
	"os"
)

func Abbreviation() {
	var abbreviation string
	var phrase string
	fmt.Println("Enter phrase")
	phrase = readString()

	for _, v := range strings.Split(phrase, " ") {
		if unicode.IsLetter(rune(v[0])) {
			abbreviation += strings.ToUpper(string(strings.Split(v, "")[0]))
		}
	}
	
	fmt.Printf("abbreviation: %v\n", abbreviation)
}

func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}