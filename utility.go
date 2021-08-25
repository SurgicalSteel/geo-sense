package main

import "strconv"

func parseCSV(csvLines [][]string, orderLocation Coordinate) []HaversineCalculationData {
	haversineCalculationData := make([]HaversineCalculationData, len(csvLines))

	var (
		tempData                    HaversineCalculationData
		tempLatitude, tempLongitude float64
	)

	if len(csvLines) == 0 {
		return haversineCalculationData
	}

	for i, line := range csvLines {
		tempLatitude, _ = strconv.ParseFloat(line[2], 64)
		tempLongitude, _ = strconv.ParseFloat(line[3], 64)
		tempData = HaversineCalculationData{
			OrderLocation: orderLocation,
			WarehouseLocation: Coordinate{
				Latitude:  tempLatitude,
				Longitude: tempLongitude,
			},
			WarehouseName: line[0],
		}
		haversineCalculationData[i] = tempData
	}

	return haversineCalculationData
}

// Let it out, let it out
// Feel the empty space
// So insecure
// Find the words
// And let it out

// Staring down, staring down
// Nothing comes to mind
// Find the place
// Turn the water into wine

// But I feel I'm getting nowhere
// And I'll never see the end

// So I wither
// And render myself helpless
// I give in
// And everything is clear
// I breakdown
// And let the story guide me

// Turn it on
// Turn it on
// Let the feelings flow
// Close your eyes
// See the ones you used to know

// Open up open up
// Don't struggle to relate
// Lure it out
// Help the memory escape
// Still this barrenness consumes me
// And I feel like giving up

// So I wither
// And render myself helpless
// I give in
// And everything is clear
// I breakdown
// And let the story guide me

// I wither
// And give myself away

// Like reflections on the page
// The world's what you create

// I drown in hesitation
// My words come crashing down
// And all my best creations
// Burning to the ground
// The thought of starting over
// Leaves me paralyzed

// Tear it out again
// Another one that got away

// I wither
// And render myself helpless
// I give in
// And everything is clear

// I wither
// And render myself helpless
// I give in
// And everything is clear
// I breakdown
// And let the story guide me
// I wither
// And give myself away

// Like reflections on the page
// The world's what you create
