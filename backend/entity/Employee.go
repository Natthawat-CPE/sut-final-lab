package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Name cannot be blank"`
	Email      string
	EmployeeID string `valid:"matches(^[JMS][0-9]{8}$)~ต้องขึ้นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจำนวน 8 ตัว"`
}
