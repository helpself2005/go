1、go执行的两种方式 go run ***.go 解释执行，go build ***.go 编译执行
2、引用Package
	1）在import中，你可以使用相对路径，如 ./或 ../ 来引用你的package
	2）如果没有使用相对路径，那么，go会去找$GOPATH/src/目录。

	使用相对路径 import "./haoel"  //import当前目录里haoel子目录里的所有的go文件
	使用GOPATH路径 import "haoel"  //import 环境变量 $GOPATH/src/haoel子目录里的所有的go文件
3、fmt输出格式，Println不支持，Printf才支持%式的输出
4、变量和常量
	变量的声明很像 javascript，使用 var关键字。还有一种定义变量的方式（这让我想到了Pascal语言，但完全不一样）x := 100 //等价于 var x int = 100;
	常量很简单，使用const关键字：const pi float32 = 3.1415926
5、数组
	var a [5]int   var c [2][3]int
	数组的切片操作 a := [5]int{1,2, 3,4, 5}  b := a[2:4] b = a[2:]  b = a[:4]
6、分支循环语句
	if 语句没有圆括号，而必需要有花括号，花括号必须和if一行
	switch语句没有break，还可以使用逗号case多个值
	switch i {
	    case  1:  fmt.Println("one")
	    case  2:  fmt.Println("two")
	    case  3:  fmt.Println("three")
	    case  4,5,6:  fmt.Println("four, five, six")
	    default:      fmt.Println("invalid value!")
	}
	for 语句 Go语言中没有while
	//经典的for语句 init; condition; post
	for 
	i := 0; i<10; i++{
	     fmt.Println(i)
	}
	 
	//精简的for语句 condition
	i := 1
	for 
	i<10 {
	    fmt.Println(i)
	    i++
	}
	 
	//死循环的for语句 相当于for(;;)
	i :=1
	for
	{
	    if
	i>10 {
	        break
	    }
	    i++
	}
	注意：无论任何时候，你都不应该将一个控制结构（(if、for、switch或select）的左大括号放在下一行。如果这样做，将会在大括号的前方插入一个分号，这可能导致出现不想要的结果。

7、关于分号
	和C一样，Go的正式的语法使用分号来终止语句。和C不同的是，这些分号由词法分析器在扫描源代码过程中使用简单的规则自动插入分号，因此输入源代码多数时候就不需要分号了。

8、map
	m := make(map[string]int)
	m["one"] = 1
    m["two"] = 2
    m["three"] = 3
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}

	for key, val := range m1{
        fmt.Printf("%s => %d \n", key, val)
        /*输出：(顺序在运行时可能不一样)
            three => 3
            one => 1
            two => 2*/
    }

9、指针
  	Go具有两个分配内存的机制，分别是内建的函数new和make。他们所做的事不同，所应用到的类型也不同，这可能引起混淆，但规则却很简单。

10、内存分配
	new 是一个分配内存的内建函数，但不同于其他语言中同名的new所作的工作，它只是将内存清零，而不是初始化内存。new(T)为一个类型为T的新项目分配了值为零的存储空间并返回其地址，也就是一个类型为*T的值。用Go的术语来说，就是它返回了一个指向新分配的类型为T的零值的指针。

	make(T, args)函数的目的与new(T)不同。它仅用于创建切片、map和chan（消息管道），并返回类型T（不是*T）的一个被初始化了的（不是零）实例。这种差别的出现是由于这三种类型实质上是对在使用前必须进行初始化的数据结构的引用。例如，切片是一个具有三项内容的描述符，包括指向数据（在一个数组内部）的指针、长度以及容量，在这三项内容被初始化之前，切片值为nil。对于切片、映射和信道，make初始化了其内部的数据结构并准备了将要使用的值。

11、函数
	函数可以返回多个值，一个正常值，一个是错误；函数传参可不定参数，用数组传递；同时支持函数闭包、递归

12、结构体
	Go语言中没有public, protected, private的关键字，所以，如果你想让一个方法可以被别的包访问的话，你需要把这个方法的第一个字母大写。这是一种约定。

13、接口和多态
	接口的使用是通过指针来实现,Error接口

14、defer关键字延迟语句执行到函数执行完
	panic：当panic被调用时，它将立即停止当前函数的执行并开始逐级解开函数堆栈，同时运行所有被defer的函数。如果这种解开达到堆栈的顶端，程序就死亡了。但是，也可以使用内建的recover函数来重新获得Go程的控制权并恢复正常的执行。 对recover的调用会通知解开堆栈并返回传递到panic的参量。由于仅在解开期间运行的代码处在被defer的函数之内，recover仅在被延期的函数内部才是有用的。

	可以简单地理解为recover就是用来捕捉Painc的，防止程序一下子就挂掉了。
	在defer里recover来回复程序

15、goroutine
	GoRoutine主要是使用go关键字来调用函数，你还可以使用匿名函数.
	go 多线程始终没有打印出多线程的输出
	同时支持 atomic原子操作

16、Chanenel信道
	Channal是什么？Channal就是用来通信的，就像Unix下的管道一样，在Go中是这样使用Channel的。

	Channel的阻塞
	注意，channel默认上是阻塞的，也就是说，如果Channel满了，就阻塞写，如果Channel空了，就阻塞读。于是，我们就可以使用这种特性来同步我们的发送和接收端.

	Channel阻塞的这个特性还有一个好处是，可以让我们的goroutine在运行的一开始就阻塞在从某个channel领任务，这样就可以作成一个类似于线程池一样的东西。关于这个程序我就不写了。我相信你可以自己实现的。


17、定时器
	Go语言中可以使用time.NewTimer或time.NewTicker来设置一个定时器，这个定时器会绑定在你的当前channel中，通过channel的阻塞通知机器来通知你的程序。

	注意Timer只通知一次。如果你要像C中的Timer能持续通知的话，你需要使用Ticker。


18、Socket编程

19、系统调用







