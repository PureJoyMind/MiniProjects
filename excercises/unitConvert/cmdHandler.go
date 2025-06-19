package unitconvert

import (
	"flag"
	"fmt"
	"os"
)

var unit = flag.String("d", "", "The unit to convert to. Values: f (farenheight), c (celcius), k (kelvin)")
var from = flag.String("s", "", "The unit to convert from. Values: f (farenheight), c (celcius), k (kelvin)")
var val = flag.Float64("v", 0, "The value to convert (float)")

func Convert() {
	flag.Parse()
	if *unit == "" && *from == "" {
		fmt.Fprintln(os.Stderr, "No arguments provided")
		os.Exit(1)
	}

	var res string
	switch *unit {
	case "f":
		if *from == "c" {
			res = CToF(Celsius(*val)).String()
		} else if *from == "k" {
			res = KToF(Kelvin(*val)).String()
		} else {

			res = "Un-identified source"
		}

	case "c":
		if *from == "f" {
			res = FToC(Fahrenheit(*val)).String()
		} else if *from == "k" {
			res = KToC(Kelvin(*val)).String()
		} else {

			res = "Un-identified source"
		}

	case "k":
		if *from == "f" {
			res = FToK(Fahrenheit(*val)).String()
		} else if *from == "c" {
			res = CToK(Celsius(*val)).String()
		} else {

			res = "Un-identified source"
		}
	default:
		res = "Un-identified unit"
	}

	fmt.Printf("%2.2f => %s\n\r", *val, res)

}
