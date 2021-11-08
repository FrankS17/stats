package stats

import (
	"reflect"
	"testing"

	"github.com/FrankS17/bank/v2/pkg/types"
)

func TestCategoriesAVG_nil(t * testing.T) {
	var payments []types.Payment
	result := CategoriesAVG(payments)

	if len(result) != 0 {
	 t.Error("Result len != 0")
	}
}

func TestCategoriesAVG_empty(t * testing.T) {
	payments := []types.Payment{}
	result := CategoriesAVG(payments)

	if len(result) != 0 {
	 t.Error("Result len != 0")
	}
}

func TestCategoriesAVG(t * testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 20_00},
		{ID: 2, Category: "auto", Amount: 10_00},
		{ID: 3, Category: "food", Amount: 20_00},
		{ID: 4, Category: "food", Amount: 10_00},
		{ID: 5, Category: "auto", Amount: 10_00},
	}
	expected := map[types.Category]types.Money{
		"auto":13_33,
		"food":15_00,
	}

	result := CategoriesAVG(payments)

	if !reflect.DeepEqual(expected,result) {
	 t.Errorf("invalid result, expected: %v, actual: %v",expected,result)
	}
}

func TestPeriodsDynamic(t * testing.T) {
	
	first := map[types.Category]types.Money{
		"auto":10_00,
		"food":10_00,
		"mobi":10_00,
	}

	second := map[types.Category]types.Money{
		"auto":30_00,
		"food":30_00,
		"phone":30_00,
	}


	expected := map[types.Category]types.Money{
		"auto":20_00,
		"food":20_00,
		"mobi":-10_00,
		"phone":30_00,
	}

	result := PeriodsDynamic(first,second)

	if !reflect.DeepEqual(expected,result) {
	 t.Errorf("invalid result, expected: %v, actual: %v",expected,result)
	}
}
