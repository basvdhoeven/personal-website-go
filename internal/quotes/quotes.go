package quotes

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

type QuoteRetriever struct {
	Quotes map[string][]string
}

func NewQuotesRetriever() *QuoteRetriever {
	return &QuoteRetriever{
		Quotes: make(map[string][]string),
	}
}

func (qr *QuoteRetriever) LoadQuotesFromTextFiles(folder string) error {
	files, err := os.ReadDir(folder)
	if err != nil {
		return err
	}

	for _, file := range files {
		// to do validate that file is a txt file
		fileName := file.Name()
		category := strings.TrimSuffix(fileName, filepath.Ext(fileName))

		textFile, err := os.Open(folder + "/" + fileName)
		if err != nil {
			return fmt.Errorf("error reading quotes from file: %v", err)
		}
		defer textFile.Close()

		var lines []string
		scanner := bufio.NewScanner(textFile)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if scanner.Err() != nil {
			return scanner.Err()
		}

		qr.Quotes[category] = lines
	}

	return nil
}

func (qr *QuoteRetriever) GetRandom(category string) (string, error) {
	quotes, ok := qr.Quotes[category]
	if !ok {
		return "", errors.New("invalid quote category")
	}

	return quotes[rand.Intn(len(quotes))], nil
}
