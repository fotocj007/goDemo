package main

import "fmt"

func main()  {
	demoV4()

	fmt.Println("ddddd")
}


//自定义一个类型,类型是函数，两个参数和一个返回值
type add func(a, b int) int

func demov1()  {
	var a add = func(a, b int) int{
		return a + b
	}

	s := a(1,3)
	fmt.Println(s)
}

//高阶函数,参数是一个函数
func simple(as func(a,b int) int) int {
	return as(12,10)
}
func demoV2()  {
	f := func(a, b int) int{
		return a + b
	}

	ps := simple(f)
	fmt.Println(ps)
}

//返回一个函数
func sv2() func(a,b int) int{
	fs := func(a, b int) int{
		return a + b
	}

	return fs
}
func demoV3()  {
	fs := sv2()
	all := fs(10,2)
	fmt.Println(all)
}

//闭包
func addStr() func(string ) string{
	t := "he:"
	c := func(b string) string{
		t = t + " " + b
		return t
	}

	return c
}
func demoV4()  {
	a := addStr()

	p1 := a("AAA")
	p2 := a("BBB")
	fmt.Println(p1)
	fmt.Println(p2)
}