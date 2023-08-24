package zaputils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Progress is a struct which contains the progress of the scan.
type Progress struct {
	Fail int32 `json:"fail"`
	Warn int32 `json:"warn"`
	Pass int32 `json:"pass"`
}

func GetProgress(file string) (Progress, error) {
	var progress Progress

	f, err := os.Open(file)
	if err != nil {
		return progress, fmt.Errorf("failed to open json file: %w", err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return progress, fmt.Errorf("failed to read json file: %w", err)
	}

	if err := json.Unmarshal(data, &progress); err != nil {
		return progress, fmt.Errorf("failed to unmarshal progress: %w", err)
	}

	return progress, nil
}
