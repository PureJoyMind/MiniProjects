package excercises

import (
	"flag"
	"fmt"
	"strings"
)

var n *bool = flag.Bool("n", false, "omit trailing newline")
var sep *string = flag.String("s", " ", "separator")

func Echo() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
