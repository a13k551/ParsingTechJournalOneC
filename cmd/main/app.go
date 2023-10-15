package main

import (
	"fmt"
	"path/filepath"

	"github.com/a13k551/ParsingTechJournalOneC/internal/config"
	"github.com/a13k551/ParsingTechJournalOneC/internal/finder"
)

func main() {

	var findedStrings []string

	conf := config.Get()

	findedFiles, err := filepath.Glob(conf.Mask)

	if err != nil {
		panic(err)
	}

	result := make(chan []string, len(findedFiles))

	for _, filename := range findedFiles {
		go finder.FindByRegExp(filename, conf.Expression, result)
	}

	for i := 0; i < len(findedFiles); i++ {
		findedStrings = append(findedStrings, <-result...)
	}

	res := fmt.Sprintf("%d matches", len(findedStrings))
	fmt.Println(res)

}