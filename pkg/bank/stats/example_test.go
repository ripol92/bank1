package stats

import (
	"github.com/ripol92/bank/pkg/bank/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{ID: 1, Amount: types.Money(10)},
		{ID: 2, Amount: types.Money(20)},
		{ID: 3, Amount: types.Money(30)},
	}

	avg := Avg(payments)
	fmt.Println(avg)

	// Output: 20
}
