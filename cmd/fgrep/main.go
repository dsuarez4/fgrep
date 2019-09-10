package main


import (
	"flag"
	"fmt"
)

var options struct {
	version bool
}

func main() {

	// TODO: Need to figure out way to do this in
	//flag.Usage() = func() {
	//	fmt.Printf("usage: %s [options] [query]\n", os.Args[0] )
	//	flag.PrintDefaults()
	//}
	fmt.Printf("this program works")

	flag.BoolVar(&options.version, "version", false, "print version and exit")


}
