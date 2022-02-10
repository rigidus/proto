package main

//int pt(){
// return 1;
//}
import "C"
import "fmt"

func main() {
	fmt.Println(C.pt())
}

// env CGO_ENABLED=1 GOOS=linux go build test.go
