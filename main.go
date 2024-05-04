package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func mapper(data []byte, n int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	chunk := string(data[:n])
	c <- len(strings.Fields(chunk))
}

func reducer(count *int, c chan int) {
	for num := range c {
		*count += num
	}
}

func countWords(filename string) (int, error) {
	var count int = 0
	fileIn, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer fileIn.Close()

	reader := bufio.NewReader(fileIn)
	chunkSize := 512

	processChan := make(chan int)
	var wg sync.WaitGroup

	for {
		buf := make([]byte, chunkSize)
		numBytes, err := reader.Read(buf)

		if err != nil {
			fmt.Printf("error reading file: %s\n", err)
			return 0, err
		}
		wg.Add(1)
		go mapper(buf, numBytes, processChan, &wg)

		if _, err := reader.Peek(1); err == io.EOF {
			break
		}
	}
	go func() {
		wg.Wait()
		close(processChan)
	}()

	reducer(&count, processChan)

	return count, nil
}

func main() {
	var filename string

	flag.StringVar(&filename, "filename", "", "Filename to process")

	flag.Parse()

	if filename == "" {
		fmt.Println("invalid file name")
		os.Exit(1)
	}

	wc, err := countWords(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Words counted: %d\n", wc)
}
