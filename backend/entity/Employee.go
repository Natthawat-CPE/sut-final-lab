package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Name cannot be blank"`
	Email      string
	EmployeeID string
}
