package finder

import (	
	"os"
	"regexp"
)

func FindByRegExp(filename, expression string, result chan []string) {

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