package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DiscountRepositoryMock struct {
	mock.Mock
}

func (drm DiscountRepositoryMock) FindCurrentDiscount() int {
	args := drm.Called()
	return args.Int(0)
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
		{
			name:                  "should apply 20",
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        150,
			expectedAmount:        130,
		},
		{
			name:                  "should apply 40",
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        200,
			expectedAmount:        160,
		},
		{
			name:                  "should apply 60",
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        350,
			expectedAmount:        290,
		},
		{
			name:                  "should not apply",
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        50,
			expectedAmount:        50,
		},
		{
			name:                  "should not apply when discount is zero",
			minimumPurchaseAmount: 100,
			discount:              0,
			purchaseAmount:        50,
			expectedAmount:        50,
		},
	}

	for _, tc := range testCases {
		// サブテスト
		t.Run(tc.name, func(t *testing.T) {
			discountRepositoryMock := DiscountRepositoryMock{}
			discountRepositoryMock.On("FindCurrentDiscount").Return(tc.discount)
			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositoryMock)
			if err != nil {
				t.Fatalf("could not instantiate the calculator: %s", err.Error())
			}
			amount := calculator.Calculate(tc.purchaseAmount)

			assert.Equal(t, tc.expectedAmount, amount)
		})
	}
}
