package utils

import (
	"log"
	"os"
	"strings"
)

func ImportNewlineSeparatedData(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
