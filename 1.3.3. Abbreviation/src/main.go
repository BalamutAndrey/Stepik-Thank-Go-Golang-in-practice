package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var abbr []rune
	phrase := readString()
	arrWords := strings.Fields(phrase)
	for wordNumb := range arrWords {
		currWord := []rune(arrWords[wordNumb])
		if unicode.IsLetter(currWord[0]) {
			abbr = append(abbr, unicode.ToUpper(currWord[0]))
		}
	}
	fmt.Println(string(abbr))
}

func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
