package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type OpenSALT struct {
	CFDocument CFDocument `json:"CFDocument"`
	CFItems    []CFItems  `json:"CFItems"`
}
type CFDocument struct {
	OfficialSourceURL string `json:"officialSourceURL"`
	Title             string `json:"title"`
}
type CFItems struct {
	FullStatement     string `json:"fullStatement"`
	HumanCodingScheme string `json:"humanCodingScheme,omitempty"`
}

func printStandards(data OpenSALT) {
	items := make(map[string]string)
	var keys []string

	for _, item := range data.CFItems {
		if item.HumanCodingScheme != "" {
			if _, ok := items[item.HumanCodingScheme]; !ok {
				items[item.HumanCodingScheme] = item.FullStatement
				keys = append(keys, item.HumanCodingScheme)
			}
		}
	}

	sort.Strings(keys)

	fmt.Printf("%s - %s\n", data.CFDocument.Title, data.CFDocument.OfficialSourceURL)
	fmt.Println(strings.Repeat("-", 50))
	for _, key := range keys {
		fmt.Printf("%-20s%.150s\n", key, items[key])
	}
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("%d Standards found\n", len(keys))
}

func parseFile(path string) (OpenSALT, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return OpenSALT{}, err
	}

	var data OpenSALT
	err = json.Unmarshal(content, &data)
	if err != nil {
		return OpenSALT{}, err
	}

	return data, nil
}

func main() {
	var path string
	flag.StringVar(&path, "p", "opensalt-framework.json", "Path to opensalt JSON file")

	flag.Parse()

	data, err := parseFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printStandards(data)
}
