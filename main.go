package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/globalsign/publicsuffix"
)

func extractTLD(input string) (string, string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", "", fmt.Errorf("invalid input: empty string")
	}

	rootDomain, err := publicsuffix.EffectiveTLDPlusOne(input)
	if err != nil {
		return "", "", fmt.Errorf("processing error: %w", err)
	}

	return rootDomain, input, nil
}

func processInputStream(r *bufio.Reader) (map[string][]string, error) {
	domainMap := make(map[string][]string)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("input read error: %w", err)
		}

		rootDomain, fullDomain, err := extractTLD(line)
		if err != nil {
			log.Printf("Warning: %v", err)
			continue
		}

		// Append the full domain (subdomains and root domain) under the root domain
		domainMap[rootDomain] = append(domainMap[rootDomain], fullDomain)
	}
	return domainMap, nil
}

func removeDuplicatesAndSort(domainMap map[string][]string) map[string][]string {
	for rootDomain, domains := range domainMap {
		uniqueDomains := make(map[string]bool)
		var uniqueList []string

		// Remove duplicates
		for _, domain := range domains {
			if !uniqueDomains[domain] {
				uniqueDomains[domain] = true
				uniqueList = append(uniqueList, domain)
			}
		}

		// Sort the domains
		sort.Strings(uniqueList)
		domainMap[rootDomain] = uniqueList
	}

	return domainMap
}

func writeToFiles(domainMap map[string][]string) error {
	for rootDomain, domains := range domainMap {
		// Create the filename based on the root domain
		filename := fmt.Sprintf("%s.txt", rootDomain)

		// Open or create the file
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("file open error: %w", err)
		}
		defer file.Close()

		// Write all the domains (including subdomains) under the respective root domain
		for _, domain := range domains {
			if _, err := file.WriteString(domain + "\n"); err != nil {
				return fmt.Errorf("file write error: %w", err)
			}
		}
	}

	return nil
}

func main() {
	// Parse command-line flags
	createFiles := flag.Bool("n", false, "create .txt files for each root domain and save domains")
	flag.Parse()

	inputReader := bufio.NewReader(os.Stdin)

	// Process input stream
	domainMap, err := processInputStream(inputReader)
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}

	// Remove duplicates and sort the domains under each root domain
	sortedDomainMap := removeDuplicatesAndSort(domainMap)

	// Print the sorted, unique list for each domain
	for rootDomain, domains := range sortedDomainMap {
		fmt.Printf("%s: %v\n", rootDomain, domains)
	}

	// If -n flag is provided, write the TLDs and subdomains to respective .txt files
	if *createFiles {
		if err := writeToFiles(sortedDomainMap); err != nil {
			log.Fatalf("Error writing to files: %v", err)
		}
	}
}
