package stats

import "github.com/ripol92/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	count := len(payments)

	sum := types.Money(0)

	for _, payment := range payments {
		sum += payment.Amount
	}

	return sum / count
}