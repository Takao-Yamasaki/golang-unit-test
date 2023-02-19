package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type DiscountRepositoryMock struct{}

func (drm DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

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
		{name: "should apply 20", minimumPurchaseAmount: 100, purchaseAmount: 150, expectedAmount: 130},
		{name: "should apply 40", minimumPurchaseAmount: 100, purchaseAmount: 200, expectedAmount: 160},
		{name: "should apply 60", minimumPurchaseAmount: 100, purchaseAmount: 350, expectedAmount: 290},
	}

	for _, tc := range testCases {
		// サブテスト
		t.Run(tc.name, func(t *testing.T) {
			discountRepositoryMock := DiscountRepositoryMock{}
			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositoryMock)
			if err != nil {
				t.Fatalf("could not instantiate the calculator: %s", err.Error())
			}
			amount := calculator.Calculate(tc.purchaseAmount)

			assert.Equal(t, 2, amount)
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
