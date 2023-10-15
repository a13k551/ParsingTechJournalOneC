package main

import (
	"fmt"

	"github.com/a13k551/ParsingTechJournalOneC/internal/config"
	"github.com/a13k551/ParsingTechJournalOneC/internal/finder"
)

func main() {

	conf := config.Get()

	findedStrings := finder.FindStringsInFiles(conf.Mask, conf.Expression)

	res := fmt.Sprintf("%d matches", len(findedStrings))

	fmt.Println(res)

}
