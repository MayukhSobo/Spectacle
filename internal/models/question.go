package models

import "spectacle/internal/types"

type Question struct {
	Ques string
	Ans  string
	Inp  types.Input
}

func NewQuestion(q string) Question {
	return Question{Ques: q}
}
