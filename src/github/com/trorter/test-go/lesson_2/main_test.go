package main

import "testing"

func TestOfStructures(t *testing.T) {
	person := Person{
		Name: "Ivan",
		Age:  18,
		Sex:  "male",
	}

	person.Grow(10)
	if person.Age != 28 {
		t.Error("Expected 28, got ", person.Age)
	}

	person.Print()

	superHuman := SuperHuman{
		Id: Id{
			Id: 1,
		},
		Person: Person{
			Name: "Superman",
			Age:  30,
			Sex:  "female",
		},
		Gifts: []string{"Fly", "Laser"},
	}

	superHuman.Print()

	fakeStuff := fakeStuff(1)
	if text := fakeStuff.doStuff(); text != "I'm fake" {
		t.Error("Expected \"I'm fake\", got ", text)
	}

	realStuff := realStuff("realStuff")
	if text := realStuff.doStuff(); text != "realStuff" {
		t.Error("Expected \"realStuff\", got ", text)
	}
}

func TestOfMain(t *testing.T) {

	main()
}
