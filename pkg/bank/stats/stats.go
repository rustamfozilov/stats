package stats

import "github.com/rustamfozilov/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var  summ types.Money
	for _, payment := range payments {
		summ += payment.Amount
	}
	average := summ / types.Money(len(payments))
	return average
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	var sum types.Money
	for _, payment:= range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}