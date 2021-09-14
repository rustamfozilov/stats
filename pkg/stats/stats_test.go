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
		}, {
			Amount:   40,
			Category: "Магазин",
		}, {
			Amount:   100,
			Category: "Авто",
		},
	}

	want := map[types.Category]types.Money{
		"Аптека":  100,
		"Магазин": 120,
		"Авто":    100,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(want, result) {
		t.Errorf("invalid result, want: %v, result %v", want, result)
	}

}

func TestPeriodsDynamic_1(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
	}
	expected := map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
	}
	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("invalid result, expected: %v, result: %v", expected, result)
	}

}

func TestPeriodsDynamic_2(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{

		"food": 20,
	}
	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
	}
	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("invalid result, expected: %v, result: %v", expected, result)
	}
}

func TestPeriodsDynamic_3(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 10,
		"food": 25,
		"mobile": 5,
	}
	expected := map[types.Category]types.Money{
		"auto": 0,
		"food": 5,
		"mobile": 5,
	}
	result := PeriodsDynamic(first,second)
	if !reflect.DeepEqual(result,expected){
		t.Errorf("invalid result, expected: %v, result: %v", expected,result)
	}

}