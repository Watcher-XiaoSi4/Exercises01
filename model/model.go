package model

import (
	"gorm.io/gorm"
)

const (
	Exercises01_status_active    = 0
	Exercises01_status_completed = 1
)

type Todomvc struct {
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
