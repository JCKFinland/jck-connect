package seed

import (
	"os"
	"path/filepath"
	"sort"
)

// LoadFiles returns all SQL seed files in execution order.
func LoadFiles() ([]string, error) {

	files, err := filepath.Glob("seeds/*.sql")
	if err != nil {
		return nil, err
	}

	sort.Strings(files)

	return files, nil
}

// ReadFile loads a SQL file into memory.
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
