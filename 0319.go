package main

import (
	"fmt"
	"strconv"
)

// 定数(グローバル) =>一度定義した定数は変更不可
const Pi = 3.14

// 連続する整数の値を生成
const (
	c0 = iota
	c1
	c2
)

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

// 関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("function")
	}
}

// interface
func anything(a interface{}) {
	fmt.Println(a)
}

func main() {

	var x interface{} = 3
	e := x.(int)
	fmt.Println(e + 3)

	f := ReturnFunc()
	f()

	i5 := Double(100)
	fmt.Println(i5)

	ii, iii := Div(9, 3)
	fmt.Println(ii, iii)

	fmt.Println(c0, c1, c2)
	fmt.Println(Pi)

	var s string = "100"
	fmt.Printf("s = %T\n", s)
	// 文字列型から数値型
	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	var i2 int = 200
	fmt.Println(i2)
	fmt.Printf("i2 = %T\n", i2)

	// 数値型から文字列型
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	// 文字列型からbyte配列
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	// byte配列から文字列型
	h2 := string(b)
	fmt.Println(h2)

	// 文字列の結合
	s3 := "ABC"
	s3 += "DEF"
	fmt.Println(s3)

	s4 := "ABC" + "DEF"
	fmt.Println(s4)

	// if
	a := 2
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("else")
	}

	for m := 0; m < 10; m++ {
		if m == 3 {
			fmt.Println("continue")
			continue
		}
		if m == 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(m)
	}

	// for文(配列)
	arr := [3]int{1, 2, 3}
	for m2 := 0; m2 < len(arr); m2++ {
		fmt.Println(arr[m2])
	}

	arr2 := [3]string{"GO", "Python", "Html"}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

	// for文(スライス)
	sl := []string{"GO", "Python", "Html"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 300}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// switch
	/*
		n := 5
		switch n {
		case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		default:
			fmt.Println("default")
		}
	*/

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("default")
	}

	// ラベル付きfor
Loop:
	for {
		for {
			for {
				fmt.Println("START")
				break Loop
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("処理をしない")
	}
	fmt.Println("END")

	// jが1の時だけ処理したい
Loop:
	for i8 := 0; i8 < 3; i8++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i8, j, i8*j)
		}
		fmt.Println("処理をしない")
	}

}
