// ex7.1 create WordCounter and LineCounter. Using bufio package
package counters

import (
	"bufio"
	"bytes"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}

	*w += WordCounter(count)
	return count, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	count := 0
	for scanner.Scan() {
		count++
	}
	*l += LineCounter(count)
	return count, nil
}
