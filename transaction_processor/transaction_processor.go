package transaction_processor

import "time"

type TransactionData struct {
	AccountFrom       int64
	AccountTo         int64
	CreateData        time.Time
	CurrencyCode      int
	Amount            float64
	TransactionType   int
	TransactionStatus int
}

type Kafka interface {
	GetTransactionForProcess() (td TransactionData, err error)
	PutMessageToNotification(transactionFrom, transactionTo TransactionData)
}

type DB interface {
	PutTransaction(transactionFrom, transactionTo TransactionData) error
}

type process struct {
	k  Kafka
	db DB
}

func (p *process) ProcessTransaction() {
	for {
		td, _ := p.k.GetTransactionForProcess()
		transactionFrom, transactionTo := PrepareTransaction(td)
		err := p.db.PutTransaction(transactionFrom, transactionTo)
		if err != nil {
			p.k.PutMessageToNotification(transactionFrom, transactionTo)
		}

	}
}

func PrepareTransaction(td TransactionData) (transactionFrom TransactionData, transactionTo TransactionData) {
	return
}
