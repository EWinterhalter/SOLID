package main

import (
	"fmt"
)

// Single Responsibility Principle (SRP) Каждый класс должен иметь только одну ответственность.

type Student struct {
	ID    int
	Name  string
	Grade int
}

// Open/Closed Principle (OCP) Классы должны быть открыты для расширения, но закрыты для модификации.

type Printer interface {
	Print(student Student)
}

type ConsolePrinter struct{}

func (cp ConsolePrinter) Print(student Student) {
	fmt.Printf("ID: %d, Name: %s, Grade: %d\n", student.ID, student.Name, student.Grade)
}

// Dependency Inversion Principle (DIP) Модули верхнего уровня не должны зависеть от модулей нижнего уровня. Оба типа модулей должны зависеть от абстракций.

type StudentService struct {
	printer Printer
}

func (ss StudentService) PrintStudentInfo(student Student) {
	ss.printer.Print(student)
}

func main() {
	student := Student{ID: 1, Name: "Name", Grade: 20}
	printer := ConsolePrinter{}
	studentService := StudentService{printer: printer}

	studentService.PrintStudentInfo(student)
}
