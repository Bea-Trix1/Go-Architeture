package singleresponsabilitymodel

import (
	"os"
)

func PrintNews(filename string, journal Journal) {
	_ = os.WriteFile(filename, []byte(journal.String()), 0644)
}
