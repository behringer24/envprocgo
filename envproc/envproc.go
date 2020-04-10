package envproc

import (
	"fmt"
)

// Envproc returns hello world
func Envproc() string {
	return "Hello, world."
}

//ParseFlags parses command line parameters
func ParseFlags() {
	//args[1] := flag.String("word", "foo", "a string")

	//return args;
}

func main() {
	fmt.Println(Envproc())
}
