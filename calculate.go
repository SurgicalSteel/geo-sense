package main

import "math"

//CalculateDistance calculates distance of two given points using Haversine formula
func CalculateDistance(pointA, pointB Coordinate) float64 {
	var distance float64 = 0.0

	var radPerDeg float64 = math.Pi / 180.0
	const earthRadiusKilometer float64 = 6371.0

	var deltaLatitudeRad float64 = (pointB.Latitude - pointA.Latitude) * radPerDeg
	var deltaLongitudeRad float64 = (pointB.Longitude - pointA.Longitude) * radPerDeg

	var latPointARad float64 = pointA.Latitude * radPerDeg
	var latPointBRad float64 = pointB.Latitude * radPerDeg

	var sinDeltaLatitudeRad float64 = math.Sin(deltaLatitudeRad / 2)
	var sinDeltaLongitudeRad float64 = math.Sin(deltaLongitudeRad / 2)

	var a float64 = squareFloat64(sinDeltaLatitudeRad) + (math.Cos(latPointARad) * math.Cos(latPointBRad) * squareFloat64(sinDeltaLongitudeRad))
	var c float64 = 2.0 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance = earthRadiusKilometer * c

	return distance
}

func squareFloat64(a float64) float64 {
	return a * a
}

// Mother, tell me
// I long to hear the stories
// Just like long ago
// All these memories
// Start to fade before me
// I cannot let them go
// No, I can't let them go

// From the cradle to the grave
// It's a fear I can't escape
// Who will be my hiding place
// When you're gone?

// Father, stay here
// Don't leave me like the other
// Know, I need you so
// All you gave me
// Father, how you loved me
// Treated like your own
// Don't ever let me go

// From the cradle to the grave
// It's a fear I can't escape
// Who will be my hiding place
// When you're gone?

// There's no way to deny
// The brevity of life

// As time keeps marching on
// All we have is lost

// As time keeps marching on
// All we have is lost

// Nothing lasts forever
// Nothing stays the same

// From the cradle to the grave
// It's a fear I can't escape
// Who will be my hiding place
// When you're gone?

// From the cradle to the grave
// It's a fear I can't escape
// Who will be my hiding place
// When you're gone?

// How am I to carry on?
// How am I to carry on? (as time keeps marching on)
// How am I to carry on? (all we have is lost)
// Nothing lasts forever
// Nothing ever stays the same
