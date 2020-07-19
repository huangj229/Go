package main

import (
	    "./test"
        "flag"
)


var (
        testFlag int
)

func init() {
        flag.IntVar(&testFlag,"t", 0, "test Go script: [0]:basic func usage; [1]:net")
        flag.Parse()
}

func main() {
    switch testFlag {
    case 0 :
        test.TestFunc()
    case 1 :
        test.TestNet()
    }
    


}
