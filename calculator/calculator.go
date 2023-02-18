package calculator

// 割引計算機
type DiscountCalculator struct {
	minimumPurchaseAmount int // 最低購入金額
	discountAmount        int // 割引額
}

// 割引を適用するための最低購入額と割引額を受け取る
func NewDiscountCalculator(minimumPurchaseAmount int, discountAmount int) *DiscountCalculator {
	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountAmount:        discountAmount,
	}
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {
	// 購入金額が最低購入金額よりも高い場合は割引が適用される
	if purchaseAmount > c.minimumPurchaseAmount {

		// 乗数(自動的に丸められる)
		multiplier := purchaseAmount / c.minimumPurchaseAmount

		return purchaseAmount - (c.discountAmount * multiplier)
	}
	// それ以外は購入金額を返す
	return purchaseAmount
}
