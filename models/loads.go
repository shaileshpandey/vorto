package models

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type Loads []Load

func (loads *Loads) ReadLoadsFromFiles(problemPath string, files []fs.DirEntry) {
	for _, file := range files {
		err := loads.readLoadsFromFile(fmt.Sprintf("%s/%s", problemPath, file.Name()))
		if err != nil {
			fmt.Println("Error reading the file:", err)
			os.Exit(1)
		}
		drivers := Drivers{}
		drivers.assignLoadToDrivers(*loads)
		drivers.displayLoads()
	}
}

func (loads *Loads) readLoadsFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.EqualFold("loadNumber pickup dropoff", line) {
			load, err := formatLoad(line)
			if err != nil {
				return err
			}
			*loads = append(*loads, load)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (loads Loads) displayLoads() {
	var loadIDs []int
	for _, load := range loads {
		loadIDs = append(loadIDs, load.ID)
	}
	fmt.Println(loadIDs)
}
