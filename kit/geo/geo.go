package geo

import "math"

// Haversine formula to calculate the great-circle distance between two points
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// Earth radius in kilometers
	const R = 6371

	// Convert latitude and longitude from degrees to radians
	lat1Rad := radians(lat1)
	lon1Rad := radians(lon1)
	lat2Rad := radians(lat2)
	lon2Rad := radians(lon2)

	// Difference in coordinates
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	// Haversine formula
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	distance := R * c

	return distance
}

// Converts degrees to radians
func radians(deg float64) float64 {
	return deg * math.Pi / 180
}
