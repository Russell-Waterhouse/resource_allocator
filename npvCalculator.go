package resourceAllocator

func add(a, b float64) float64 {
  return a + b
}

type hasNpv interface {
  npv() float64
}

type cashflowingInvestment struct {
  cashFlows []float64
  discountRate float64
  investmentCost float64
}

func (n *cashflowingInvestment) npv() float64 {
  var total float64
  for i, cashFlow := range n.cashFlows {
    total += cashFlow / (1 + n.discountRate) * float64(i)
  }
  return total - n.investmentCost
}

