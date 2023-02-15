package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeIDNegativeLetter(t *testing.T) {
	g := NewGomegaWithT(t)

	emp := Employee{
		Name:       "นาถวัฒน์ สลางสิงห์",
		Email:      "aha0037@gmail.com",
		EmployeeID: "A12345678",
	}

	ok, err := govalidator.ValidateStruct(emp)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("ต้องขึ้นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจำนวน 8 ตัว"))
}

func TestEmployeeIDNegativeNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	emp := Employee{
		Name:       "นาถวัฒน์ สลางสิงห์",
		Email:      "aha0037@gmail.com",
		EmployeeID: "S123456789",
	}

	ok, err := govalidator.ValidateStruct(emp)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("ต้องขึ้นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจำนวน 8 ตัว"))
}
