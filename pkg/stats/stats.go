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



func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money{
	totalSum := map[types.Category]types.Money{}
	count := map[types.Category]types.Money{}
	average := map[types.Category]types.Money{}


	for _, payment := range payments {
		if payment.Amount < 0 {
			continue
		}

		if payment.Status == types.StatusFail {
			continue
		}

		totalSum[payment.Category] += payment.Amount
		count[payment.Category] += 1
		average[payment.Category] = totalSum[payment.Category] / count[payment.Category]
	}

	return average
}


func PeriodsDynamic(first,second map[types.Category]types.Money) map[types.Category]types.Money{
	third := map[types.Category]types.Money{}

	for ind, _ := range second {

		third[ind] = second[ind] - first[ind]

		for indF, _ := range first {
			if indF != ind {
				third[indF] = second[indF] - first[indF]
			}
		}

	}

	return third
}