package util

func PriceCountingHelper(price float64) string {
	if price <= 50 {
		return "0-50"
	} else if price <= 100 {
		return "51-100"
	} else if price <= 150 {
		return "101-150"
	} else if price <= 200 {
		return "151-200"
	}
	return "200++"
}
