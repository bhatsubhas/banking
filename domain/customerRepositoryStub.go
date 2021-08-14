package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (cs CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return cs.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Subhas", "Udupi", "576226", "15/08/1987", "1"},
		{"1002", "Suresh", "Mangaluru", "573226", "15/08/1977", "1"},
		{"1003", "Harish", "Kumble", "666226", "15/08/1967", "1"},
		{"1004", "Mahesh", "Bengaluru", "560226", "15/08/1987", "1"},
	}
	return CustomerRepositoryStub{customers}
}
