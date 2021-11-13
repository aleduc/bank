package customer

type CustomerData struct {
	Name, Surname, MobileNumber, Address, Login, Password string
	ID                                                    int64
}

type DB interface {
	CreateCustomer(cd CustomerData) (id int64, err error)
}

type Redis interface {
	Put(customerID int64, token string) (err error)
	Get(customerID int64) (err error)
	Delete(customerID int64) (err error)
}

func Create(cd CustomerData) (token string, err error) {
	return
}

func Login(cd CustomerData) (token string, err error) {
	return
}

func VerifyToken(customerID int64, token string) (err error) {
	return
}

func Logout(customerID int64, token string) (err error) {
	return
}
