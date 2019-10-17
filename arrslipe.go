package main

import (
	"fmt"
)

func main()  {


	fmt.Println("ddd")
}

func arrs()  {
	var as[3] int
	as[1] = 3
	fmt.Println(as)

	ap := [3] string{"ddd","dede","ff"}
	fmt.Println(ap)

	ai := [...] string {"ferfr","gdgd"}
	lss := len(ai)
	fmt.Println(ai,lss)

	//值类型,修改只会修改副本
	bi := ai
	bi[0] = "aaa"
	fmt.Println(ai)
	fmt.Println(bi)


	for i := 0; i < len(ai); i++{
		fmt.Println(ai[i])
	}

	for i,v := range ai{
		fmt.Println(i,v,ai[i])
	}

	for _,v := range ai{
		fmt.Println(v)
	}
}

func slicp()  {
	//1
	a := [5] int {12,34,1,3,1} //数组
	var b []int = a[1:4] //切片,取数组1~（4-1）的值
	fmt.Println(b)

	b[0] = 999
	fmt.Println(a) //a中第二个会被修改
	fmt.Println(b)

	//2
	c := []int{3,45,1,55,23,12}
	fmt.Println(c)
	for i := range c{
		c[i] ++
	}
	fmt.Println(c)


	//3
	ma := make([]int ,3)
	fmt.Println(ma,len(ma),cap(ma))
	ma = append(ma,78,55) //增加多个
	fmt.Println(ma,len(ma),cap(ma)) //len=5,cap=6 。。cap为3的倍数
	//没有声明长度和容量,系统自己优化
	ao := []int{1, 2, 3}
	fmt.Println(ao,len(ao), cap(ao))
	ao = append(ao, 4, 5, 6, 7)
	fmt.Println(ao,len(ao), cap(ao)) // len=7,cap=8

	s := append(ma[:1],ma[3:]...) //删除一个
	fmt.Println(s)


	var ms []string
	fmt.Println(ms,len(ms),cap(ms))
	ms = append(ms,"dfdf")
	fmt.Println(ms,len(ms),cap(ms))

	//函数传递修改值会影响源切片
	fmt.Println(ma)
	changeSp(ma)
	fmt.Println(ma)

	pis := [][]string{
		{"c","ccc"},
		{"fffff"},
		{"b","bbbb"},
	}

	for _,v1 := range pis{
		for _,v2 := range v1{
			fmt.Printf("%s ",v2)
		}
		fmt.Print("\n")
	}


	//4
	cps := [...]string{"fff","fefe","gegeg","eee"}
	nep := cps[:len(cps) - 2]
	cpnes := make([]string,len(nep))
	copy(cpnes,nep) //copy后对切片修改不会响应源切片

	fmt.Println(nep)
	fmt.Println(cpnes)
	cpnes[0] = "aaaa"
	fmt.Println(nep)
	fmt.Println(cpnes)

}

func changeSp(numb []int)  {
	numb[0] = 99999
}

func maoArr()  {
	pmap := make(map[string]int)
	pmap["sf"] = 23
	pmap["ge0"] = 45
	fmt.Println(pmap)

	smap := map[string]int{
		"ssd" :23,
		"ge" : 45,
	}

	smap["hr"] = 11
	fmt.Println(smap)

	fmt.Println(pmap["ge0"])

	val, ok := pmap["gege"]
	fmt.Println(val,ok)

	for k,v := range smap{
		fmt.Println(k,v)
	}

	delete(pmap,"ge0")

	m1s := map[string] int{"dd":12,"cc":22}
	m2s := map[string] int{"dd":12,"cc":22}
	bos := composer(m1s,m2s)
	fmt.Println(bos)
}

func composer(m1,m2 map[string]int) bool {
	if len(m1) != len(m2){
		return false;
	}

	for k1,v1 := range m1{
		if v2,ok := m2[k1]; ok{
			if v1 != v2{
				return false
			}
		}
	}

	return true
}

func tObj()  {
	b := 34
	var a *int = &b  //&获取地址, a指向b
	fmt.Printf("%T \n",a)
	fmt.Println(a,*a)

	*a++
	fmt.Println(b)

	c := 22
	cs := &c
	bFuc(cs)
	fmt.Println(c,*cs)

	//这种方式向函数传递一个数组指针参数，并在函数内修改数组。尽管它是有效的，但却不是 Go 语言惯用的实现方式。我们最好使用切片来处理。
	arr := [3]int{12,13,14}
	arrFuc(arr[:])
	fmt.Println(arr)
}
func bFuc(val *int)  {
	*val = 10
}
func arrFuc(val []int)  {
	val[0] = 1
}


//冒泡排序
func sortArr(arr *[5]int)  {
	fmt.Println("排序前--",(*arr))

	temp := 0
	lens := len(*arr)
	for i := 0; i< lens - 1; i++{
		for j := 0; j < lens -1 - i; j++ {
			if (*arr)[j] > (*arr)[j + 1]{
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}
	}

	fmt.Println("排序后--",(*arr))

}