package data

import (
	"encoding/csv"
	"os"
	"strings"
)

func ReadCSV(filepath string) ([][]string, error) {
	fileContents, err := readFile(filepath)
	if err != nil {
		return [][]string{}, err
	}

	records, err := stringToCSV(fileContents)
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func stringToCSV(s string) ([][]string, error) {
	reader := csv.NewReader(strings.NewReader(s))
	reader.Comma = ',' // Delimeter

	records, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func readFile(filepath string) (string, error) {
	b := make([]byte, 100_000)
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}

	n, err := f.Read(b)
	if err != nil {
		return "", err
	}

	// string of bytes read in
	return string(b[:n]), nil
}
