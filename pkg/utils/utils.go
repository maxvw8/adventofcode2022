package utils

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// StringParseFn defines function which receives text and try to parse to a type.
// It returns the converted type or an error if the convertion was not possible
type StringParseFn[T any] func(string) (T, error)

func ReadFile[T any](r io.Reader, parser func(string) (T, error)) ([]T, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []T
	var line string
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		x, err := parser(line)
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func ReadBlocks[T any](r io.Reader, parser func(string) (T, error)) ([][]T, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result [][]T
	var block []T
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			//block finished
			result = append(result, block)
			block = nil

		} else {
			x, e := parser(line)
			if e != nil {
				return result, e
			}
			block = append(block, x)
		}
	}
	if len(block) != 0 {
		result = append(result, block)
	}
	return result, scanner.Err()
}

// todo :)
func SplitBlocks(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return 0, nil, nil
	}
	first := bytes.IndexByte(data, '\n')
	if first < 0 || len(data) < first+1 {
		return len(data), data, nil
	}
	if i := bytes.IndexByte(data[first+1:], '\n'); i < 1 {
		token = data[0 : first+1] //skipping second \n"
		advance = first + 2
		return advance, token, nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
