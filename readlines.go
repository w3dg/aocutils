package aocutils

import (
	"bufio"
	"os"
)

// ReadFileLines reads the lines of the specified split on newlines and returns as a `[]string`.
// If the file does not exist or the scanner returns an error, it returns an error.
func ReadFileLines(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ReadFileLines reads the lines of the specified split on two newlines, that is in two blocks separated by two newlines and returns each block as a slice as part of another slice contains all the blocks.
// If the file does not exist or the scanner returns an error, it returns an error.
func ReadFileBlocks(filename string) ([][]string, error) {
	lines, err := ReadFileLines(filename)
	if err != nil {
		return nil, err
	}

	blocks := [][]string{}

	currblock := []string{}
	for _, v := range lines {
		if len(v) == 0 {
			blocks = append(blocks, currblock)
			currblock = []string{}
		} else {
			currblock = append(currblock, v)
		}
	}

	if len(currblock) != 0 {
		blocks = append(blocks, currblock)
	}

	return blocks, nil
}
