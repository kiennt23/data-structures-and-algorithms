package main

import "fmt"

type IComposite interface {
	perform()
}

type Leaflet struct {
	name string
}

func (leaf Leaflet) perform() {
	fmt.Println("Leaflet " + leaf.name)
}

type Branch struct {
	name     string
	leaves   []Leaflet
	branches []Branch
}

func (branch Branch) perform() {
	fmt.Println("Branch " + branch.name)
	for _, childBranch := range branch.branches {
		childBranch.perform()
	}
	for _, leaf := range branch.leaves {
		leaf.perform()
	}
}

func (branch *Branch) addLeaf(leaf Leaflet) {
	branch.leaves = append(branch.leaves, leaf)
}

func (branch *Branch) addBranch(childBranch Branch) {
	branch.branches = append(branch.branches, childBranch)
}

func (branch *Branch) getLeaves() []Leaflet {
	return branch.leaves
}

func main() {
	branch := &Branch{name: "Root"}
	childBranch := Branch{name: "Child Branch"}
	branch.addBranch(childBranch)
	leaf1 := Leaflet{name: "Leaf 1"}
	branch.addLeaf(leaf1)
	leaf2 := Leaflet{name: "Leaf 2"}
	branch.addLeaf(leaf2)
	branch.perform()
}
