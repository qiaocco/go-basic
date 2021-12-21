package main

import (
	"flag"
	"fmt"
	myflag "go-basic/10-interface/11-flag"
	"os"
	"time"
)

var period = flag.Duration("peroid", time.Second, "sleep period")
var temp = myflag.SheshiFlag("temp", 20, "the temperature")

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])
	flag.Parse()
	fmt.Printf("parsed value= %v...", *temp)
	fmt.Println()

}
