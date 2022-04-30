package model

import (
	"gorm.io/gorm"
)

const (
	todomvc_status_active    = 0
	todomvc_status_completed = 1
)

type ToDoMvc struct {
	gorm.Model
	Item   string
	Status uint
}

type TodomvcAdd struct {
	Item string
}

type TodomvcDel struct {
	Id uint
}

type TodomvcUpdate struct {
	Item   string
	Id     uint
	Status uint
}

type TodomvcFind struct {
	Item   string
	Status uint
}
