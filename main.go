package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	stopChan := make(chan int, 1)

	go func(stopChan chan int) {
		count := 0
		for {
			select {
			case _ = <-stopChan:
				log.Println("loop stopped")
				break
			default:
				time.Sleep(1 * time.Second)
				count++
				log.Println("count", count)
			}
		}
	}(stopChan)

	for {
		select {
		case ch := <-signalChan:
			stopChan <- 1
			code := catchSig(ch)
			os.Exit(code)
		}
	}

}

func catchSig(sig os.Signal) int {
	switch sig {
	case syscall.SIGHUP:
		log.Println("SIGHUP Happend! ", sig)
		return 0
	case syscall.SIGTERM:
		log.Println("SIGTERM Happend! ", sig)
		return 0
	case syscall.SIGKILL:
		log.Println("SIGKILL Happend! ", sig)
		return 0
	default:
		log.Println("Other singal Happend! ", sig)
		return 1
	}
}
