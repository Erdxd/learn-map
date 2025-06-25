package main

import (
	"fmt"
)

type mmap = map[string]string

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

	}

}
func addbookmark(m mmap) mmap {
	var a string
	var b string

	fmt.Println("Название закладки")
	fmt.Scan(&a)
	fmt.Println("Адрес закладки")
	fmt.Scan(&b)

	for {
		if len(b) > 7 {

			if b[:8] == "https://" || b[:7] == "http://" {
				m[a] = b
				return m

			} else if b == "4" || a == "4" {
				break

			} else {
				fmt.Println("Введите URL (начинается с http:// или с https://). Попробуйте снова ")
				break

			}
		} else {
			fmt.Println("Введите URL (начинается с http:// или с https://). Попробуйте снова ")
			break
		}
	}

	return m

}
func Printm(m mmap) {
	if len(m) == 0 {
		fmt.Println("У вас еще нету закладок")
	} else {
		for key, value := range m {
			fmt.Println(key, "=", value)

		}
	}

}
