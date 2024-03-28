package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

func main() {
    // Read words from file
    words, err := readWordsFromFile("words.txt")
    if err != nil {
        fmt.Println("Error reading words:", err)
        return
    }

    // Prompt the user to input a pattern
    fmt.Print("Escreva o padrão (e.g., p_an__): ")
    var pattern string
    fmt.Scanln(&pattern)

    // Compile the regular expression pattern
    regexPattern := convertPatternToRegex(pattern)

    // Find and print matching words
    matchedWords := findMatchingWords(words, regexPattern)
    fmt.Println("Palavras:", matchedWords)
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
