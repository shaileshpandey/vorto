package main

import (
	"fmt"
	"os"
	"strings"

	"main/models"
)

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)
	if len(argsWithoutProg) != 2 || !strings.EqualFold(argsWithoutProg[0], "path_to_problem") {
		fmt.Println(`Usage: Make sure you have Go development environment setup on a MAC machine. You can pick one of two options. 
			Option1 (MAC) - go run main.go path_to_problem actual_directory_with_file
			Option 2 (MAC) - go build . &&  ./main path_to_problem actual_directory_with_file
					e.g. go build . &&  ./main path_to_problem  ~/Downloads/Training\ Problems`)
		os.Exit(1)
	}

	problemPath := argsWithoutProg[1]
	if files, err := os.ReadDir(problemPath); err != nil {
		fmt.Println(err.Error())
	} else {
		loads := models.Loads{}
		loads.ReadLoadsFromFiles(problemPath, files)
	}
}
