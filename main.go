package main

import (
	"./test"
        "flag"
        "os"
        "os/signal"
        "syscall"
)


var (
        testFlag bool
)

func init() {
        flag.BoolVar(&testFlag,"t",false,"test Go script")
        flag.Parse()
}

func main() {
    if testFlag {
        test.Test()
    }
    

    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
    <-sigCh
}
