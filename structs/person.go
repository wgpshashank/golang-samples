package main
import (
	"fmt"
)

type People interface {
	SayHello()
}

type Person struct {
	name string
	age int
	city,phone string
}
//A person method
func (p Person ) SayHello() {
	fmt.Printf("Hi, I am %s, %d years old\n", p.name, p.age)
}

func main() {
	var shiju Person
	shiju.name="ShijuVar"
	shiju.age=30
	shiju.city="Kochi"
	shiju.phone="+91-94003372xx"

	shaheer:= Person {
		name : "Shaheer",
		age : 18,
		city : "Kochi",
		phone : "+91-94003372xx",
	}

	thomson:= Person {"Thomson",45,"Kochi","+91-94003372xx"}
	peopleArr := []People{shiju, shaheer,thomson}

	//Iterating through the array of People types and call methods.
	for k, _ := range peopleArr {
		peopleArr[k].SayHello()
	}
}
