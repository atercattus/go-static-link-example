package main

import (
	"fmt"
)

/*
#cgo LDFLAGS: -L./ -lsmth
#include "smth.h"
*/
import "C"

func main() {
	fmt.Println(`before`)
	fmt.Println(C.testX());
	fmt.Println(`after`)
	return
}
