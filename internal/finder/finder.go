package finder

import (	
	"os"
	"regexp"
	"path/filepath"
)

func FindStringsInFiles(mask, expression string) []string{
	
	var findedStrings []string
	findedFiles, err := filepath.Glob(mask)

	if err != nil {
		panic(err)
	}

	result := make(chan []string, len(findedFiles))

	for _, filename := range findedFiles {
		go findByRegExp(filename, expression, result)
	}

	for i := 0; i < len(findedFiles); i++ {
		findedStrings = append(findedStrings, <-result...)
	}

	return findedStrings

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