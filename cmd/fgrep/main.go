package main


import (
	"flag"
	"fmt"
	"github.com/dsuarez4/fgrep/meta"
	"os"
)

var options struct {
	version bool
}

func main() {

	fmt.Printf("this program compiles. That's it \n")

	flag.BoolVar(&options.version, "version", false,
		"print version and exit")
	flag.BoolVar(&options.version, "v", false,
		"print version and exit (shorthand)")

	if options.version {
		fmt.Printf("%s\n", meta.Meta())
		os.Exit(0)
	}
	flag.Parse()


}
