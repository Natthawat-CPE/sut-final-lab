package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameEmployeeNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	emp := Employee{
		Name:       "",
		Email:      "aha0037@gmail.com",
		EmployeeID: "S12345678",
	}

	ok, err := govalidator.ValidateStruct(emp)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Name cannot be blank"))
}
