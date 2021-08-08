package encapsulation

import "testing"

type Employee struct {
	ID   string
	Name string
	Age  int
}

func TestCreatedEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee)
	e2.ID = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Logf("e is %T", e)
	t.Log(e)
	t.Logf("e1 is %T", e1)
	t.Log(e1)
	t.Log(e1.ID)
	t.Logf("e2 is %T", e2)
	t.Log(e2)
}
