package csvutil

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/htanivar/rutils/file"
)

// CSVToRecords reads a CSV file and converts it to a slice of map records.
func CSVToRecords(path string) ([]map[string]string, error) {
	if err := file.ValidateType(".csv", path); err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1

	headers, err := r.Read()
	if err != nil {
		return nil, err
	}

	var records []map[string]string

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		record := make(map[string]string, len(headers))
		for i, h := range headers {
			if i < len(row) {
				record[h] = row[i]
			} else {
				record[h] = ""
			}
		}
		records = append(records, record)
	}

	return records, nil
}
