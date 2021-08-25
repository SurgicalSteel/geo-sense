package main

// Coordinate is the basic struct which represent location of each point
type Coordinate struct {
	//Latitude represents the latitude of a coordinate. Positive for North, Negative for South
	Latitude float64
	//Longitude represents the longitude of a coordinate. Positive for East, Negative for West
	Longitude float64
}

// HaversineCalculationData is the parameter struct for calculating haversine formula to get distance from order location and warehouse location and become the unit of job
type HaversineCalculationData struct {
	OrderLocation     Coordinate
	WarehouseLocation Coordinate
	WarehouseName     string
	Distance          float64
}

// Sometimes
// I feel the fear of
// Uncertainty stinging clear
// And I, can't help but ask myself how much I'll let the fear
// Take the wheel and steer

// It's driven me before and seems to have a vague
// Haunting mass appeal
// But lately I'm beginning to find that I
// Should be the one behind the wheel

// Whatever tomorrow brings I'll be there
// With open arms and open eyes yeah
// Whatever tomorrow brings
// I'll be there, I'll be there

// So, if I
// Decide to waiver my
// Chance to be one of
// The hive
// Will I choose water over wine
// And hold my own and drive?
// Aah ah ooo

// It's driven me before and it seems to be the way
// That everyone else gets around
// But lately I'm beginning to find that when
// I drive myself my light is found

// Whatever tomorrow brings I'll be there
// With open arms and open eyes yeah
// Whatever tomorrow brings
// I'll be there I'll be there

// Would you choose water over wine?
// Hold the wheel and drive

// Whatever tomorrow brings I'll be there
// With open arms and open eyes yeah
// Whatever tomorrow brings
// I'll be there I'll be there
