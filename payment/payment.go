package payment

import "time"

type TransactionData struct {
	AccountFrom       int64
	AccountTo         int64
	CreateData        time.Time
	ActualTime        time.Time
	CurrencyCode      int
	Amount            float64
	TransactionType   int
	TransactionStatus int
}

type Kafka interface {
	PutTransaction(td TransactionData) (err error)
}

func PutTransaction(accountFrom, accountTo int64, amount float64, currencyCode int) error {
	return nil
}
