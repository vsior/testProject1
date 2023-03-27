package aPackage

import (
	"fmt"
	"testProject1/aPackage/bPackage"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func init() {
	fmt.Println("init 3")
}

const MyConstant int = 110
const privateConstant int = 404

func A() {
	fmt.Println("aPackage A()!!!!")
}

func B() {
	fmt.Println("privateConstant:", privateConstant)
	bPackage.B()
}
