package main

import "fmt"

type DataTransferObject interface {
	Id() string
}

type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

type CustomerDTO struct {
	id   string
	name string
	ssn  string
}

func (customer CustomerDTO) Id() string {
	return customer.id
}

type EmployeeDTO struct {
	id   string
	name string
}

func (employee EmployeeDTO) Id() string {
	return employee.id
}

type ManagerDTO struct {
	id   string
	name string
	dept string
}

func (manager ManagerDTO) Id() string {
	return manager.id
}

type AddressDTO struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

func (address AddressDTO) Id() string {
	return address.id
}

func (dtoFactory DataTransferObjectFactory) getDataTransferObjectFactory(dtoType string) DataTransferObject {
	dto := dtoFactory.pool[dtoType]
	if dto == nil {
		switch dtoType {
		case "customer":
			dtoFactory.pool[dtoType] = CustomerDTO{id: "1"}
		case "employee":
			dtoFactory.pool[dtoType] = EmployeeDTO{id: "2"}
		case "manager":
			dtoFactory.pool[dtoType] = ManagerDTO{id: "3"}
		case "address":
			dtoFactory.pool[dtoType] = AddressDTO{id: "4"}
		}
		dto = dtoFactory.pool[dtoType]
	}
	return dto
}

func main() {
	factory := DataTransferObjectFactory{make(map[string]DataTransferObject)}
	customer := factory.getDataTransferObjectFactory("customer")
	fmt.Printf("%v", customer)
	employee := factory.getDataTransferObjectFactory("employee")
	fmt.Printf("%v", employee)
	manager := factory.getDataTransferObjectFactory("manager")
	fmt.Printf("%v", manager)
	address := factory.getDataTransferObjectFactory("address")
	fmt.Printf("%v", address)
}
