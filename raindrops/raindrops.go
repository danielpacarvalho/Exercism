// Package raindrops makes sounds in the form of string based on what factor the number has
package raindrops

import (
	"math"
	"strconv"
)

// Convert creates a string based on the factor of the received number
func Convert(pingos int) string {
	volta := ""
	factor3 := math.Remainder(float64(pingos), 3)
	factor5 := math.Remainder(float64(pingos), 5)
	factor7 := math.Remainder(float64(pingos), 7)

	if factor3 == 0 {
		volta += "Pling"
	}
	if factor5 == 0 {
		volta += "Plang"
	}
	if factor7 == 0 {
		volta += "Plong"
	}

	//If there are no factors, just return the received number
	if factor3 != 0 && factor5 != 0 && factor7 != 0 {
		volta = strconv.Itoa(pingos)
	}

	return volta
}
