package main

import "fmt"

// import "time"
import "os"
import "os/signal"
import "syscall"

type signalHandler func(s os.Signal, arg interface{})

type signalSet struct {
	m map[os.Signal]signalHandler
}

func signalSetNew() *signalSet {
	ss := new(signalSet)
	ss.m = make(map[os.Signal]signalHandler)
	return ss
}

func (set *signalSet) register(s os.Signal, handler signalHandler) {
	if _, found := set.m[s]; !found {
		set.m[s] = handler
	}
}

func (set *signalSet) handle(sig os.Signal, arg interface{}) (err error) {
	if _, found := set.m[sig]; found {
		set.m[sig](sig, arg)
		return nil
	} else {
		return fmt.Errorf("No handler available for signal %v", sig)
	}

	panic("won't reach here")
}

func main() {
	ss := signalSetNew()
	handler := func(s os.Signal, arg interface{}) {
		fmt.Printf("handle signal: %v\n", s)
	}

	ss.register(syscall.SIGINT, handler)
	ss.register(syscall.SIGUSR1, handler)
	ss.register(syscall.SIGUSR2, handler)

	c := make(chan os.Signal)
	done := make(chan bool, 1)
	signal.Notify(c)

	go func() {
		fmt.Println("在这边阻塞。。。")
		sig := <-c
		fmt.Println("收到signal信号量。。。")
		//for sig := range ss.m {
		err := ss.handle(sig, nil)
		if err != nil {
			fmt.Printf("unknown signal received: %v\n", sig)
			os.Exit(1)
		}
		done <- true
		//}
	}()
	<-done
}
