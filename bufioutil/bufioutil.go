// Package bufioutil implements some bufio utility functions.
package bufioutil

import "bufio"
import "io"
import "os"

// Reader implements buffering for an io.Reader object.
type Reader struct {
	backend *bufio.Reader
}

// NewReader returns a new Reader.
func NewReader(r io.Reader) (br Reader) {
	br.backend = bufio.NewReader(r)
	return br
}

// ReadLine returns a single line, not including the end-of-line bytes.
func (br Reader) ReadLine() (line string, err error) {
	line, err = br.backend.ReadString('\n')
	if err != nil {
		return "", err
	}
	// skip end-of-line bytes.
	if line[len(line)-1] == '\n' {
		drop := 1
		if line[len(line)-2] == '\r' {
			drop = 2
		}
		line = line[:len(line)-drop]
	}
	return line, nil
}

// ReadLines returns all lines, not including the end-of-line bytes.
func (br Reader) ReadLines() (lines []string, err error) {
	for {
		line, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}

// ReadLines returns all lines, not including the end-of-line bytes.
func ReadLines(filePath string) (lines []string, err error) {
	fr, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fr.Close()
	br := NewReader(fr)
	for {
		line, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}
