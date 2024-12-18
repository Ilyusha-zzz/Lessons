package main

import "fmt"

func TruncateByteString() {
	var text string
	var width int
	fmt.Println("Enter line and width\nformat: text width")
	fmt.Scanf("%s %d", &text, &width)

	maxWidth := width
    if width > len(text) {
		maxWidth = len(text)
	}

	var truncated string
	for i := 0; i < maxWidth; i++ {
		truncated += string(text[i])
	}
	if len(text) > width {
		truncated += "..."
	}

	fmt.Printf("Truncated text: %v\n", truncated)
}