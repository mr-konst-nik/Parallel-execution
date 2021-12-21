package pExec

import (
	"fmt"
	"sync"
	"time"
)

func typeAWG(wg *sync.WaitGroup) {
	//some logic
	time.Sleep(2 * time.Second)
	fmt.Printf("Process A finished\n")
	wg.Done()
}

func typeBWG(wg *sync.WaitGroup) {
	//some logic
	time.Sleep(5 * time.Second)
	fmt.Printf("Process B finished\n")
	wg.Done()
}

func typeCWG(wg *sync.WaitGroup) {
	//some logic
	time.Sleep(7 * time.Second)
	fmt.Printf("Process C finished\n")
	wg.Done()
}

func pExecutionWG(f []func(wg *sync.WaitGroup)) {
	var wg sync.WaitGroup
	for _, fnc := range f {
		wg.Add(1)
		go fnc(&wg)
	}
	wg.Wait()
}

func ExecWG() {

	funcArr := []func(wg *sync.WaitGroup){typeAWG, typeBWG, typeCWG}

	fmt.Println("...Start goroutines...")
	pExecutionWG(funcArr)
	fmt.Println("...Finish goroutines...")
}
