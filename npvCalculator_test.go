package resourceAllocator

import "testing"

func TestAdd(t *testing.T) {
  got := add(1, 2)
  if got != 3 {
    t.Errorf("add(1, 2) = %f; want 3", got)
  }
}

func TestCashflowingInvestmentNpv(t *testing.T) {
  investment := cashflowingInvestment{
    cashFlows: []float64{100, 100, 100},
    discountRate: 0.1,
  }
  got := investment.npv()
  if got != 272.72727272727275 {
    t.Errorf("investment.npv() = %f; want 272.72727272727275", got)
  }
}

func TestCashflowingInvestmentNpv2(t *testing.T) {
  investment := cashflowingInvestment{
    cashFlows: []float64{100, 100, 100},
    discountRate: 0.1,
    investmentCost: 100,
  }
  got := investment.npv()
  if got != 172.72727272727275 {
    t.Errorf("investment.npv() = %f; want 172.72727272727275", got)
  }
}
