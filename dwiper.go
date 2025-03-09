package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dsa-ferreira/dwiper/internal/files"
	"github.com/dsa-ferreira/dwiper/internal/terminal"
	"github.com/dsa-ferreira/dwiper/internal/validator"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("I need a folder to clear!")
		os.Exit(1)
	}
	dir := os.Args[1]
	possibleCadidates := make([]string, 0)

	for _, file := range files.GetDirFiles(dir) {
		info, err := file.Info()
		if err != nil {
			log.Fatal(err)
		}

		if validator.ImmediateRemoval(file.Name()) {
			files.Remove(dir, file.Name())
		}

		if validator.CheckCandidacy(info) {
			possibleCadidates = append(possibleCadidates, info.Name())
		}
	}

	for _, file := range terminal.InteractiveSelection(possibleCadidates) {
		files.Remove(dir, file)
	}

}
