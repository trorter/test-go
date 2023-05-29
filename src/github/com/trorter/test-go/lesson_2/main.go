package main

import "fmt"

type Person struct {
	Name string `json:"name",xml:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
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

type iStuff interface {
	doStuff() string
}

type realStuff string

func (r realStuff) doStuff() string {
	return string(r)
}

type fakeStuff int

func (f fakeStuff) doStuff() string {
	return "I'm fake"
}

type complexStuff struct {
	iStuff
	name string
}

func (c complexStuff) doComplexStuff() {
	fmt.Println(c.name, "=", c.doStuff())
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

	println("=========")

	fakeStuff := fakeStuff(1)
	realStuff := realStuff("realStuff")

	complexFake := complexStuff{
		iStuff: fakeStuff,
		name:   "fakeStuff",
	}

	complexReal := complexStuff{
		iStuff: realStuff,
		name:   "realStuff",
	}

	complexFake.doComplexStuff()
	complexReal.doComplexStuff()
}
