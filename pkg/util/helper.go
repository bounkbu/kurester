package util

func PriceCountingHelper(price float64) string {
	if price >= 200 {
		return "200++"
	} else if price >= 151 {
		return "151-200"
	} else if price >= 101 {
		return "101-150"
	} else if price >= 51 {
		return "51-100"
	}
	return "0-50"
}
