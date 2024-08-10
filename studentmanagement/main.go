package main

import (
	"fmt"
	"studentmanagement/storage"
	"studentmanagement/student"
)

func main() {
	// Initialize storage
	s := storage.NewStudentStorage()

	// Create students
	s.AddStudent(student.Student{ID: 1, Name: "Alice", Age: 20})
	s.AddStudent(student.Student{ID: 2, Name: "Bob", Age: 22})

	// List all students
	fmt.Println("All students:")
	students := s.GetAllStudents()
	for _, stu := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", stu.ID, stu.Name, stu.Age)
	}

	// Find a student by ID
	fmt.Println("\nFinding student with ID 1:")
	stu, err := s.GetStudentByID(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", stu.ID, stu.Name, stu.Age)
	}

	// Update student information
	s.UpdateStudent(student.Student{ID: 2, Name: "Bob Updated", Age: 23})
	fmt.Println("\nUpdated student list:")
	students = s.GetAllStudents()
	for _, stu := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", stu.ID, stu.Name, stu.Age)
	}

	// Delete a student
	s.DeleteStudent(1)
	fmt.Println("\nStudent list after deletion:")
	students = s.GetAllStudents()
	for _, stu := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", stu.ID, stu.Name, stu.Age)
	}
}
