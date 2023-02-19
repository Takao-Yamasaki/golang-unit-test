package calculator

import "testing"

// テーブルドリブンテスト
func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{name: "should apply 20", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 150, expectedAmount: 130},
		{name: "should apply 40", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 200, expectedAmount: 160},
		{name: "should apply 60", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 350, expectedAmount: 290},
		{name: "zero minimum purchase amount", minimumPurchaseAmount: 0, discount: 20, purchaseAmount: 50, expectedAmount: 50},
	}

	for _, tc := range testCases {
		// サブテスト
		t.Run(tc.name, func(t *testing.T) {
			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
			if err != nil {
				// t.Errorf("could not instantiate the calculator %s", err.Error())
				t.Fatalf("could not instantiate the calculator: %s", err.Error())
			}
			amount := calculator.Calculate(tc.purchaseAmount)
			if amount != tc.expectedAmount {
				t.Errorf("expected %v, got %v", tc.expectedAmount, amount)
			}
		})
	}
}
