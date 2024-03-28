package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
        // Prompt the user to input a pattern
		if isPortuguese {
			color.New(color.FgMagenta).Println("Idioma: Português")
		} else {
			color.New(color.FgMagenta).Println("Language: English")
		}
		color.New(color.FgHiMagenta).Println("Escreva o padrão (ex. c_s_), 't' para trocar o idioma ou escreva 'q' para sair: ")
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
			printMatchingWordsGrid(words_pt, regexPattern)
		} else {
			printMatchingWordsGrid(words_en, regexPattern)
		}
    }
}

// readWordsFromFile reads words from a file and returns them as a slice of strings.
func readWordsFromFile(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var words []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words = append(words, scanner.Text())
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

// findMatchingWords finds words from the given list that match the given regular expression pattern.
func findMatchingWords(words []string, pattern string) []string {
    var matchedWords []string
    regex := regexp.MustCompile(pattern)
    for _, word := range words {
        if regex.MatchString(word) {
            matchedWords = append(matchedWords, word)
        }
    }
    return matchedWords
}

// printMatchingWordsGrid finds words from the given list that match the given regular expression pattern
// and prints them in a colorful grid.
func printMatchingWordsGrid(words []string, pattern string) {
    matchedWords := findMatchingWords(words, pattern)
    if len(matchedWords) == 0 {
        fmt.Println("Nenhuma palavra encontrada.")
        return
    }

	if len(matchedWords) > 300 {
		color.New(color.FgHiRed).Println("A lista completa não pode ser exibida")
        matchedWords = matchedWords[:300]
    }

    // Determine the number of columns in the grid
    numCols := 3
    numRows := (len(matchedWords) + numCols - 1) / numCols

    // Print the matched words in a colorful grid
    for i := 0; i < numRows; i++ {
        for j := 0; j < numCols; j++ {
            index := i*numCols + j
            if index < len(matchedWords) {
                // Colorize the word
                color.New(color.FgHiGreen).Printf("%-22s", matchedWords[index])
            }
        }
        fmt.Println()
    }
}
