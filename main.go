package main

import "fmt"

var a, b, c int //= 6, 7, 8
var name string

func main() {
	name = "Ivan"
	mess, _ := enter_club(16)
	fmt.Println(mess)

}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello, %s -  %d лет ", name, age)
}

func enter_club(age int) (string, bool) {
	if age >= 18 {
		mess := fmt.Sprintf("%s вам доступен вход в клуб,ведь вам %d лет ", name, age)
		return mess, true

	} else {
		mess := fmt.Sprintf("%s вам недоступен вход в клуб,ведь вам всего %d лет ", name, age)
		return mess, false

	}
}
