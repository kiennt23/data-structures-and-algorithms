package main

import "fmt"

type IRealObject interface {
	performAction()
}

type RealObject struct {
}

func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performing action")
}

type VirtualProxy struct {
	realObject IRealObject
}

// Need to pass a pointer to a VirtualProxy
func (virtualProxy *VirtualProxy) performAction() {
	fmt.Println("VirtualProxy performing action")
	if virtualProxy.realObject == nil {
		fmt.Println("Creating real object")
		virtualProxy.realObject = &RealObject{}
	}
	virtualProxy.realObject.performAction()
}

func main() {
	virtualProxy := VirtualProxy{}
	virtualProxy.performAction()
	virtualProxy.performAction()
}
