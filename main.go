package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/fatih/color"
)

func main() {
    // Read words from file
    words_pt, err := readWordsFromFile("words_pt.txt")
    if err != nil {
        fmt.Println("Error reading words:", err)
        return
    }

	words_en, err := readWordsFromFile("words_en.txt")
    if err != nil {
        fmt.Println("Error reading words:", err)
        return
    }

	// Main loop
    scanner := bufio.NewScanner(os.Stdin)

	isPortuguese := true
    for {
		c := color.New(color.FgMagenta)
		chi := color.New(color.FgHiMagenta)
        // Prompt the user to input a pattern
		if isPortuguese {
			c.Println("Idioma: Português")
		} else {
			c.Println("Idioma: Inglês")
		}
		chi.Println("Escreva o padrão (ex. c_s_), 't' para trocar o idioma ou escreva 'q' para sair: ")
        scanner.Scan()
        input := scanner.Text()

        // Check if the user wants to quit
        if input == "q" {
            fmt.Println("Saindo...")
            return
        }

		if input == "t" {
            isPortuguese = !isPortuguese
			continue
        }

        // Compile the regular expression pattern
        regexPattern := convertPatternToRegex(input)

        // Find and print matching words in a colorful grid
		if isPortuguese {
			matchedWords := findMatchingWords(words_pt, regexPattern)
			displayWordsByLength(matchedWords)
		} else {
			matchedWords := findMatchingWords(words_en, regexPattern)
			displayWordsByLength(matchedWords)
		}
		fmt.Println("--------------------------------------------------------------------------------")
    }
}

// readWordsFromFile reads words from a file and returns them as a slice of strings.
func readWordsFromFile(filename string) (map[int][]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    words := make(map[int][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		wordLen := len(word)
		words[wordLen] = append(words[wordLen], word)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

    return words, nil
}

// convertPatternToRegex converts a pattern string into a regular expression pattern.
// For example, "_ l _ _ t" becomes "^.l..t$".
func convertPatternToRegex(pattern string) string {
	// Replace underscores with dots (any character)
	regexPattern := strings.ReplaceAll(pattern, "_", ".")

	// Replace '*' with '.*' to match any number of characters
	regexPattern = strings.ReplaceAll(regexPattern, "*", ".*")

	// Add anchors to match the whole string
	regexPattern = "^" + regexPattern + "$"

	return regexPattern
}

// findMatchingWords finds words from the given map of word lengths that match the given regular expression pattern.
func findMatchingWords(words map[int][]string, pattern string) []string {
	var matchedWords []string
	regex := regexp.MustCompile(pattern)
	for _, wordList := range words {
		for _, word := range wordList {
			if regex.MatchString(word) {
				matchedWords = append(matchedWords, word)
			}
		}
	}
	return matchedWords
}

// displayWordsByLength divides words by length and displays each list separately.
func displayWordsByLength(words []string) {
	if(len(words) == 0) {
		color.New(color.FgRed).Println("Nenhuma palavra encontrada.")
		return
	}
    // Map to store words by length
    wordsByLength := make(map[int][]string)

    // Group words by length
    for _, word := range words {
        wordLen := len(word)
        wordsByLength[wordLen] = append(wordsByLength[wordLen], word)
    }

	var lengths []int
    for length := range wordsByLength {
        lengths = append(lengths, length)
    }
    sort.Ints(lengths)

	cgreen := color.New(color.FgGreen)
	cyellow := color.New(color.FgYellow)

	for _, length := range lengths {
        wordList := wordsByLength[length]

        // Print header with colored length
        cyellow.Printf("Palavras de tamanho %d:\n", length)

        // Print words
        for i, word := range wordList {
            if i%3 == 0 {
				cgreen.Printf("%-25s", word)
			} else if i%3 == 1 {
				cgreen.Printf("%-25s", word)
			} else {
				cgreen.Printf("%s\n", word)
			}
        }

		// Print new line for odd number of words
        if len(wordList)%3 != 0 {
            cgreen.Println()
        }
    }
}