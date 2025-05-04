package main

import (
	"fmt"
	"unsafe"
)


// unsafe.Pointer access the underlying address without
// regards for golang type system
func main() {
	i := 10
	iptr := unsafe.Pointer(&i)
	fmt.Println(iptr)
	fmt.Println(*(*int)(iptr))
	// fmt.Println(8(*string)(iptr))
	// fmt.Println(8(*float64)(iptr))


	// pointer arithmetic
	arr := []int{1,2,3,4,5,6}
	arrPtr := unsafe.Pointer(&arr[0])
	next := (*int)(unsafe.Pointer(uintptr(arrPtr) + unsafe.Sizeof(arr[0])))
	fmt.Println(*next)

	for i := 0; i < len(arr); i++ {
		next := (*int)(unsafe.Pointer(uintptr(arrPtr) + uintptr(i) * unsafe.Sizeof(arr[0])))
	fmt.Println(*next)

	}
}