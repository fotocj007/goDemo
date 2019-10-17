package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main()  {
	demoV12()

	fmt.Println("goroutines")
}

//协程
func demoG1()  {
	go godemo() //go关键字开启协程

	for i:= 1; i <= 10;i++{
		fmt.Println("main-----",i)
		time.Sleep(time.Second)
	}
}
func godemo()  {
	for i:= 1; i <= 10; i++{
		fmt.Println("heee---",i)
		time.Sleep(time.Second)
	}
}

//全局锁
var (
	myMap = make(map [int]int , 10)
	lock sync.Mutex //全局互斥锁
)
func demoG2()  {
	for i := 1;i <= 20; i++{
		go test(i)
	}

	time.Sleep(time.Second * 5)//  主线程等待10秒,等协程执行完成

	for i,v := range myMap{
		fmt.Println(i,"---",v)
	}
}
func test(n int)  {
	res := 1;
	for i:= 1;i <= n; i++{
		res *= i
	}

	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

//管道 发送与接收默认是阻塞的
func demoGv3()  {
	done := make(chan int)
	go chans(done)

	data := <- done
	fmt.Println("chan----:",data)

	defer func() {
		fmt.Println("defer------")
		if err := recover(); err != nil{
			fmt.Println("err:==" , err)

			return
		}
	}()

	//这段报错能捕获
	//as := 0
	//ps := 12 / as
	//fmt.Println(ps)

	//已经取完了,再取就会报错,错误无法捕获
	datav2 := <- done
	fmt.Println(datav2)
}
func chans(done chan int)  {
	fmt.Println("helloww----")

	time.Sleep(time.Second * 3)
	done <- 12
}

func demoGv4()  {
	cnum := 4
	chanv1 := make(chan int)
	chanv2 := make(chan int)

	go calV1(cnum,chanv1)
	go calv2(cnum, chanv2)

	sv1,sv2 := <- chanv1, <- chanv2
	to := sv1 + sv2
	fmt.Println(to)
}
func calV1(num int,chans chan  int)  {
	num = num * 2
	chans <- num
}
func calv2(num int,chans chan int)  {
	num = num * 3
	chans <- num
}

//单向通道
func demoV5()  {
	sche := make(chan<- int)

	go func(cha chan<- int) {
		cha <- 13
		fmt.Println("11111")
	}(sche)
	//da := <-sche //报错,无法读取

	//双向通道转为单向通道,可行。反之不行
	che1 := make(chan int)
	go func(chs chan<- int) {
		chs <- 15
	}(che1)
	fmt.Println(<-che1)
}

//遍历
func demoV6()  {
	ch := make(chan int)
	go func(chas chan int) {
		for i := 1; i <= 10; i++{
			chas <- i
		}
		close(chas)
	}(ch)

	//用遍历通道必须关闭 close
	//for v := range ch{
	//	fmt.Println(v)
	//}

	//遍历2
	for{
		v, ok := <- ch
		if !ok {
			break
		}

		fmt.Println(v)
	}
}

//缓冲信道
/**
我们还可以创建一个有缓冲（Buffer）的信道。只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。同样，只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据
 */
func demoV7()  {
	//ch := make(chan string,2)
	//ch <- "dererr"
	//ch <- "gegggg"
	//fmt.Println(<- ch)
	//fmt.Println(<- ch)

	chs := make(chan int,2)

	go func(cs chan int) {
		for i := 1;i < 5; i ++{
			cs <- i //管道容量为2,所以会同时写入1，2，然后阻塞等待，读取后再写入
			fmt.Println("successfully wrote", i, "to ch")
		}
		close(cs)
	}(chs)

	time.Sleep(time.Second * 2)
	for v := range chs{
		fmt.Println(v)
		time.Sleep(time.Second * 2)
	}
}

//工作池
func demoV8()  {
	var wg sync.WaitGroup

	for i := 1;i < 4; i++{
		wg.Add(1)
		go doo1(i,&wg) //传递地址
	}

	wg.Wait()
	fmt.Println("main------")
}
func doo1(num int, wgs *sync.WaitGroup)  {
	fmt.Println("entry----",num)
	time.Sleep(time.Second * 2)

	wgs.Done()
}

//工作池实现
type Job struct {
	id int
	ran int
}
type Result struct {
	job Job
	sumo int
}
func digist(num int) int {
	num = num * 2
	time.Sleep(time.Second * 2)
	return num
}
func worker(jobs chan Job,resu chan Result,wg *sync.WaitGroup)  {
	//从工作任通道去值,进行计算后放入打印通道中
	for job := range jobs{
		oup := Result{job, digist(job.ran)}
		resu <- oup
	}

	wg.Done()
}
func createWork(nos int,jobs chan Job,res chan Result)  {
	var wg sync.WaitGroup
	for i := 0;i < nos; i++{
		wg.Add(1)
		go worker(jobs,res,&wg)  //开始工作
	}

	wg.Wait() //等待工作结束
	close(res) //关闭结果通道,方便后面遍历
}
func allcate(nos int,jobs chan Job)  {
	for i := 0;i < nos; i++{
		rnum := rand.Intn(100)
		job := Job{i,rnum}
		jobs <- job
	}
	close(jobs)
}
func resilt(done chan bool,resus chan Result)  {
	fmt.Println("11-------")
	for resu := range resus{
		fmt.Println("jos---",resu.job.id,resu.job.ran,resu.sumo)
	}
	fmt.Println("22222-------")
	done <- true
}
func demoV9()  {
	stT := time.Now()
	nos := 30
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)
	dan := make(chan bool)

	go allcate(nos,jobs) //创建工作
	go resilt(dan,results) //打印工作,results信道没有时阻塞

	fmt.Println("----创建工作池---")
	onif := 10 //增加工作池数量,时间会变短
	createWork(onif,jobs,results) //创建工作池,执行工作,完成后任务投递到results信道中

	//该条语句不能放在创建之后
	//1.results的缓冲区只有10条 2.createworkerpool 阻塞了后续读取结果操作 所以造成了死锁 你可以尝试 把 results 的缓冲区设置为100 这样就不会死锁了
	//go resilt(dan,results)

	<-dan

	end := time.Now()
	diff := end.Sub(stT)
	fmt.Println("time-----",diff.Seconds())
}


