package main

import (
	"fmt"
	"runtime"
	"time"
	"unicode/utf8"
)

func main()  {

}

func vales(bs int) int {
	bs = 4
	return bs
}

func demoGo()  {
	//runtime.GOMAXPROCS(1)
	////并发
	//for i := 0; i < 10; i++ {
	//	go func(ks int) {
	//		println(ks)
	//	}(i)
	//}
	////runtime.Gosched()
	//time.Sleep(time.Second)


	//依次执行,只开了一个P
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go println(i)
	}
	//runtime.Gosched() //添加后会输出  9,0,1...8
	time.Sleep(time.Second)
}

//定义返回值必须是int64的函数
func Hash() int64 {
	s := 34
	return int64(s)
}

//变量.常量定义
func v3()  {
	const sd  = "fffff"
	var sm = 23.234

	fmt.Println(int(sm))
	fmt.Println(len(sd)) //字符串byte长度
	fmt.Println(utf8.RuneCountInString(sd)) //字符串长度

	//同时定义多个
	var sw, sh = 12,33
	fmt.Print(sw,sh)

	var s1,s2 int
	s1 = 33
	s2 = 34
	fmt.Print(s1,s2)

	var (
		a int = 12
		b float64 = 32
	)

	const (
		Pi float64 = 32.1
		ts = true
		kls = "feeee"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(kls)
}

func v2()  {
	ah := "fffff"  //定义字符串,知道变量初始值时 (该方式只能在函数体内使用)
	var a string = "cjdddd"

	fmt.Println(len(a))
	fmt.Println(a[5])

	fmt.Println(len(ah))

	var eq bool
	var b = "cjddttdd"

	eq = a == b
	fmt.Println(eq)
}

func tps()  {
	var x = 100
	var p = &x
	fmt.Println(p,*p)

	*p = 20
	fmt.Println(*p,x)
}
