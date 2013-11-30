package fileinput

import (
	"bufio"
	"io"
	"os"
)

func input(io_ch chan string) {
	var readers []io.Reader
	var file *os.File
	var err error

	// Find all files that we can open
	for i := 1; i < len(os.Args); i++ {
		if file, err = os.Open(os.Args[i]); err == nil {
			readers = append(readers, file)
		}
	}

	// Add stdin if we found no files in the arg list
	if len(readers) == 0 {
		readers = append(readers, os.Stdin)
	}

	// And now read from all the input sources
	for _, reader := range readers {
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			io_ch <- scanner.Text()
		}
	}

	close(io_ch)
}

func Input() chan string {
	io_ch := make(chan string)
	go input(io_ch)
	return io_ch
}
