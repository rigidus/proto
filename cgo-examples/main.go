package main

/*
   #cgo CFLAGS: -I${SRCDIR}/ctestlib
   #cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/ctestlib
   #cgo LDFLAGS: -L${SRCDIR}/ctestlib
   #cgo LDFLAGS: -ltest

   #include <test.h>
*/
import "C"
import (
    "bytes"
    "fmt"
    "unsafe"
)

type Status int

const (
    Pending Status = iota
    Done
)
type User struct {
    Username string
    Visits   int
}

func evenNumberCallback(num int) {
    fmt.Println("odd number: ", num)
}

func userCallback(user unsafe.Pointer) {
    u := (*User)(unsafe.Pointer(user))
    u.Visits++
}

func main() {
    fmt.Println("\nNumbers")
    a := 1
    b := 2
    sum := int(C.sum(C.int(a), C.int(b)))
    fmt.Print(sum, "\n\n")
    fmt.Println("Get string")
    getString := C.GoString(C.get_string())
    fmt.Println(getString)
    stringBytes := C.GoBytes(unsafe.Pointer(C.get_string()), 24)
    fmt.Println(stringBytes[0:bytes.Index(stringBytes, []byte{0})])
    fmt.Println()
    fmt.Println("Send string")
    str := "lorem ipsum"
    cStr := C.CString(str)
    C.print_string(cStr)
    C.free(unsafe.Pointer(cStr))
    fmt.Println()
    fmt.Println("Send byte array")
    data := []byte{1, 4, 2}
    cBytes := (*C.uchar)(unsafe.Pointer(&data[0]))
    cBytesLength := C.size_t(len(data))
    fmt.Print("bytes: ")
    C.print_buffer(cBytes, cBytesLength)
    fmt.Println()
    fmt.Println("Get and pass struct")
    point := C.struct_point{}
    point.x = 0
    point.y = 2
    fmt.Println(point)
    fmt.Print(C.point_diff(point), "\n\n")
    fmt.Println("Pass void pointer")
    C.pass_void_pointer(unsafe.Pointer(&point.y))
    fmt.Println()
    fmt.Println("Access enum")
    fmt.Print(C.enum_status(Pending) == C.PENDING, C.PENDING, C.DONE, "\n\n")
    fmt.Println("Pass callback")
    c := registerCallback(evenNumberCallback, nil)
    C.generate_numbers(5, c)
    unregisterCallback(c)
    user := User{
        Username: "johndoe",
    }
    cWithParams := registerCallback(userCallback, unsafe.Pointer(&user))
    C.user_action(cWithParams)
    unregisterCallback(cWithParams)
    fmt.Println(user)
}
