package main

import "fmt"

type Employee struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

func main() {
	var employee1 Employee

	// employee1描述
	employee1.ID = 10001
	employee1.Name = "Tom"
	employee1.Address = "xxxxx"
	employee1.Phone = "1349302xxxx"

	// 打印employee1信息
	printEmployee(&employee1)

}

func printEmployee(employee *Employee) {
	fmt.Printf("employee ID: %d\n", employee.ID)
	fmt.Printf("employee Name: %s\n", employee.Name)
	fmt.Printf("employee Address: %s\n", employee.Address)
	fmt.Printf("employee Phone: %s\n", employee.Phone)
}
