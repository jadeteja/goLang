package main

import "fmt"

func main(){
	var EmployeeName string
	EmployeeName = "Teja"
	var EmployeeID int
	EmployeeID = 3948;
	salary := 30000.50;
	isPermanent := true;
	const companyName = "Orienray";

	fmt.Printf("EmployeeName is %s\n",EmployeeName);
	fmt.Printf("Employee Id is %d\n", EmployeeID);
	fmt.Printf("Salary is %.2f\n", salary);
	fmt.Printf("Is Permanent %t", isPermanent);
	fmt.Printf("Company Name is %s \n", companyName);
}