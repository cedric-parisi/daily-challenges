package julius

import (
	"math"
)

func leaps(year1, year2 uint64) uint64 {
	count := uint64(0)

	if year1 >= year2 {
		return count
	}

	// There are 218 leap years every 900 years:
	// 900/4 = 225 are leap years
	// 900/100 = 9 are exception
	// the year 200 and 600 are exception to the exception, therefore 2 leap years
	// 225 - 9 + 2 = 218
	if diff := year2 - year1; diff >= 900 {
		count = uint64(218 * math.Floor(float64(diff/900)))
		// We move year1 to the start of the remaining years
		year1 = year2 - diff%900
	}

	// for every years in range
	// the difference between year1 and year2 is never > 900
	for i := year1; i < year2; i++ {
		if isLeap(i) {
			count++
		}
	}
	return count
}

func isLeap(year uint64) bool {
	return (year%4 == 0 && year%100 != 0) || (year%900 == 200 || year%900 == 600)
}
