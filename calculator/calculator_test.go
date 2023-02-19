package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type DiscountRepositoryMock struct {
	DiscountValue int
}

func (drm DiscountRepositoryMock) FindCurrentDiscount() int {
	return drm.DiscountValue
}

// テーブルドリブンテスト
func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		purchaseAmount        int
		discount              int
		expectedAmount        int
	}

	testCases := []testCase{
		{
			name:                  "should apply 20",
			minimumPurchaseAmount: 100,
			purchaseAmount:        150,
			discount:              20,
			expectedAmount:        130,
		},
		{
			name:                  "should apply 40",
			minimumPurchaseAmount: 100,
			purchaseAmount:        200,
			discount:              20,
			expectedAmount:        160,
		},
		{
			name:                  "should apply 60",
			minimumPurchaseAmount: 100,
			purchaseAmount:        350,
			discount:              20,
			expectedAmount:        290,
		},
		{
			name:                  "should not apply",
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			discount:              20,
			expectedAmount:        50,
		},
		{
			name:                  "should not apply when discount is zero",
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			discount:              0,
			expectedAmount:        50,
		},
	}

	for _, tc := range testCases {
		// サブテスト
		t.Run(tc.name, func(t *testing.T) {
			discountRepositoryMock := DiscountRepositoryMock{
				DiscountValue: tc.discount,
			}
			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositoryMock)
			if err != nil {
				t.Fatalf("could not instantiate the calculator: %s", err.Error())
			}
			amount := calculator.Calculate(tc.purchaseAmount)

			assert.Equal(t, tc.expectedAmount, amount)
		})
	}
}

// 失敗するテスト（0除算）
func TestDiscountCalculatorShouldFailWithZeroMinimumAmount(t *testing.T) {
	discountRepositoryMock := DiscountRepositoryMock{}
	_, err := NewDiscountCalculator(0, discountRepositoryMock)
	if err == nil {
		t.Fatalf("should not create the calculator with zero purchase amount")
	}
}
