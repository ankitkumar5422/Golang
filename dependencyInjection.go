package main

import (
	"go/format"
	"log"
)

type Database interface {
	GetUser() string
	GetAllUsers() [] string
}

type DefaultDatabase struct{}

func (db DefaultDatabase) GetUser() string {
	return "bob"
}

type FamousDatabase struct{}

func (db FamousDatabase) GetUser() string {
	return "Will smith"
}
func (db FamousDatabase) GetAllUsers() [] string {
	return []string{}
}

type Greeter interface {
	Greet(userName string)
}

type NiceGreeter struct{}

func (ng NiceGreeter) Greet(userName string) {
	log.Printf(format:"Hello %s !! Nice to meet you", userName)
}

type Program struct{
	Db Database
	Greeter Greeter
}
func (p Program) Execute(){
	user :=p.Db.GetUser()
	p.Greeter.Greet(user)
}

func main(){
	db := FamousDatabase{}
	Greeter := NiceGreeter{}
	p := Program{
		Db:  db,
		Greeter: greeter,
	}

	p.Execute()
}