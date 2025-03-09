package validator

import (
	"os"
	"regexp"
	"time"

	"github.com/dsa-ferreira/dwiper/internal/terminal"
)

var regexes = []string{"^Pokemon.*\\.zip$", "bla", "choo"}
var selectedRegexes []string

func ImmediateRemoval(fileName string) bool {
	if selectedRegexes == nil {
		selectedRegexes = terminal.InteractiveSelection(regexes)
	}

	for _, r := range selectedRegexes {
		regex, _ := regexp.Compile(r)
		if regex.MatchString(fileName) {
			return true
		}
	}
	return false
}

func CheckCandidacy(file os.FileInfo) bool {
	duration := time.Now().Sub(file.ModTime())
	if duration.Hours() > 24 {
		return true
	}
	return false

}
