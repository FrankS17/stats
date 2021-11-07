package stats

import "github.com/FrankS17/bank/v2/pkg/types"

// Avg находмт среднюю сумму одного платежа
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	count := 0

	for _, payment := range payments {
		
		if payment.Amount < 0 {
			continue
		}
		
		if payment.Status == types.StatusFail {
			continue
		}

		sum += types.Money(payment.Amount)
		count++
	}

	avg := sum / types.Money(count)
	
	return avg
}


// TotalInCategory находит сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment,category types.Category) types.Money {
	
	sum := types.Money(0)

	for _, payment := range payments {

		if payment.Amount < 0 {
			continue
		}

		if payment.Category != category {
			continue
		} 

		if payment.Status == types.StatusFail {
			continue
		}

		sum += types.Money(payment.Amount)
	}
	
	return sum
}