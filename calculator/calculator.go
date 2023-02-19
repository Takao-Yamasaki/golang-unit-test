package calculator

import (
	"errors"

	"github.com/Takao-Yamasaki/golang-unit-test/database"
)

// 割引計算機
type DiscountCalculator struct {
	minimumPurchaseAmount int                 // 最低購入金額
	discountRepository    database.Repository // 割引額
}

// 割引を適用するための最低購入額と割引額を受け取る
func NewDiscountCalculator(minimumPurchaseAmount int, discountRepository database.Repository) (*DiscountCalculator, error) {

	if minimumPurchaseAmount == 0 {
		return &DiscountCalculator{}, errors.New("minimum purchase amount could not be zero")
	}

	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountRepository:    discountRepository,
	}, nil
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {
	// 購入金額が最低購入金額よりも高い場合は割引が適用される
	if purchaseAmount > c.minimumPurchaseAmount {

		// 乗数(自動的に丸められる)
		multiplier := purchaseAmount / c.minimumPurchaseAmount

		discount := c.discountRepository.FindCurrentDiscount()

		return purchaseAmount - (discount * multiplier)
	}
	// それ以外は購入金額を返す
	return purchaseAmount
}
