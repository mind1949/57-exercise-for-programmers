package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var employees = []string{
	"zhao",
	"qian",
	"sun",
	"li",
	"zhou",
}

func outputEmployees() {
	count := len(employees)
	fmt.Printf("There are %d employees.\n", count)
	for _, employee := range employees {
		fmt.Println(employee)
	}
}

func prompt(msg string) (input string) {
	if !strings.HasSuffix(msg, " ") {
		msg += " "
	}
	fmt.Printf(msg)

	_input := bufio.NewScanner(os.Stdin)

	if _input.Scan() {
		input = _input.Text()
	} else {
		return prompt(msg)
	}

	return
}

func getEmployee() (employeeName string) {
	_employeeName := prompt("enter an employee name to remove: ")

	if !isContain(employees, _employeeName) {
		fmt.Println("Sorry, this employee is not exist.(please enter again)")
		return getEmployee()
	}

	employeeName = _employeeName
	return
}

var findIndex = func(s []string, element string) (index int) {
	index = -1
	for _index, value := range s {
		if value == element {
			index = _index
			break
		}
	}
	return
}

var isContain = func(slice []string, element string) bool {
	index := findIndex(slice, element)
	if index > 0 {
		return true
	}
	return false
}

var remove = func(s []string, index int) []string {
	copy(s[index:], s[index+1:])
	return s[:len(s)-1]
}

func deleteEmployee(employee string) {
	employeeIndex := findIndex(employees, employee)
	employees = remove(employees, employeeIndex)
}

func main() {
	outputEmployees()

	employeeName := getEmployee()
	deleteEmployee(employeeName)

	fmt.Println("")

	outputEmployees()
}
