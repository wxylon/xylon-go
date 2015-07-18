//http://www.cnblogs.com/getong/archive/2013/03/19/2970045.html

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	done := make(chan bool)
	receive_channel := make(chan chan bool)
	finish := make(chan bool)

	for i := 0; i < runtime.NumCPU(); i++ {
		go do_while_select(i, receive_channel, finish)
	}

	fmt.Println("len(receive_channel)--->", len(receive_channel))

	go handle_exit(done, receive_channel, finish)

	<-done
	os.Exit(0)

}
func handle_exit(done chan bool, receive_channel chan chan bool, finish chan bool) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	chan_slice := make([]chan bool, 0)
	for {
		//fmt.Println("len(receive_channel)--->",len(receive_channel));
		select {
		case <-sigs:
			for _, v := range chan_slice {
				v <- true
			}
			for i := 0; i < len(chan_slice); i++ {
				<-finish
			}
			done <- true
			runtime.Goexit()
		case single_chan := <-receive_channel:
			log.Println("the single_chan is ", single_chan)
			chan_slice = append(chan_slice, single_chan)
		default:
			//log.Println("default")
		}
	}
}
func do_while_select(num int, rece chan chan bool, done chan bool) {
	quit := make(chan bool)
	rece <- quit
	for {
		time.Sleep(time.Second * 3)
		select {
		case <-quit:
			log.Println("the ", num, "is quit")
			done <- true
			runtime.Goexit()
		default:
			//简单输出
			log.Println("the ", num, "is running")

		}
	}
}
