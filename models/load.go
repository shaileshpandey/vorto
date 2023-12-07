package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Load struct {
	ID      int
	Pickup  Coordinate
	DropOff Coordinate
}

func formatLoad(line string) (load Load, err error) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return Load{}, fmt.Errorf("invalid load format: %s", line)
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return Load{}, fmt.Errorf("failed to parse load ID: %v", err)
	}
	load.ID = id
	load.formatPickupPoint(parts[1])
	load.formatDropOffPoint(parts[2])
	return load, nil
}

func (load *Load) formatPickupPoint(coordinates string) {
	load.Pickup.formatPoint(coordinates)
}

func (load *Load) formatDropOffPoint(coordinates string) {
	load.DropOff.formatPoint(coordinates)
}
