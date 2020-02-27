package model

import "fmt"

// factory pattern & Encapsulation of fields
type student struct {
	name  string
	score float64
}

func (s *student) String() string {
	return fmt.Sprintf("name = %s, score = %.2f", s.name, s.score)
}

func (s *student) SetName(name string) {
	s.name = name
}

func (s *student) GetName() string {
	return s.name
}

func (s *student) SetScore(score float64) {
	s.score = score
}

func (s *student) GetScore() float64 {
	return s.score
}

func NewStudent(name string, score float64) *student {
	return &student{name, score}
}
