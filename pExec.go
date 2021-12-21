package pExec

import (
	"fmt"
	"time"
)

func pExecution(f []func(c chan string)) {

	lenF := len(f)
	ch := make(chan string, lenF)

	for _, fnc := range f {
		go fnc(ch)
	}

	for i := 0; i < lenF; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
}

func typeA(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "Process A finished"
}

func typeB(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Process B finished"
}

func typeC(c chan string) {
	time.Sleep(7 * time.Second)
	c <- "Process C finished"
}

func Exec() {

	funcArr := []func(c chan string){typeA, typeB, typeC}

	fmt.Println("...Start goroutines...")
	pExecution(funcArr)
	fmt.Println("...Finish goroutines...")
}
