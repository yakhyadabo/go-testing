package model

type Employee struct {
	id    int
	name  string
	phone string
	email string
}

func makeEmployee(id int, name, phone, email string) Employee { // Return Value
 return Employee{
 	id: id,
 	name: name,
	phone: phone,
	email: email,
 }
}

func (employee Employee) GetId() int {
	return employee.id
}
func (employee Employee) GetName() string {
	return employee.name
}
func (employee Employee) GetPhone() string {
	return employee.phone
}
func (employee Employee) GetEmail() string {
	return employee.email
}
func (employee *Employee) SetId(id int) {
	employee.id = id
}
func (employee *Employee) SetName(name string) {
	employee.name = name
}
func (employee *Employee) SetPhone(phone string) {
	employee.phone = phone
}
func (employee *Employee) SetEmail(email string) {
	employee.email = email
}