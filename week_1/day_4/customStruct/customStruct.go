package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func orderByAge(persons []Person) []Person {
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Age > persons[j].Age
	})
	return persons
}

func makePerson(name string, age int) Person {
	p := Person{Name: name, Age: age}
	fmt.Println("Created new person with name", name, "and age", age)
	return p
}

func main() {
	original_persons := make([]Person, 0, 10)

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("Name_%d", i)
		original_persons = append(original_persons, makePerson(name, i+10))
	}
	fmt.Println("Created list of persons", original_persons)
	ordered_persons := orderByAge(original_persons)
	fmt.Println("Final list of persons", ordered_persons)

}
