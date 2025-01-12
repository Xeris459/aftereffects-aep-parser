package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	aep "github.com/xeris459/aftereffects-aep-parser/src"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <path-to-aep-file>", os.Args[0])
	}

	// Path file AEP dari argumen
	filePath := os.Args[1]

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s", filePath)
	}

	project, err := aep.Open(filePath)
	// project, err := Open("F:/Project Video/hi3contest.aep")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	projectJSON, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(string(projectJSON))
}
