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
	fmt.Printf("employee1 ID: %d\n", employee1.ID)
	fmt.Printf("employee1 Name: %s\n", employee1.Name)
	fmt.Printf("employee1 Address: %s\n", employee1.Address)
	fmt.Printf("employee1 Phone: %s\n", employee1.Phone)
}
