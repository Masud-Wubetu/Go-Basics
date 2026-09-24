package main

import "fmt"

type Payable interface {
	fmt.Stringer
	CalculatePay() float64
}

type salariedEmployee struct {
	Name string
	AnualSalary float64
}

func (se salariedEmployee) CalculatePay() float64 {
	return se.AnualSalary / 12.0
}

func (se salariedEmployee) String() string {
	return fmt.Sprintf("salaried: %s (Anual: $%.2f)", se.Name, se.AnualSalary)
}

type HourlyEmployee struct {
	Name string
	HourlyRate float64
	HoursWorked float64
}

func (he HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %.1f)", he.Name, he.HourlyRate, he.HoursWorked)
}

type CommissionEmployee struct {
	Name string
	BaseSalary float64
	CommissionRate float64
	SalesAmount float64
}

func (ce CommissionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommissionRate * ce.SalesAmount)
}

func (ce CommissionEmployee) String() string {
	return fmt.Sprintf("Commission: %s (Base: $%.2f, CommRate: %.2f%%, Sales: $%.2f)",
		ce.Name, ce.BaseSalary, ce.CommissionRate * 100, ce.SalesAmount)
}

func PrintEmployeeSummary[P fmt.Stringer](employee P) {
	fmt.Printf("---Processing: %s\n", employee)
}

func ProcessPayroll(employees []Payable) {
	fmt.Println("\n--- Processing Payroll ---")
	totalPayroll := 0.0
	for _, emp := range employees {
		PrintEmployeeSummary(emp)
		pay := emp.CalculatePay()
		fmt.Printf("Monthly pay: $%.2f\n", pay)
		totalPayroll += pay
	}

	fmt.Printf("Total monthly payroll: $%.2f\n", totalPayroll)
	fmt.Println("----------------")
}

func main() {

	fmt.Println("Welcome to the payroll processor")
	salEmp := salariedEmployee{Name: "Masud Wubetu", AnualSalary: 72000.0}
	hrEmp := HourlyEmployee{Name: "Miftah Ebrahim", HourlyRate: 25.00, HoursWorked: 160.0}
	commEmp := CommissionEmployee{Name: "Khalid Suabir", BaseSalary: 2000.00, CommissionRate: 0.10, SalesAmount: 15000.0}

	payrollList := []Payable{
		salEmp,
		hrEmp,
		commEmp,
		HourlyEmployee{Name: "Diana Prince", HourlyRate: 30.00, HoursWorked: 150.0},
	}

	ProcessPayroll(payrollList)

}
