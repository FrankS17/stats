package stats

import (
	"fmt"
	"github.com/FrankS17/bank/v2/pkg/types"
)

func Example_Avg() {
	
	payment := []types.Payment{
		{
			ID:     1,
			Amount: 20_00,
			Status: types.StatusFail,
		},
		{
			ID:     2,
			Amount: 20_00,
			Status: types.StatusOk,
		},
		{
			ID:     3,
			Amount: -20_00,
			Status: types.StatusOk,
		},
		{
			ID:     3,
			Amount: 40_00,
			Status: types.StatusOk,
		},
	}

	fmt.Println(Avg(payment))
	// Output: 
	// 3000		
	}

	
func Example_TotalInCategory() {

	payment := []types.Payment{
		{
			ID:     1,
			Amount: 20_00,
			Category: "A",
		},
		{
			ID:     2,
			Amount: 20_00,
			Category: "B",
		},
		{
			ID:     3,
			Amount: 20_00,
			Category: "A",
		},
		{
			ID:     4,
			Amount: -20_00,
			Category: "A",
		},
		{
			ID:     4,
			Amount: 25_00,
			Category: "F",
		},
		{
			ID:     4,
			Amount: 30_00,
			Category: "F",
		},
		{
			ID:     4,
			Amount: 20_00,
			Category: "P",
		},
	}

	fmt.Println(TotalInCategory(payment,"A"))
	// Output:
	// 4000
}




 
