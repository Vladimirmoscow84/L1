package main

import "fmt"

/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.*/

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
	Power int
}

func (h Human) getAge() int {
	return h.Age
}
func (h Human) introduce() string {
	return fmt.Sprintf("my name is %s", h.Name)
}

func main() {
	h := Human{Name: "Tod", Age: 25}
	a := Action{Human: h, Power: 30}

	fmt.Println("Hi,", a.introduce())
	fmt.Println("My age is", a.getAge())
	fmt.Println("My power is", a.Power)

}
