package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adrianosela/pipeline"
)

func getScanner(filename string) (*os.File, *bufio.Scanner) {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	return file, bufio.NewScanner(file)
}

func main() {
	fd, scanner := getScanner("file.txt")
	defer fd.Close()

	p := pipeline.New()

	p.SetSource("readFileLines", func() (interface{}, error) {
		if ok := scanner.Scan(); !ok {
			return nil, pipeline.ErrorSourceFinished
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		return scanner.Text(), nil
	})

	p.AddStage("parseRepoName", func(in interface{}) (interface{}, error) {
		url, ok := in.(string)
		if !ok {
			return nil, fmt.Errorf("received non string input")
		}
		return strings.TrimPrefix(url, "https://github.com/"), nil
	})

	p.SetSink("printRepoName", func(in interface{}) error {
		repoName, ok := in.(string)
		if !ok {
			return fmt.Errorf("received non string input")
		}

		fmt.Println(repoName)
		return nil
	})

	p.Run()
}
