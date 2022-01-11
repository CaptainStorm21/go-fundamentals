package main

import (
	"fmt"
	"reflect"
)

var (
	name string = "Mary Stewart"
	course string = "Geology"
	module, clip int = 9, 12
)
func main ()  {
	fmt.Println("Name is set to ", name, "takes", course, "class.")
	fmt.Println("Module and clip are set to ", module, "and", clip, ".")

	fmt.Println("Name is", reflect.TypeOf(name))
	fmt.Println("Course is", reflect.TypeOf(course))
	fmt.Println("Module is", reflect.TypeOf(module))
	fmt.Println("Clip is", reflect.TypeOf(clip))

	
}