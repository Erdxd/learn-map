package main

import (
	"fmt"
)

func main() {
	m := map[string]string{}
menu:
	for {

		user := ""
		fmt.Println("Что хотите сделать?")
		fmt.Scan(&user)
		if user == "1" {
			Printm(m)

		} else if user == "2" {
			m = addbookmark(m)

		} else if user == "3" {
			c := ""
			fmt.Println("Название закладки которую хотите удалить")
			fmt.Scan(&c)
			delete(m, c)

		} else if user == "4" {
			break menu
		}

		for key, value := range m {
			fmt.Println(key, "=", value)

		}

	}

}
func addbookmark(m map[string]string) map[string]string {
	var a string
	var b string

	fmt.Println("Название закладки")
	fmt.Scan(&a)
	fmt.Println("Адрес закладки")
	fmt.Scan(&b)
	m[a] = b
	return m

}
func Printm(m map[string]string) {
	if len(m) == 0 {
		fmt.Println("У вас еще нету закладок")
	}

}
