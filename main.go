package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"main/models"
)

func main() {
	argsWithoutProg := os.Args[1:]
	var filePath string
	if len(argsWithoutProg) != 1 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Problem File Full Path (No Absolute Path Please)=>")
		filePath, _ = reader.ReadString('\n')
		filePath = strings.Split(filePath, "\n")[0]
	} else {
		filePath = argsWithoutProg[0]
	}

	if file, err := os.Open(filePath); err != nil {
		log.Println(err.Error())
	} else {
		loads := models.Loads{}
		if err := loads.ReadLoadsFromFile(file); err != nil {
			log.Println(err.Error())
		}
		defer file.Close()
	}
}
