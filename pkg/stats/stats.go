package stats

import "github.com/rustamfozilov/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	var  summ types.Money
	for _, payment := range payments {
		if payment.Status != types.StatusFail{
			summ += payment.Amount
		}

	}
	average := summ / types.Money(len(payments))
	return average
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	var sum types.Money
	for _, payment:= range payments {
		if payment.Category == category && payment.Status != types.StatusFail{
			sum += payment.Amount
		}
	}
	return sum
}