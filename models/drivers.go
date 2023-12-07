package models

import "sort"

type Drivers []Driver

var zeroPoint Coordinate = Coordinate{0, 0}

func (drivers *Drivers) assignLoadToDrivers(loads []Load) {
	var currentDriver Driver

	sort.Slice(loads, func(i, j int) bool {
		return zeroPoint.euclidDistance(loads[i].Pickup) < zeroPoint.euclidDistance(loads[j].Pickup)
	})

	for _, load := range loads {
		lastLoadIndex := len(currentDriver.Loads) - 1
		if len(currentDriver.Loads) == 0 ||
			zeroPoint.euclidDistance(currentDriver.Loads[lastLoadIndex].DropOff)+
				currentDriver.Loads[lastLoadIndex].DropOff.euclidDistance(load.Pickup) < 720 {
			currentDriver.Loads = append(currentDriver.Loads, load)
		} else {
			*drivers = append(*drivers, currentDriver)
			currentDriver = Driver{}
			currentDriver.Loads = append(currentDriver.Loads, load)
		}
	}

	if len(currentDriver.Loads) > 0 {
		*drivers = append(*drivers, currentDriver)
	}
}

func (drivers Drivers) displayLoads() {
	for _, driver := range drivers {
		driver.Loads.displayLoads()
	}
}
