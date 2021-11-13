package account

type Data struct {
	ID            int64
	accountNumber string
	sortCode      int
}

type DB interface {
	CreateAccount(ad Data) (id int64, err error)
	DeleteAccount(customerID int64, accountID int64)
}

func CreateAccount(customerID int64, d Data) (err error) {
	return
}

func DeleteAccount(customerID int64, d Data) (err error) {
	return
}
