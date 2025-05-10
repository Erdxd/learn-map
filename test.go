package main

import (
	"fmt"
)

func main() {
	z := ""
	for {

		m := map[string]string{}
		user := ""
		fmt.Println("Что хотите сделать?")
		fmt.Scan(&user)
		if user == "1" {
			for key, value := range m {
				fmt.Println(key, value)

			}

		} else if user == "2" {
			a := ""
			b := ""
			fmt.Println("Название закладки")
			fmt.Scan(&a)
			fmt.Println("Адрес закладки")
			fmt.Scan(&b)
			m["a"] = b

		} else if user == "3" {
			c := ""
			fmt.Println("Название закладки которую хотите удалить")
			fmt.Scan(&c)
			delete(m, c)

		} else if user == "4" {

		}
		fmt.Println("Хочешь продолжить? да/нет")
		fmt.Scan(&z)
		if z == "нет" {
			break
		}
	}

}
