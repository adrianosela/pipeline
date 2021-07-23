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
	p := pipeline.New(readFileLines)

	p.AddStage("parseRepoName", func(in interface{}) (interface{}, error) {
		url, _ := in.(string)
		return strings.TrimPrefix(url, "https://github.com/"), nil
	})

	p.AddStage("printRepoName", func(in interface{}) (interface{}, error) {
		repoName, _ := in.(string)
		fmt.Println(repoName)
		return nil, nil
	})

	p.Run()
}
