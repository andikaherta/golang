package models

type Todo struct {
	Id        int
	Item      string
	Assignee  string
	Deadline  string
	Completed int
}
