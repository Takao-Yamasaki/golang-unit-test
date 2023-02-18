package copycalculator

import "testing"

// 割引が適用される場合のテスト
func TestDiscountApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(150)

	if 130 != amount {
		t.Fail()
	}
}

// 割引が適用されない場合のテスト
func TestDiscountNotApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(50)

	if amount != 60 {
		// Error = Log + Fail
		// t.Logf("expected 50, got %v", amount)
		// t.Fail()
		t.Errorf("expected 50, got %v, failed because the discount was not expected to be applied", amount)
	}
}
