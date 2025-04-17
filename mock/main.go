package main

import (
	"fmt"
	"strconv"

	"example.com/mock-catalog-item/catalogItem"
	// "example.com/mock/catalogItem"
)

func main() {
	fmt.Println("Hello, World!")

	fieldA := getUserInput("Field A: ")
	fieldB := getUserInput("Field B: ")
	fieldC := getUserInput("Field C: ")
	fieldD := getUserInput("Field D: ")
	fieldDInt, err := strconv.Atoi(fieldD)
	if err != nil {
		fmt.Println(err)
	}
	fieldE := getUserInput("Field E: ")

	userCatalogItem, err := catalogItem.New(fieldA, fieldB, fieldC, fieldDInt, fieldE)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userCatalogItem)
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
