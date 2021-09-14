package stats

import (
	"github.com/rustamfozilov/bank/v2/pkg/types"
	"reflect"
	"testing"
)


func TestCategoriesAvg(t *testing.T) {
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
				}, {
					Amount:   50,
					Category: "Аптека",
				},{
					Amount:   40,
					Category: "Магазин",
				},{
					Amount:   100,
					Category: "Авто",
				},
			}

			want := map[types.Category]types.Money{
				"Аптека": 100,
				"Магазин": 120,
				"Авто": 100,
			}

	result :=CategoriesAvg(payments)

if !reflect.DeepEqual(want,result) {
	t.Errorf("invalid result, want: %v, result %v",want,result)
}

}