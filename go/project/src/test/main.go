package main

import (
	"fmt"
	"time"
)


type Person struct {
	FirstName, LastName string
	Dob time.Time
	Email, Location string
}
func (p *Person) ChangeLocation(newLocation string) {
	p.Location= newLocation
}
func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}
//A person method
func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", p.Dob.String(), p.Email,
		p.Location)
}


func main()  {
	p := Person{
		"Shiju",
		"Varghese",
		time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
		"shiju@email.com",
		"Kochi",
	}
	p.ChangeLocation("Santa Clara")
	p.PrintName()
	p.PrintDetails()
}
