package model


import (
	"testing"
)

func TestAddEmployee(t *testing.T) {
	company := NewCompany("Zeta LTD", "00011122222244", "zeta.ltd@email.com")
	employee := makeEmployee(100, "John Doe", "0000111222", "john.doe@email.com")

	company.addEmloyee(employee)

	ans := company.getEmloyee(employee.GetId())
	want := employee.GetId()
	if ans.GetId() != 100 {
		t.Errorf("got %d, want %d", ans.GetId(), want)
	}
}
