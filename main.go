package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

func processInputStream(r *bufio.Reader) error {
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("input read error: %w", err)
		}

		output, err := extractTLD(line)
		if err != nil {
			log.Printf("Warning: %v", err)
			continue
		}

		fmt.Println(output)
	}
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	if err := processInputStream(inputReader); err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}
