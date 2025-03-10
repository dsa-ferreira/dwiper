package validator

import (
	"os"
	"regexp"
	"time"

	"github.com/dsa-ferreira/dwiper/internal/config"
	"github.com/dsa-ferreira/dwiper/internal/terminal"
)

var selectedRegexes []string

var configuration config.Configuration = config.ParseConfig()

func ImmediateRemoval(fileName string) bool {
	if selectedRegexes == nil {
		selectedRegexes = terminal.InteractiveSelection(configuration.Regexes)
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
	if duration.Hours() > configuration.KeepDuration {
		return true
	}
	return false

}
