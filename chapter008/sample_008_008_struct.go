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

	fmt.Printf("形式传参之前，employee ID: %d\n", employee1.ID)
	operateEmployee1(employee1)
	fmt.Printf("形式传参之后，employee ID: %d\n", employee1.ID) // ID值是不会变化的。

	fmt.Println("***********************************************")

	fmt.Printf("指针传参之前，employee ID: %d\n", employee1.ID)
	operateEmployee2(&employee1)
	fmt.Printf("指针传参之后，employee ID: %d\n", employee1.ID) // ID值是发生变化。

}

// 形式传参（形参传值）
func operateEmployee1(employee Employee) {
	employee.ID = 10010
}

// 指针传参（形参传址）
func operateEmployee2(employee *Employee) {
	employee.ID = 10020
}
