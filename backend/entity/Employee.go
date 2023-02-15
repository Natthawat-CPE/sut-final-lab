package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string
	Email      string
	EmployeeID string
}
