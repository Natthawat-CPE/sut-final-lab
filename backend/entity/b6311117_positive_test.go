package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestTrueFormat(t *testing.T) {
	g := NewGomegaWithT(t)

	emp := Employee{
		Name:       "Natthawat",
		Email:      "aha0037@gmail.com",
		EmployeeID: "S12345678",
	}

	ok, err := govalidator.ValidateStruct(emp)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
