/*Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}. */
package main

import (
	"fmt"
	"reflect"
)

func printType(i interface{}) {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int:
		fmt.Println("Переменная типа int")
	case reflect.String:
		fmt.Println("Переменная типа string")
	case reflect.Bool:
		fmt.Println("Переменная типа bool")
	case reflect.Chan:
		fmt.Println("Переменная типа channel")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	printType(10)
	printType("hello")
	printType(true)
	printType(make(chan int))
}
