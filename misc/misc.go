package misc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Typefile struct {
	Position struct {
		X float32 `json:"X"`
		Y float32 `json:"Y"`
	} `json:"position"`
	Size struct {
		W float32 `json:"W"`
		H float32 `json:"H"`
	} `json:"size"`
	Name string `json:"name"`
}

type Typefiles []Typefile

// Function to read JSON file into the slice of filetime structs
func Readjsonfile(filepath string, vara any) error {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open JSON file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %w", err)
	}
	if err := json.Unmarshal(byteValue, vara); err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	return nil
}
