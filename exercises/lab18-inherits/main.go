package main

import "fmt"

type Student struct {
	name  string
	age   int
	score int
}

type Pupil struct {
	Student
	Level string
}

type Graduated struct {
	Student
	Level string
}

// common func
func (s *Student) Profile() {
	fmt.Printf("student = %s age = %d score = %d\n", s.name, s.age, s.score)
}

func (s *Student) SetScore(score int) {
	s.score = score
}

// Pupil func
func (p *Pupil) Testing() {
	fmt.Printf("[Student level: %s] %s have a test of first-exam.\n", p.Level, p.name)
}

// Graduated func
func (g *Graduated) Writing() {
	fmt.Printf("[Student level: %s] %s is writing paper.\n", g.Level, g.name)
}

func main() {
	pupil := &Pupil{
		Student: Student{name: "Tom", age: 8, score: 95},
		Level:   "pupil",
	}

	graduated := &Graduated{
		Student: Student{name: "Jackie", age: 22, score: 88},
		Level:   "graduated",
	}

	// before setScore
	pupil.Profile()
	graduated.Profile()

	fmt.Println()
	// setScore
	pupil.SetScore(90)
	graduated.SetScore(99)
	// call func separately
	pupil.Testing()
	graduated.Writing()
	
	fmt.Println()
	// after setScore
	pupil.Profile()
	graduated.Profile()
}
