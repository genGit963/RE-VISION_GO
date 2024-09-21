package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func learn_Struct_Embedding() {
	fmt.Println("\n--------- Learn Struct Embedding ----------- ")

	co := container{base: base{num: 2}, str: "some str"}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)
	/*
		Since container embeds base, the methods of base also become methods
		 of a container. Here we invoke a method that was embedded from
		  base directly on co.
	*/
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer | interface: ", d, d.describe())

}

func main() {
	learn_Struct_Embedding()
	println()
}
