package payment

import (
	"fmt"

	"github.com/shopspring/decimal"
)

const centsPerUnit = 100

// YuanToFen converts a decimal amount string (e.g. "10.50") to minor units (int64).
// Uses shopspring/decimal for precision.
func YuanToFen(amountStr string) (int64, error) {
	d, err := decimal.NewFromString(amountStr)
	if err != nil {
		return 0, fmt.Errorf("invalid amount: %s", amountStr)
	}
	return d.Mul(decimal.NewFromInt(centsPerUnit)).IntPart(), nil
}

// FenToYuan converts minor units (int64) to major units as a float64 for interface compatibility.
func FenToYuan(fen int64) float64 {
	return decimal.NewFromInt(fen).Div(decimal.NewFromInt(centsPerUnit)).InexactFloat64()
}
