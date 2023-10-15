package finder

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func FindStringsInFiles(mask, expression string) []string {

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

	myRegExp, err := regexp.Compile("-[0-9]{7,15},CALL.*")

	if err != nil {
		panic(err)
	}

	matchedStrings := myRegExp.FindAll(file, 100000)

	for _, matchedStringByte := range matchedStrings {

		replaceCallrRegexp, _ := regexp.Compile("CALL.*:")
		replaceCall := replaceCallrRegexp.ReplaceAll(matchedStringByte, []byte(""))

		replaceMemoryRegexp, _ := regexp.Compile(",Memory.*")
		replaceMemory := replaceMemoryRegexp.ReplaceAll(replaceCall, []byte(""))

		matchedString := string(replaceMemory[:])
		findedStrings = append(findedStrings, matchedString)

		fmt.Println(matchedString)
	}

	result <- findedStrings

}
