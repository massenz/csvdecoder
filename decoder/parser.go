package decoder

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

type Record = map[string]string
type Records = []Record

const (
	PipeSeparator           = '|'
	DefaultTrimLeadingSpace = true
)

func NewPipeReader(r io.Reader) *csv.Reader {
	c := csv.NewReader(r)
	c.Comma = PipeSeparator
	c.TrimLeadingSpace = DefaultTrimLeadingSpace
	return c
}

// ReadRecordsFromFile reads a file and returns a slice of records.
func ReadRecordsFromFile(filename string) (Records, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadPipeSeparatedLines(file)
}

// ReadPipeSeparatedLines generalizes to a Reader and returns a slice of records.
func ReadPipeSeparatedLines(psr io.Reader) (records Records, err error) {
	records = make(Records, 0)
	r := NewPipeReader(psr)

	// The first line is the header and contains all the field names.
	header, err := r.Read()
	// The csv reader does not trim trailing spaces, so we do it here.
	for i, s := range header {
		header[i] = strings.TrimSpace(s)
	}
	for {
		values, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return records, err
		}
		record, err := NewRecord(header, values)
		if err != nil {
			return records, err
		}
		records = append(records, record)
	}
	return records, nil
}

// NewRecord creates a new record from a header and a set of values.
func NewRecord(keys []string, values []string) (Record, error) {
	if len(keys) < len(values) {
		return nil, csv.ErrFieldCount
	}
	var record = make(Record)
	for i, v := range values {
		record[keys[i]] = strings.TrimSpace(v)
	}
	return record, nil
}
