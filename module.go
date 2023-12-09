package readinput

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(fileName string) ([]string, error) {
	// Opening file input.txt
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("there was a problem reading the file: %v", err)
	}

	// Creating a scanner for the opened file
	scanner := bufio.NewScanner(file)

	// Tells scanner how it wants to split (by lines)
	// bufio.ScanLines is default, but we can be explicit here
	scanner.Split(bufio.ScanLines)

	// Scanning through file and construct output
	output := []string{}
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	return output, nil
}
