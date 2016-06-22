package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name   Name
	School School
}
type Name struct {
	FirstName string
	LastName  string
}

type School struct {
	City   string
	Region string
}

func main() {
	person := Person{
		Name:   Name{"Chen", "Xu"},
		School: School{"Hangzhou", "XiHu"}}
	saveJson("person.json", person)

}

func saveJson(filename string, key interface{}) {
	output, _ := os.Create(filename)
	encoder := json.NewEncoder(output)
	err := encoder.Encode(key)
}
