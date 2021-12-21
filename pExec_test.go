package pExec

import (
	"testing"
)

func BenchmarkExecWG(t *testing.B) {

	ExecWG()

}

func BenchmarkExec(t *testing.B) {

	Exec()

}
