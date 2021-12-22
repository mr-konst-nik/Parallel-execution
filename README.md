## Parallel execution
Comparison of parallel execution calls with sync.WaitGroup and without it.

## Usage example 
Statring without sync.WaitGroup
```go
//create functions
func typeA(c chan string) {
	//some logic
	c <- "Process A finished"
}

// send a list of functions
funcArr := []func(c chan string){typeA, typeB, typeC}
Exec(funcArr)
```
Statring with sync.WaitGroup
```go
//create functions
func typeAWG(wg *sync.WaitGroup) {
	//some logic
	time.Sleep(2 * time.Second)
	fmt.Printf("Process A finished\n")
	wg.Done()
}

// send a list of functions
funcArr := []func(wg *sync.WaitGroup){typeAWG, typeBWG, typeCWG}
ExecutionWG(funcArr)
```
## Benchmark

```
go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: pExec
cpu: Intel(R) Core(TM) i3-8100 CPU @ 3.60GHz
BenchmarkExecWG-4              1        7002519400 ns/op            3008 B/op         28 allocs/op
BenchmarkExec-4                1        7016264800 ns/op            1312 B/op         20 allocs/op
PASS
ok      pExec  14.168s
```
