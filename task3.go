package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollno  int
	name    string
	address string
}

type StudentList struct {
	list []*Student
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollno = rollno
	s.name = name
	s.address = address
	return s
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

func p(limit int, sl *StudentList) {
	no := 0
	for i := 0; i < limit; i++ {
		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 25), no, strings.Repeat("=", 25))
		fmt.Printf("Student roll no \t %v\n", *&sl.list[i].rollno)
		fmt.Printf("Student name \t\t %v\n", *&sl.list[i].name)
		fmt.Printf("Student address \t %v\n", *&sl.list[i].address)
		no++
	}
}

func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Ali", "AAAA")
	student.CreateStudent(22, "Bilal", "BBBB")
	p(2, student)

}
