package utilities

import (
	"encoding/json"
	"io"
	"os"
)

// AppendJSON appends new data to a JSON file, maintaining a single JSON object.
func AppendJSON(filename string, newData map[string]interface{}) error {
	// Open the file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read existing data
	var existingData map[string]interface{}
	if fileStat, _ := file.Stat(); fileStat.Size() > 0 {
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&existingData); err != nil && err != io.EOF {
			return err
		}
	} else {
		existingData = make(map[string]interface{})
	}

	// Merge new data into existing data
	for k, v := range newData {
		existingData[k] = v
	}

	// Write updated data back to the file
	file.Seek(0, 0)  // Move to the start of the file
	file.Truncate(0) // Clear the file content
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: for pretty-printing
	if err := encoder.Encode(existingData); err != nil {
		return err
	}

	return nil
}
