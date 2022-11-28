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
	} else if price <= 250 {
		return "201-250"
	} else if price <= 300 {
		return "251-300"
	} else if price <= 350 {
		return "301-350"
	} else if price <= 400 {
		return "351-400"
	} else if price <= 450 {
		return "401-450"
	}
	return "500++"
}
