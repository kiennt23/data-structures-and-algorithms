package main

import "fmt"

type IProcessor interface {
	process()
}

type Adaptee struct {
	adapterType string
}

func (adaptee Adaptee) convert()  {
	fmt.Println("Adaptee convert method")
}

type Adapter struct {
	adaptee Adaptee
}

func (adapter Adapter) process()  {
	fmt.Println("Adapter process method")
	adapter.adaptee.convert()
}

func main() {
	var processor IProcessor = Adapter{}
	processor.process()
}


