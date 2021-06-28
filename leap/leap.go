// Package leap checks if the year is a leap year in the Gregorian calendar
package leap

// IsLeapYear returns true if a given year is leap
func IsLeapYear(ano int) bool {
	//On every year that is divisible by 4
	if ano%4 == 0 {
		//Except every year that is evenly divisible by 100
		if ano%100 == 0 {
			//Unless the year is also evenly divisible by 400
			if ano%400 == 0 {
				return true
			} else {
				return false
			}
		}
		return true
	}
	return false
}
