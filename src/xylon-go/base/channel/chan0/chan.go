package main

import "fmt"
import "os"
import "os/signal"
import "syscall"
import "time"
import "math/rand"

// 该 chan 代表 线程 A 放入数据时, 只有当线程B取走数据时, 线程A才继续执行
var count = make(chan int)

// 同上
//var count = make(chan int, 0)
// 该 chan 只能存放一个数据, 当该 chan 中 已经存放一个数据时,其他的存放数据的线程,将被阻塞.
//var count = make(chan int, 1)

func main() {

	sigs := make(chan os.Signal, 2)
	done := make(chan bool)
	exit := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			time.Sleep(time.Second * 1)
			num := rand.Intn(100)
			select {
			case sig := <-sigs:
				fmt.Println(sig)
				done <- true
				//os.Exit(0)
				goto gameover
			default:
				fmt.Println("准备放入chan 数据:", num)
				count <- num
				fmt.Println("成功放入chan 数据:", num)
			}
		}
	gameover:
		fmt.Println("exiting 1")
	}()

	go func() {
		for {
			time.Sleep(time.Second * 5)
			select {
			case <-done:
				//os.Exit(0)
				exit <- true
				goto gameover
			case num := <-count:
				fmt.Println("成功取出chan 数据:", num)
			default:
				fmt.Println("无数据")
			}
		}
	gameover:
		fmt.Println("exiting 2")
	}()

	fmt.Println("awaiting signal")
	<-exit
	fmt.Println("exited")
}
