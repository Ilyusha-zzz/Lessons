package main

import "fmt"

func LanguageByCodeSwitch() {
	var code string
	fmt.Println("Enter language code")
	fmt.Scan(&code)

	switch code {
	case "ru":
		fmt.Println("Russian")
	case "rus":
		fmt.Println("Russian")
	case "en":
		fmt.Println("English")
	case "fr":
		fmt.Println("French")
	default:
		fmt.Println("Unknown language")
	}
}