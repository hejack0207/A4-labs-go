package main

import "fmt"
import "testing"

func panicRecover() {
	defer catch()
	triggerPanic()
}

func catch() {
	if err := recover(); err != nil {
		fmt.Println("catch the panic:", err)
	}
}

func triggerPanic() {
	panic("error occurred!")
}

func main() {
	panicRecover()
}
