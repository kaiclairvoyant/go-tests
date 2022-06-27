package main

import "errors"

type Person struct {
	name     string
	lastName string
	age      int
}

const AgeOfMajority = 18
const ReallyOldAge = 150

func NewPerson(name string, lastName string, age int) (*Person, error) {
	if name == "invalid" {
		return nil, errors.New("name is invalid, literally")
	}
	if len(name) <= 1 {
		return nil, errors.New("name is too short")
	}
	if len(lastName) <= 1 {
		return nil, errors.New("last name is too short")
	}
	if age < AgeOfMajority {
		return nil, errors.New("user is under the age of majority")
	}
	if youAreProbablyAVampire(age) {
		return nil, errors.New("no vampires allowed")
	}

	var p Person
	p.name = name
	p.lastName = lastName
	p.age = age
	return &p, nil
}

func youAreProbablyAVampire(age int) bool {
	return age >= ReallyOldAge
}
