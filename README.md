# Parallel execution
Example of working goroutines without primitives of synchronization.

# Usage example 

```go
#create functions
func typeA(c chan string) {
	#your code here
	c <- "Process A finished"
}

# send a list of functions
funcArr := []func(c chan string){typeA, typeB, typeC}
pExecution(funcArr)
```
