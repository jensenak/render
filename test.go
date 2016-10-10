package main

import (
	"fmt"

	"./obj"
)

func main() {
	fmt.Println("Hi")
	vert, face := obj.Parse("head.obj")
	fmt.Println(vert[0:3])
	fmt.Println(face[0:3])
}
