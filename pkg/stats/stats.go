package stats

import (
	"github.com/rustamfozilov/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	summ := types.Money(0)
	c := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			summ += payment.Amount
			c++
		}
	}
	avg := summ / types.Money(c)
	return avg
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	var sum types.Money
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}
	return sum
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}
	CountOfCategories := countCategoryInPayments(categories, payments)
	for category := range categories {
		categories[category] = categories[category] / CountOfCategories[category]
	}
	return categories
}

func countCategoryInPayments(categories map[types.Category]types.Money, payments []types.Payment) map[types.Category]types.Money {
	CountOfCategories := map[types.Category]types.Money{}
	for category := range categories {
		var counter types.Money
		for _, payment := range payments {
			if payment.Category == category {
				counter++
			}
		}
		CountOfCategories[category] = counter
	}
	return CountOfCategories
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money  {
	Dynamic := map[types.Category]types.Money{}
	for firstCategory := range first {
		 Dynamic[firstCategory] = second[firstCategory] - first[firstCategory]
	}
	for secondCategory := range second {
		      Dynamic[secondCategory]= second[secondCategory] - first[secondCategory]
	}
	return Dynamic
}