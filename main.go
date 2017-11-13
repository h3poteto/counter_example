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

	exitChan := make(chan int)
	go func() {
		for {
			s := <-signalChan
			switch s {
			case syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL:
				log.Println("exit")
				exitChan <- 0
			default:
				exitChan <- 1
			}
		}
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			log.Println("count")
		}
	}()

	code := <-exitChan
	os.Exit(code)
}
