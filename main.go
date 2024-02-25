package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Print("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}
	font, _ := os.ReadFile(os.Args[2] + ".txt")
	read_line := string(os.Args[1])
	read_line = strings.TrimRight(read_line, "\\n")
	lines := strings.Split(string(read_line), "\\n")
	fmt.Println(lines)
	if string(os.Args[1]) == "" {
		return
	}
	strfont := string(font)
	alphabet := strings.Split(strfont, "\n")
	for _, line := range lines {
		for i := 1; i < 9; i++ {
			for _, char := range line {
				os.Stdout.WriteString(alphabet[(char-32)*9+rune(i)])
			}
			os.Stdout.WriteString("\n")
		}
	}
}
