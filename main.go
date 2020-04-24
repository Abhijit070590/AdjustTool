package main

import (
	"adjust/tool"
	"flag"
	"os"
)

func main() {
	var parallel *uint
	parallel = flag.Uint("parallel", 10, "number of parallel routines")
	flag.Parse()
	t := tool.AdjustTool{}
	if len(os.Args) == 1 {
		println("Please enter Urls")
		return
	}
	switch os.Args[1] {
	case "-parallel":
		if len(os.Args[3:]) == 0 {
			println("Please enter Urls")
			return
		}
		t.Run(os.Args[3:], true, *parallel)
	default:
		t.Run(os.Args[1:], true, *parallel)
	}
}
