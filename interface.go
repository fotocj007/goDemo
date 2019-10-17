package main

import (
	"fmt"
)

func main()  {
	fmt.Println("interface")
}
/****************************/
//断言
func cs()  {
	var s interface{} = 80
	psdd(s)

	var st interface{} = "dddd"
	psdd(st)
}
func psdd(i interface{})  {
	s , ok := i.(int) //提取i底层的int值,判断类型是否满足
	fmt.Println(s , ok) //80,true   0，false

	//通过选择进行打印
	switch i.(type) {
	case string:
		fmt.Println(i.(string))
	case int:
		fmt.Println(i.(int))
	}

}

/****************************/
//定义接口
type Inarr interface {
	totleAll() int
}

//定义两个结构体
type Per1 struct {
	ps1 int
	ps2 int
}
type Per2 struct {
	as int
}

//结构体实现接口的方法
func (ps Per1) totleAll() int {
	return ps.ps1 + ps.ps2
}
func (ps Per2) totleAll() int {
	return ps.as
}
//汇总数据
func totleAlls(s []Inarr) int {
	a := 0
	for _,v := range s{
		a = a + v.totleAll()
	}
	return a
}

//调用,今后扩充Per时非常方便
func get()  {
	p1 := Per1{12,33}
	p2 := Per2{10}

	allP := []Inarr{p1,p2} //per实现了接口的方法,所以可以赋值给接口类型
	alls := totleAlls(allP)
	fmt.Println(alls)
}
/****************************/

/****************************
/****************************
user := Users{"nasss",12}
var vis Inters
vis = user
vis.FindV()
 */
func (us Users) FindV()  {
	fmt.Println(us.name)
}

type Inters interface {
	FindV()
}
type Users struct {
	name string
	age int
}
/****************************/
/****************************/
//指针接受者与值接受者
type Pinser interface {
	Dists() int
}
type Pss struct {
	name string
	eag int
}

func (us *Pss) Dists() int { // 使用指针接受者实现
	fmt.Println(us.eag)
	return us.eag
}

func getPos()  {
	pss := Pss{"cjddd",13}
	var p1 Pinser
	//p1 = pss  //非法
	p1 = &pss
	p1.Dists()
}

/****************************/