package stats

import (
	"fmt"
	"github.com/rustamfozilov/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 100,
		},
		{
			Amount: 200,
		},
	}

	fmt.Println(Avg(payments))
	// Output:
	// 150
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount:   100,
			Category: "Аптека",
		},
		{
			Amount:   150,
			Category: "Аптека",
		},
		{
			Amount:   200,
			Category: "Магазин",
		},{
			Amount:   1,
			Category: "Аптека",
		},
	}

	fmt.Println(TotalInCategory(payments, "Аптека"))
	// Output:
	//251
}