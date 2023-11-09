package main

import "fmt"

type Product interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

type Desktop struct {
	Computer
}

func newLaptop() Product {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

func newDesktop() Product {
	return &Desktop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

func GetComputerFactory(computerType string) (Product, error) {
	switch computerType {
	case "laptop":
		return newLaptop(), nil
	case "desktop":
		return newDesktop(), nil
	}
	return nil, fmt.Errorf("the computer type doesn't match with any type in this factory")
}

func printNameAndStock(p Product) {
	fmt.Printf("Product Details:\n Name: %s\n Stock: %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
