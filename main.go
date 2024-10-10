package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/globalsign/publicsuffix"
)

func extractTLD(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", fmt.Errorf("invalid input: empty string")
	}

	result, err := publicsuffix.EffectiveTLDPlusOne(input)
	if err != nil {
		return "", fmt.Errorf("processing error: %w", err)
	}

	return result, nil
}

func processInputStream(r *bufio.Reader) ([]string, error) {
	var tlds []string
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("input read error: %w", err)
		}

		output, err := extractTLD(line)
		if err != nil {
			log.Printf("Warning: %v", err)
			continue
		}

		tlds = append(tlds, output)
	}
	return tlds, nil
}

func removeDuplicatesAndSort(tlds []string) []string {
	uniqueTLDs := make(map[string]bool)
	var result []string

	for _, tld := range tlds {
		lowerTLD := strings.ToLower(tld)
		if !uniqueTLDs[lowerTLD] {
			uniqueTLDs[lowerTLD] = true
			result = append(result, tld)
		}
	}

	sort.Strings(result) // Sort the results in lexicographical order
	return result
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	tlds, err := processInputStream(inputReader)
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}

	sortedUniqueTLDs := removeDuplicatesAndSort(tlds)

	for _, tld := range sortedUniqueTLDs {
		fmt.Println(tld)
	}
}
