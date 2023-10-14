package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/a13k551/ParsingTechJournalOneC/internal/config"
)

func main() {

	var findedStrings []string

	conf := config.Get()

	findedFiles, _ := filepath.Glob(conf.Mask)

	result := make(chan []string, len(findedFiles))

	for _, filename := range findedFiles {
		go findByRegExp(filename, conf.Expression, result)
	}

	for i := 0; i < len(findedFiles); i++ {
		findedStrings = append(findedStrings, <-result...)

	}

	fmt.Println(len(findedStrings))

}

func findByRegExp(filename, expression string, result chan []string) {

	var findedStrings []string

	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	myRedExp, err := regexp.Compile(expression)

	if err != nil {
		panic(err)
	}

	matchedStrings := myRedExp.FindAll(file, 10000)

	for _, matchedStringByte := range matchedStrings {

		matchedString := string(matchedStringByte[:])

		findedStrings = append(findedStrings, matchedString)
	}

	result <- findedStrings

}
