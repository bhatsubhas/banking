package service

import (
	"github.com/bhatsubhas/banking/domain"
	"github.com/bhatsubhas/banking/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (cs DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return cs.repo.FindAll()
}

func (cs DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return cs.repo.ById(id)
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
