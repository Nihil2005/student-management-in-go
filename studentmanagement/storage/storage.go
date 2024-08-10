package storage

import (
	"errors"
	"studentmanagement/student"
)

type StudentStorage struct {
	students map[int]student.Student
}

func NewStudentStorage() *StudentStorage {
	return &StudentStorage{
		students: make(map[int]student.Student),
	}
}

func (s *StudentStorage) AddStudent(stu student.Student) {
	s.students[stu.ID] = stu
}

func (s *StudentStorage) GetStudentByID(id int) (student.Student, error) {
	stu, exists := s.students[id]
	if !exists {
		return student.Student{}, errors.New("student not found")
	}
	return stu, nil
}

func (s *StudentStorage) GetAllStudents() []student.Student {
	students := []student.Student{}
	for _, stu := range s.students {
		students = append(students, stu)
	}
	return students
}

func (s *StudentStorage) UpdateStudent(stu student.Student) {
	s.students[stu.ID] = stu
}

func (s *StudentStorage) DeleteStudent(id int) {
	delete(s.students, id)
}
