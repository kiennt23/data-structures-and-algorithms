package main

import "fmt"

type Processor struct {
}

func (processor Processor) process() {
	fmt.Println("Processor processing")
}

type ProcessDecorator struct {
	processor *Processor
}

func (processDecorator ProcessDecorator) process() {
	fmt.Println("ProcessDecorator processing")
	if processDecorator.processor != nil {
		processDecorator.processor.process()
	}
}

func main() {
	processor := &Processor{}
	processDecorator := &ProcessDecorator{processor: processor}
	processDecorator.process()
}
