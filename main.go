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
    words, err := readWordsFromFile("words.txt")
    if err != nil {
        fmt.Println("Error reading words:", err)
        return
    }

	// Main loop
    scanner := bufio.NewScanner(os.Stdin)
    for {
          // Prompt the user to input a pattern
    fmt.Print("Escreva o padrão (ex. c_s_) ou escreva 'q' para sair: ")
        scanner.Scan()
        input := scanner.Text()

        // Check if the user wants to quit
        if input == "q" {
            fmt.Println("Saindo...")
            return
        }

        // Compile the regular expression pattern
        regexPattern := convertPatternToRegex(input)

        // Find and print matching words in a colorful grid
    	printMatchingWordsGrid(words, regexPattern)
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
    pattern = strings.ReplaceAll(pattern, "_", ".") // Replace underscores with dots (any character)
    return "^" + pattern + "$"                      // Add anchors to match the whole word
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

	if len(matchedWords) > 100 {
		color.New(color.FgHiRed).Println("A lista completa não pode ser exibida")
        matchedWords = matchedWords[:100]
    }

    // Determine the number of columns in the grid
    numCols := 5
    numRows := (len(matchedWords) + numCols - 1) / numCols

    // Print the matched words in a colorful grid
    for i := 0; i < numRows; i++ {
        for j := 0; j < numCols; j++ {
            index := i*numCols + j
            if index < len(matchedWords) {
                // Colorize the word
                color.New(color.FgHiGreen).Printf("%-15s", matchedWords[index])
            }
        }
        fmt.Println()
    }
}
