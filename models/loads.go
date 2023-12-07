package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Loads []Load

func (loads *Loads) ReadLoadsFromFile(file *os.File) error {
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
	drivers := Drivers{}
	drivers.assignLoadToDrivers(*loads)
	drivers.displayLoads()
	return nil
}

func (loads Loads) displayLoads() {
	var loadIDs []int
	for _, load := range loads {
		loadIDs = append(loadIDs, load.ID)
	}
	fmt.Println(loadIDs)
}
