package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"BPL/parser"
)

// Finding all files with the extension "bpl"
func WalkMatch(root, pattern string) ([]string, error) {
    var matches []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
            return err
        } else if matched {
            matches = append(matches, path)
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    return matches, nil
}

// Reading all WalkMatch files, extracting text and putting into string to parse
func main() {
	stringSlice := []string{}
	files, err := WalkMatch("examples/", "*.bpl")
	for _, element := range files {
		byte, err := ioutil.ReadFile(element)
		if err != nil {
			log.Println("main.go - main() Error loading files:", err.Error())
			return
		}
		stringSlice = append(stringSlice, string(byte))
	}

	parsedBPLFile := parser.Parse(strings.Join(files,", "), strings.Join(stringSlice, " "))
	prettyJSON, err := json.MarshalIndent(parsedBPLFile, "", "   ")

	if err != nil {
		log.Println("main.go - main() Error marshalling JSON:", err.Error())
		return
	}

	log.Println(string(prettyJSON))
}