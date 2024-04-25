package model


type Company struct {
	name      string
	phone     string
	email     string
	employees map[int]Employee
}

func NewCompany(name, phone, email string) *Company { // Return Pointer (Statefull Struct)
	return &Company{
		name: name,
		phone: phone,
		email: email,
		employees: make(map[int]Employee),
	}
}

func (company *Company) addEmloyee(employee Employee) {
	company.employees[employee.GetId()] = employee
}

func (company *Company) getEmloyee(id int) Employee {
	return company.employees[id]
}

func (company *Company) removeEmloyee(id int) {
	delete(company.employees,id)
}

func (company *Company) SetName(name string) {
	company.name = name
}

func (company *Company) SetPhone(phone string) {
	company.phone = phone
}

func (company *Company) SetEmail(email string) {
	company.email = email
}