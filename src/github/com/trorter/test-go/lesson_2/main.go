package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

type Id struct {
	Id   int
	Name string
}

type SuperHuman struct {
	Id
	Person
	Gifts []string
}

func (p Person) Print() {
	fmt.Println("{Name=", p.Name, ", age=", p.Age, ", sex=", p.Sex, "}")
}

func (p SuperHuman) Print() {
	fmt.Println("{Name=", p.Person.Name, ", age=", p.Age, ", sex=", p.Sex, "gifts=", p.Gifts, "}")
}

func (p *Person) Grow(years int) {
	p.Age += years
}

func main() {
	testPerson1 := Person{
		Name: "Test1",
		Age:  11,
	}

	testPerson1.Print()

	testSuperHuman := SuperHuman{
		Person: Person{
			Name: "Superman",
			Age:  30,
			Sex:  "male",
		},
		Id: Id{
			Id:   1,
			Name: "test",
		},
		Gifts: []string{"fly", "laser", "x-ray"},
	}

	testSuperHuman.Print()
}
