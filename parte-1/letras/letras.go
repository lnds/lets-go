package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	args := os.Args[1:]
	var vowels, consonants, digits bool
	var files []string
	for i, arg := range args {
		switch arg {
		case "-v":
			vowels = true
		case "-c":
			consonants = true
		case "-d":
			digits = true
		default:
			if strings.HasPrefix(arg, "-v") {
				log.Fatal("argumento %s inválido", args[i])
			} else {
				files = append(files, arg)
			}
		}
	}

	processFiles(files, vowels, consonants, digits)
}

func processFiles(files []string, vowels, consonants, digits bool) {

	var letters = map[rune]int{}
	var countVowels, countConsonants, countDigits int

	for _, file := range files {
		buf, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		runes := string(buf)
		for _, r := range runes {

			if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				continue
			}
			letters[r]++
			c := unicode.ToLower(r)
			switch c {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				countDigits++
			case 'a', 'e', 'i', 'o', 'u':
				countVowels++
			default:
				countConsonants++
			}
		}
	}

	fmt.Println("Letra|Cantidad")

	for key, val := range letters {
		fmt.Printf("%c | %d\n", key, val)
	}
	if vowels {
		fmt.Println("Vocales:", countVowels)
	}
	if consonants {
		fmt.Println("Consonantes:", countConsonants)
	}
	if digits {
		fmt.Println("Dígitos:", countDigits)
	}
}
