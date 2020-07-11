package main

import (
	"./test"
        "flag"
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
}
