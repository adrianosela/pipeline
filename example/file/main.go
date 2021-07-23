package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adrianosela/pipeline"
)

func readFileLines(c chan interface{}) {
	defer close(c)

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	p := pipeline.New()
	p.SetSource("readFileLines", func() (interface{}, error) {
		if dataWasRead := scanner.Scan(); !dataWasRead {
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
