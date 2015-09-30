package main

import "fmt"

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */
type Person struct {
	Name string
	Age int
	Weight int
}

type Student struct {
	Person
	Skills []string
	Speciality string
}

func main() {
	dimas := Student{Person: Person{"Dimas Ragil T", 23, 50}, Speciality: "Backend and Mobile"}

	fmt.Printf("Name : %v\nAge : %v\nWeight : %v\nSpeciality : %v\n", dimas.Name, dimas.Age, dimas.Weight, dimas.Speciality)

	dimas.Skills = append(dimas.Skills, "Java", "Scala", "Android", "Go")
	fmt.Println("Skills :", dimas.Skills)
	fmt.Println("Total Skill :", len(dimas.Skills))

	for _, v := range dimas.Skills {
		fmt.Println(v)
	}
}
