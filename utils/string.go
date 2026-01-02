package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	Value     string
	Separator string
	Children  []Node
}

func (f Node) String() string {
	json, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshalling node: %v", err)
	}
	return string(json)
}

func convertToNodes(data string, separators []string) Node {
	value := strings.TrimSpace(data)
	separator := ""
	if len(separators) > 0 {
		separator = separators[0]
	}
	children := make([]Node, 0)

	if len(separators) > 0 {
		childValues := strings.Split(value, separator)
		children = make([]Node, len(childValues))
		for i, childValue := range childValues {
			children[i] = convertToNodes(childValue, separators[1:])
		}
	}

	return Node{Value: value, Separator: separator, Children: children}
}

func ImportMultiDimensionalData(filename string, separators []string) Node {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return convertToNodes(string(data), separators)
}
