package payment

import (
	"github.com/Takao-Yamasaki/golang-unit-test/database"
	"github.com/Takao-Yamasaki/golang-unit-test/entity"
)

type PaymentService struct {
	attemptHistoryRepository database.AttemptHistory
	gateway                  Gateway
}

func NewPaymentService(
	attemptHistoryRepository database.AttemptHistory,
	gateway Gateway,
) *PaymentService {
	return &PaymentService{
		attemptHistoryRepository: attemptHistoryRepository,
		gateway:                  gateway,
	}
}

// ユーザーとクレジットカードの情報をしようして支払いが許可されたか判断するメソッド
func (c *PaymentService) IsAuthorized(user entity.User, creditCard entity.CreditCard) (bool, error) {
	// ユーザーの前回の失敗した試行の回数を取得
	totalAttempts, err := c.attemptHistoryRepository.CountFailures(user)
	if err != nil {
		return false, err
	}

	// 失敗した試行の合計が5回より大きい場合、falseを返す
	if totalAttempts > 5 {
		return false, nil
	}

	// 5回未満の場合はgatewayを呼び出して、承認を行う
	isAuthorized, err := c.gateway.IsAuthorized(user, creditCard)
	if err != nil {
		return false, err
	}

	// 承認失敗した場合は、失敗試行の回数を増やす
	if !isAuthorized {
		err := c.attemptHistoryRepository.IncrementFailure(user)
		if err != nil {
			return false, err
		}
	}

	// isAuthorizedのフラグと必要に応じてエラーを返す
	return isAuthorized, nil
}
