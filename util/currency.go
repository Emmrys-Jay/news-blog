package util

const (
	USD = "USD"
	CAD = "CAD"
	CHF = "CHF"
	EUR = "EUR"
	JPY = "JPY"
	NGN = "NGN"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, CAD, CHF, EUR, JPY, NGN:
		return true
	}
	return false
}
