package consts

import (
	"errors"
)

var (
	// 错误码
	Success = 0
	Error   = 1

	UserId = "userId"

	TransactionTypePut      = "存款"
	TransactionTypeTake     = "取款"
	TransactionTypeTransfer = "转账"
	TransactionTypePay      = "支付"

	// BankNameNum 银行名称标识
	BankNameNum = map[string]string{
		"工商银行":   "6222",
		"农业银行":   "6214",
		"建设银行":   "6212",
		"交通银行":   "6216",
		"邮政储蓄银行": "6210",
		"中国银行":   "6213",
		"招商银行":   "6211",
		"中信银行":   "6221",
		"广发银行":   "6222",
		"光大银行":   "6223",
		"华夏银行":   "6225",
		"民生银行":   "6226",
		"平安银行":   "6228",
		"兴业银行":   "6229",
	}
)

var (
	ErrBankNameNotFound    = errors.New("不支持该银行")
	ErrBalanceNotEnough    = errors.New("余额不足")
	ErrInsufficientBalance = errors.New("余额不足")
	ErrTransferToSelf      = errors.New("不能转账给自己")
)
