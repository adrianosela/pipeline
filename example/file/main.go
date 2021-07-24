package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adrianosela/pipeline"
)

func trimURLPrefix(in interface{}) (interface{}, error) {
	url, ok := in.(string)
	if !ok {
		return nil, fmt.Errorf("received non string input")
	}
	return strings.TrimPrefix(url, "https://github.com/"), nil
}

func trimRepoOrgPrefix(in interface{}) (interface{}, error) {
	url, ok := in.(string)
	if !ok {
		return nil, fmt.Errorf("received non string input")
	}

	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("malformed url")
	}

	return strings.Join(parts[1:], "/"), nil
}

func printRepoName(in interface{}) error {
	repoName, ok := in.(string)
	if !ok {
		return fmt.Errorf("received non string input")
	}

	fmt.Println(repoName)
	return nil
}

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	readFileLine := func() (interface{}, error) {
		if ok := scanner.Scan(); !ok {
			return nil, pipeline.ErrorSourceFinished
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		return scanner.Text(), nil
	}

	p := pipeline.New()
	p.SetSource("readFileLines", readFileLine)
	p.AddStage("trimURLPrefix", 5, trimURLPrefix)
	p.AddStage("trimRepoOrgPrefix", 5, trimRepoOrgPrefix)
	p.SetSink("printRepoName", 5, printRepoName)
	p.Run()
}
