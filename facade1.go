package main

import "fmt"

type Entity interface {
	Id() string
}

type Account struct {
	id      string
	accType string
}

func (account Account) Id() string {
	return account.id
}

type Customer struct {
	id   string
	name string
}

func (customer Customer) Id() string {
	return customer.id
}

type Transaction struct {
	id       string
	amount   float64
	srcAccId string
	dstAccId string
}

func (transaction Transaction) Id() string {
	return transaction.id
}

type EntityRepository struct {
	data []Entity
}

func (entityRepository *EntityRepository) create(entity Entity) *Entity {
	fmt.Println("Creating entity with id " + entity.Id())
	entityRepository.data = append(entityRepository.data, entity)
	return &entity
}

func (entityRepository *EntityRepository) getById(id string) *Entity {
	fmt.Println("Getting entityRepository by Id " + id)
	for _, entity := range entityRepository.data {
		if entity.Id() == id {
			return &entity
		}
	}
	return nil
}

func (entityRepository *EntityRepository) deleteById(id string) *Entity {
	fmt.Println("Delete entityRepository by id" + id)
	for idx, entity := range entityRepository.data {
		if entity.Id() == id {
			entityRepository.data = append(entityRepository.data[:idx], entityRepository.data[idx+1:]...)
			return &entity
		}
	}
	return nil
}

type BranchManagerFacade struct {
	accountRepo     *EntityRepository
	customerRepo    *EntityRepository
	transactionRepo *EntityRepository
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&EntityRepository{}, &EntityRepository{}, &EntityRepository{}}
}

func (facade BranchManagerFacade) createCustomerAccount(id, name, accType string) (*Customer, *Account) {
	customer := &Customer{id: id, name: name}
	account := &Account{id: id, accType: accType}
	facade.customerRepo.create(customer)
	facade.accountRepo.create(account)
	return customer, account
}

func (facade BranchManagerFacade) createTransaction(srcAccId, dstAccId string, amount float64) *Transaction {
	transaction := &Transaction{"1_1", amount, srcAccId, dstAccId}
	facade.transactionRepo.create(transaction)
	return transaction
}

func main() {
	facade := NewBranchManagerFacade()
	_, account1 := facade.createCustomerAccount("1", "Sushant", "default")
	_, account2 := facade.createCustomerAccount("2", "Kien", "default")
	transaction := facade.createTransaction(account1.Id(), account2.Id(), 1000000)
	fmt.Printf("%v", transaction)
}
