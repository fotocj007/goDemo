package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main()  {
	mv2gete()
	fmt.Println("trysss")
}

func trysdemo()  {
	defer func() {
		err := recover()
		if err != nil{
			fmt.Println("err:==" , err)
		}
	}()

	av1 := 23
	av2 := 0
	ps := av1 / av2
	fmt.Println(ps)
	fmt.Println("ttttt----")
}

func passV1()  {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println("er---",err)
		}
	}()

	//go psss()  //不能捕获其他协程抛出的异常
	psss()

	time.Sleep(2 * time.Second)
	fmt.Println("1111----")
}
func psss()  {
	fmt.Println("2222222----")
	panic("抛出异常")
	fmt.Println("abc-------")
}

func wrfile()  {
	f,err := os.Open("test.txt")
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(f.Name(),"op------")
}

func defff()  {
	//先加入的后执行,加入一个独立的栈中,不执行
	//当函数执行完毕后执行
	defer func() {
		fmt.Println("end--111")
	}()

	defer func() {
		fmt.Println("end ---22")
	}()

	res := "geddd"
	fmt.Println(res)
}

func dev2()  {
	a := 12
	defer dfus(a)
	a = 10
	fmt.Println("nnnn----",a)
}
func dfus(a int)  {
	fmt.Println(a)
}


/*******自定义错误*******/

func trPic(st int) (int, error) {
	if st < 0{
		return 0,errors.New("小于0 -----")  //自定义一个错误
	}
	return st,nil
}
func mv2gete()  {
	res := -2

	is , er := trPic(res)
	if er != nil{
		fmt.Println("err----",er)
		panic(er)  //抛出这个错误
	}

	fmt.Println("ok-----",is)
}