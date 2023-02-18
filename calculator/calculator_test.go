package calculator

import "testing"

// テーブルドリブンテスト
func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 150, expectedAmount: 130},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 200, expectedAmount: 160},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 350, expectedAmount: 290},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 50, expectedAmount: 50},
	}

	for _, tc := range testCases {
		calculator := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
		amount := calculator.Calculate(tc.purchaseAmount)
		if amount != tc.expectedAmount {
			t.Errorf("expected %v, got %v", tc.expectedAmount, amount)
		}
	}
}

// 割引が適用される場合のテスト
func TestDiscountApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(150)
	// 150 / 100 = 1
	// 20 * 1 = 20
	if amount != 130 {
		t.Errorf("expected 130, got %v", amount)
	}
}

// 乗数のテスト
func TestDiscountMultipliedByTwoApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(200)
	// 200 / 100 = 2
	// 20 * 2 = 40
	if amount != 160 {
		t.Errorf("expected 160, got %v", amount)
	}
}

func TestDiscountMultipliedByThreeApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(350)
	// 350 / 100 = 3
	// 20 * 3 = 60
	if amount != 290 {
		t.Errorf("expected 290, got %v", amount)
	}
}

// 割引が適用されない場合のテスト
func TestDiscountNotApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(50)

	if amount != 50 {
		t.Errorf("expected 50, got %v", amount)
	}
}
