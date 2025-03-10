package files

import (
	"fmt"
	"log"
	"os"
)

func GetDirFiles(dir string) []os.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return make([]os.DirEntry, 0)
	}
	return files
}

func Remove(dir string, file string) {
	os.Remove(dir + "/" + file)
}

func ReadFile(file string) string {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	return string(fileBytes)

}