//select  select 语句用于在多个发送/接收信道操作中进行选择
//select 语句会一直阻塞，直到发送/接收操作准备就绪。如果有多个信道操作准备完毕，select 会随机地选取其中之一执行
func demoV10()  {
	ov1 := make(chan string)
	ov2 := make(chan string)

	go func(ch chan string) {
		time.Sleep(3 * time.Second)
		ch <- "ser v1"
	}(ov1)

	go func(ch chan string) {
		time.Sleep(3 * time.Second)
		ch <- "ser v2"
	}(ov2)

	time.Sleep(time.Second * 4)
	//当两个通道同时满足时,会随机选择一个
	select {
	case s1 := <- ov1:
		fmt.Println(s1)
	case s2 := <- ov2:
		fmt.Println(s2)
	default:
		fmt.Println("not value")
	}
}

//互斥体
var xs = 0
func MutexFu(wl * sync.Mutex,wg *sync.WaitGroup)  {
	//添加锁,避免竞争
	wl.Lock()
	xs = xs + 1
	wl.Unlock()

	wg.Done()
}
func demoV11()  {
	var w sync.WaitGroup
	var un sync.Mutex

	for i := 0; i < 50; i++{
		w.Add(1)
		go MutexFu(&un,&w)
	}

	w.Wait()
	fmt.Println("end to musx",xs)
}

//用互斥锁来保护一个数值型的共享资源，麻烦且效率低下。标准库的sync/atomic包对原子操作提供了丰富的支持
var totle int64
func workers(wg * sync.WaitGroup)()  {

	atomic.AddInt64(&totle,1)

	wg.Done()
}
func demoV12()  {
	var wg sync.WaitGroup

	for i := 0; i < 50; i++{
		wg.Add(1)
		go workers(&wg)
	}

	wg.Wait()
	fmt.Println("end to musx",totle)
}
