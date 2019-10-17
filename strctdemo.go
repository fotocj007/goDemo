package main

import (
	"fmt"
)

func main()  {
	funcDemo()

	fmt.Println("ss")
}

//方法
func funcDemo()  {
	ep1 := Epo{"ks",23}  //实现一个结构体
	ep1.dist() //对应的方法
	fmt.Println(ep1)
	ep1.distv(10)
	fmt.Println(ep1)
}

type Epo struct {
	name string
	cit int
}
func (e Epo) dist() {
	e.name = "aa"
	//fmt.Println(e.name,e.cit)
}
func (e *Epo) distv(tis int) {
	e.cit = tis
}


//结构体
func demos()  {
	type emop struct {
		fs string
		ks string
		gs int
	}

	emp := emop{
		fs: "sfff",
		ks: "gegeg",
		gs: 30,
	}

	emp2 := emop{"ffff","gegg", 34}

	fmt.Println(emp)
	fmt.Println(emp2)
	fmt.Println(emp.fs,emp2.gs)


	var epms emop
	epms.gs = 21
	epms.fs = "egeg"

	emp3 := struct {
		fs string
		age int
	}{
		"ffff",
		12,
	}

	fmt.Println(emp3)

	var emp4 emop
	fmt.Println(emp4)

	//匿名字段
	type per struct {
		string
		int
	}

	ps := per{"fdfdf",34}
	fmt.Println(ps)

	//套签结构体
	type adds struct {
		city, stat string
	}
	type Ads struct {
		name string
		adds
	}

	aps := adds{"citsy","sts"}
	pps := Ads{
		name:"fefef",
		adds: aps,
	}

	fmt.Println(pps.name,pps.city,pps.stat)

}
