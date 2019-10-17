package main

import (
	"fmt"
	_ "demopack"
)

func main()  {
	//sf := demopack.Area(12,3) //引用自定义包
	//fmt.Println(sf)

	numss := []int {12,3,12,55,3}
	argArr(12,numss...)

	fmt.Println("dddd")
}

//switch
func sws()  {
	ms := `fe`
	fmt.Println(ms)

	ks := '2' //在Go中单引号里的内容是属于int32型 （rune型）。但是单引号里只允许有一个字符。不知道为什么这样设计。
	//lms := 'g'
	fmt.Println(ks)

	fi := 5
	switch fi {
	case 1:
		fmt.Println("1111")
	case 2,3,4:
		fmt.Println("22222")
	case 5:
		fmt.Println("5555")
		fallthrough  //继续执行下一个
	case 6:
		fmt.Println("66666")
		fallthrough
	case 7:
		fmt.Println("7777")
	default:
		fmt.Println("all")
	}
}

//for
func fores()  {
	for i:= 1; i <= 10; i++{
		if i == 2{
			continue
		}

		if i >= 5{
			break
		}

		fmt.Println(i)
	}


	k := 0
	for k <= 10 {
		fmt.Printf("%d ", k)
		k += 2
	}

	for no, i := 1,1; i <= 10 && no <= 8; i,no = i+1, no+ 1{
		fmt.Println(no,i)
	}
}

//函数
func demoInit(print, no int) int {
	var totle = print * no
	return totle
}

func resTwo(le, wh int) (int,int)  {
	resI := le * wh
	ps := le + wh

	return resI, ps
}

func resWeoV2(le,wh int) (are, per int)  {
	are = le * wh
	per = le + wh
	return
}

//可变参数函数
//可变参数函数的工作原理是把可变参数转换为一个新的切片
//argArr(44,43,22,456,232)
/**
	或者
numss := []int {12,3,12,55,3}
有一个可以直接将切片传入可变参数函数的语法糖，你可以在在切片后加上 ... 后缀。如果这样做，切片将直接传入函数，不再创建新的切片
argArr(12,numss...)
 */
func argArr(num int, nums ...int) {
	fmt.Println(num)

	for i,v := range nums{
		fmt.Println(i,v)
	}
}

