package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIds []uint
	Classes map[uint]string
}

/// Metodos

func PrintClasses(c Course) {
	text := "las clases son"
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

func (c Course) PrintClasses2() {
	text := "las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
