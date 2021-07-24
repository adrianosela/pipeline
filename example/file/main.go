package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adrianosela/pipeline"
)

var (
	inFile, _  = os.Open("file.txt")
	outFile, _ = os.Create("out.txt")

	scanner = bufio.NewScanner(inFile)
)

func readRepoURL() (interface{}, error) {
	if ok := scanner.Scan(); !ok {
		return nil, pipeline.ErrorSourceFinished
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return scanner.Text(), nil
}

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

func writeRepoName(in interface{}) error {
	repoName, ok := in.(string)
	if !ok {
		return fmt.Errorf("received non string input")
	}

	if n, err := outFile.Write([]byte(repoName + "\n")); n != len(repoName) || err != nil {
		return fmt.Errorf("error writing output: %s", repoName)
	}

	return nil
}

func main() {
	defer inFile.Close()
	defer outFile.Close()

	p := pipeline.New()
	p.SetSource("readRepoURL", readRepoURL)
	p.AddStage("trimURLPrefix", 5, trimURLPrefix)
	p.AddStage("trimRepoOrgPrefix", 5, trimRepoOrgPrefix)
	p.SetSink("writeRepoName", 5, writeRepoName)
	p.Run()
}
