package geo

import "math"

// Haversine formula to calculate the great-circle distance between two points
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
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

// SphericalLawOfCosines calculates the distance between two points (given the latitude and longitude) on earth using the Spherical Law of Cosines
func SphericalLawOfCosines(lat1, lon1, lat2, lon2 float64) float64 {
	// Radius of Earth in kilometers, can be changed to 3956 for miles
	const R = 6371

	// Convert latitude and longitude from degrees to radians
	lat1Rad := radians(lat1)
	lon1Rad := radians(lon1)
	lat2Rad := radians(lat2)
	lon2Rad := radians(lon2)

	// Calculate the distance using the Spherical Law of Cosines
	distance := math.Acos(math.Sin(lat1Rad)*math.Sin(lat2Rad)+math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Cos(lon2Rad-lon1Rad)) * R

	return distance
}

// Converts degrees to radians
func radians(deg float64) float64 {
	return deg * math.Pi / 180
}
