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
