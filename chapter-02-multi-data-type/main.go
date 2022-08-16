package main

import (
	"crypto/sha256"
	"fmt"
)

var a = [3]int{1, 2, 3}

func main() {
	for i, v := range a {
		fmt.Printf("%d %d\t", i, v)
	}
	for _, v := range a {
		fmt.Printf("%d \t", v)
	}
	fmt.Println("")
	fmt.Println("===========================================================")

	d1 := sha256.Sum256([]byte("x"))
	d2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%T\t%x\n", d1, d1)
	fmt.Printf("%T\t%x\n", d2, d2)
	fmt.Println("===========================================================")
	var runes []rune
	for _, r := range "hello world" {
		runes = append(runes, r)
	}
	fmt.Println("runes:", runes)
	fmt.Println("===========================================================")
	e := [32]byte{'x', 'z'}
	notChangeArray(e)
	fmt.Printf("%T\t%v\n", e, e)
	changeArray(&e)
	fmt.Printf("%T\t%v\n", e, e)
	fmt.Println("===========================================================")
	//创建了一个K（rune）V（int） map
	maps := map[string]int{
		"500": 1,
		"400": 2,
	}
	maps["aa"] = 3

}

func notChangeArray(arr [32]byte) {
	//go 传递的数组是值传递，相当于复制了一遍，原始的数组不会随着方法里对数组的修改而改变
	arr = [32]byte{}
}

func changeArray(ptr *[32]byte) {
	//传入地址的话就可以改变了（废话），这是一种高效的改变方式
	*ptr = [32]byte{}
}
