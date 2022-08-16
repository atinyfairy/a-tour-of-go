package main

import (
	"fmt"
	"math/cmplx"
	"runtime"
)

//bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64
// uintptr 这个东西是用来存放指针地址的，一般很少用
// byte // uint8 的别名

// rune int32 的别名 用于存放Unicode字节码
//     // 表示一个 Unicode 码点

// float32 float64

// complex64 complex128

/*

 */

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

type Flags uint

const (
	FlagUp Flags = 1 << (iota * 10)
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println(FlagUp)
	fmt.Println(FlagBroadcast)
	fmt.Println(FlagLoopback)
	fmt.Println(FlagPointToPoint)
	fmt.Println(FlagMulticast)

	os := runtime.GOOS
	fmt.Println(os)

}
