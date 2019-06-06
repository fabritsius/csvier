package csvier

import (
	"encoding/csv"
	"io"
	"os"
)

type config struct {
	index     []string
	skip      int
	limit     int
	delimiter rune
}

// Read parses CSV data and returns a map for each line
func Read(csvData io.Reader, options ...func(*config) error) ([]map[string]string, error) {
	cfg := config{}
	// apply options to the config
	for _, option := range options {
		option(&cfg)
	}

	reader := csv.NewReader(csvData)
	reader.TrimLeadingSpace = true
	if cfg.delimiter != 0 {
		reader.Comma = cfg.delimiter
	}

	result := make([]map[string]string, 0)
	counter := 0
	stopAt := cfg.skip + cfg.limit
	// iterate over the lines
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		// if index isn`t set use first line as index
		if len(cfg.index) == 0 {
			cfg.index = line
			continue
		}

		counter++
		if counter > cfg.skip {
			// stop iteration if limit is reached
			if stopAt > 0 && counter > stopAt {
				break
			}
			// create a map of column:value pairs
			record := make(map[string]string)
			for i, name := range cfg.index {
				record[name] = line[i]
			}
			result = append(result, record)
		}
	}

	return result, nil
}

// ReadFile function reads CSV file and returns a map for each line
func ReadFile(csvFile string, options ...func(*config) error) ([]map[string]string, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}

	return Read(file, options...)
}

// Index is an option for Read() and allows to change csv index (column names)
func Index(items []string) func(*config) error {
	return func(cfg *config) error {
		cfg.index = items
		return nil
	}
}

// Skip is an option for Read() and allows to skip number of rows
func Skip(n int) func(*config) error {
	return func(cfg *config) error {
		cfg.skip = n
		return nil
	}
}

// Limit is an option for Read() and allows to limit number of rows in the result
func Limit(n int) func(*config) error {
	return func(cfg *config) error {
		cfg.limit = n
		return nil
	}
}

// Delimiter is an option for Read() and allows to change csv delimiter
func Delimiter(d rune) func(*config) error {
	return func(cfg *config) error {
		cfg.delimiter = d
		return nil
	}
}
