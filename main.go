package main

import (
	"./test"
        "flag"
        "os"
        "os/signal"
        "syscall"
)


var (
        testFlag int
)

func init() {
        flag.IntVar(&testFlag,"t", 0, "test Go script: [0]:basic func; [1]:net")
        flag.Parse()
}

func main() {
    switch testFlag {
    case 0 :
        test.Test()
    case 1 :
        test.TestNet()
    }
    

    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
    <-sigCh
}
